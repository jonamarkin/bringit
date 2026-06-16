package server

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ojaami/bringit/backend/internal/auth"
	"github.com/ojaami/bringit/backend/internal/config"
	"github.com/ojaami/bringit/backend/internal/mailer"
	"github.com/ojaami/bringit/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	cfg    config.Config
	logger *slog.Logger
	db     *sql.DB
	mux    *http.ServeMux

	tokenService *auth.TokenService
	mailer       *mailer.Mailer

	models struct {
		Users         *models.UserModel
		OTPs          *models.OTPModel
		Events        *models.EventModel
		Notifications *models.NotificationModel
	}
}

func New(cfg config.Config, logger *slog.Logger, db *sql.DB) *Server {
	if logger == nil {
		logger = slog.Default()
	}

	s := &Server{
		cfg:          cfg,
		logger:       logger,
		db:           db,
		mux:          http.NewServeMux(),
		tokenService: auth.NewTokenService(cfg.JWTSecret),
		mailer:       mailer.New(cfg.ResendAPIKey, cfg.MailSender),
	}
	s.models.Users = &models.UserModel{DB: db}
	s.models.OTPs = &models.OTPModel{DB: db}
	s.models.Events = &models.EventModel{DB: db}
	s.models.Notifications = &models.NotificationModel{DB: db}

	s.routes()
	return s
}

func (s *Server) Handler() http.Handler {
	return s.withSecurityHeaders(s.withCORS(s.withRequestLogging(s.mux)))
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.handleHealth)
	s.mux.HandleFunc("GET /api/v1/status", s.handleStatus)

	s.mux.HandleFunc("POST /api/v1/auth/otp/request", s.handleOTPRequest)
	s.mux.HandleFunc("POST /api/v1/auth/otp/verify", s.handleOTPVerify)
	s.mux.HandleFunc("POST /api/v1/auth/logout", s.handleLogout)

	s.mux.HandleFunc("GET /api/v1/public/events/{shareCode}", s.handlePublicEvent)
	s.mux.HandleFunc("GET /api/v1/public/events/{shareCode}/stream", s.handlePublicEventStream)
	s.mux.HandleFunc("POST /api/v1/public/events/{shareCode}/rsvp", s.handlePublicRSVP)
	s.mux.HandleFunc("POST /api/v1/public/events/{shareCode}/claims", s.handlePublicCreateClaim)
	s.mux.HandleFunc("DELETE /api/v1/public/events/{shareCode}/claims/{claimID}", s.handlePublicCancelClaim)

	protected := http.NewServeMux()
	protected.HandleFunc("GET /api/v1/auth/me", s.handleAuthMe)
	protected.HandleFunc("GET /api/v1/events", s.handleListEvents)
	protected.HandleFunc("POST /api/v1/events", s.handleCreateEvent)
	protected.HandleFunc("GET /api/v1/events/{eventID}", s.handleGetEvent)
	protected.HandleFunc("POST /api/v1/events/{eventID}/items", s.handleAddItem)
	protected.HandleFunc("GET /api/v1/events/{eventID}/notifications", s.handleListNotifications)
	protected.HandleFunc("POST /api/v1/events/{eventID}/notifications/mark-read", s.handleMarkNotificationsRead)

	s.mux.Handle("/api/v1/", s.tokenService.RequireAuth(protected))
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"service": "bringit-api",
		"status":  "ok",
	})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"environment": s.cfg.Env,
		"service":     "bringit-api",
		"status":      "ready",
	})
}

type otpRequest struct {
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func (s *Server) handleOTPRequest(w http.ResponseWriter, r *http.Request) {
	var req otpRequest
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		writeError(w, http.StatusBadRequest, "enter a valid email address")
		return
	}

	code, err := randomDigits(6)
	if err != nil {
		s.logger.Error("failed to generate otp", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to hash otp", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if err := s.models.OTPs.Insert(r.Context(), req.Email, string(hash), time.Now().Add(10*time.Minute)); err != nil {
		s.logger.Error("failed to save otp", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	emailData := map[string]any{"Code": code}
	if id, err := s.mailer.Send(req.Email, "Your BringIt sign-in code", "otp.html", emailData); err != nil {
		s.logger.Error("failed to send otp email", "error", err)
	} else if id != "" {
		s.logger.Info("otp email sent", "email", req.Email, "resend_message_id", id)
	}

	res := map[string]any{"message": "OTP sent successfully"}
	if s.cfg.Env != "production" {
		res["dev_code"] = code
	}
	writeJSON(w, http.StatusOK, res)
}

type otpVerifyRequest struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	DisplayName string `json:"display_name"`
}

func (s *Server) handleOTPVerify(w http.ResponseWriter, r *http.Request) {
	var req otpVerifyRequest
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Code = strings.TrimSpace(req.Code)
	if req.Email == "" || req.Code == "" {
		writeError(w, http.StatusBadRequest, "email and code are required")
		return
	}

	otp, err := s.models.OTPs.LatestValid(r.Context(), req.Email)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid or expired OTP")
		return
	}
	if otp.Attempts >= 3 {
		writeError(w, http.StatusUnauthorized, "too many failed attempts")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(otp.CodeHash), []byte(req.Code)); err != nil {
		_ = s.models.OTPs.IncrementAttempts(r.Context(), otp.ID)
		writeError(w, http.StatusUnauthorized, "invalid OTP")
		return
	}

	user, err := s.models.Users.UpsertEmailUser(r.Context(), req.Email, req.DisplayName)
	if err != nil {
		s.logger.Error("failed to upsert user", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	_ = s.models.OTPs.MarkUsed(r.Context(), otp.ID)

	token, err := s.tokenService.GenerateToken(user.ID)
	if err != nil {
		s.logger.Error("failed to generate token", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	s.setAuthCookie(w, token)

	writeJSON(w, http.StatusOK, map[string]any{
		"user":  user,
		"token": token,
	})
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.clearAuthCookie(w)
	writeJSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}

func (s *Server) handleAuthMe(w http.ResponseWriter, r *http.Request) {
	user, err := s.models.Users.Get(r.Context(), auth.GetUserID(r.Context()))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "user not found")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"user": user})
}

func (s *Server) handleListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := s.models.Events.ListByHost(r.Context(), auth.GetUserID(r.Context()))
	if err != nil {
		s.logger.Error("failed to list events", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"events": events})
}

func (s *Server) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var req models.EventCreateInput
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "event title is required")
		return
	}

	shareCode, err := randomCode(8)
	if err != nil {
		s.logger.Error("failed to generate share code", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	event, err := s.models.Events.Create(r.Context(), auth.GetUserID(r.Context()), shareCode, req)
	if err != nil {
		s.logger.Error("failed to create event", "error", err)
		writeError(w, http.StatusInternalServerError, "could not create event")
		return
	}
	snapshot, _ := s.models.Events.Snapshot(r.Context(), event.ID)
	writeJSON(w, http.StatusCreated, map[string]any{
		"event":     event,
		"snapshot":  snapshot,
		"share_url": s.publicEventURL(event.ShareCode),
	})
}

func (s *Server) handleGetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	if _, err := s.models.Events.GetForHost(r.Context(), eventID, auth.GetUserID(r.Context())); err != nil {
		writeModelError(w, err)
		return
	}
	snapshot, err := s.models.Events.Snapshot(r.Context(), eventID)
	if err != nil {
		writeModelError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"snapshot":  snapshot,
		"share_url": s.publicEventURL(snapshot.Event.ShareCode),
	})
}

func (s *Server) handleAddItem(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	var req models.ItemInput
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if strings.TrimSpace(req.Name) == "" {
		writeError(w, http.StatusBadRequest, "item name is required")
		return
	}
	item, err := s.models.Events.AddItem(r.Context(), eventID, auth.GetUserID(r.Context()), req)
	if err != nil {
		writeModelError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"item": item})
}

func (s *Server) handleListNotifications(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	if _, err := s.models.Events.GetForHost(r.Context(), eventID, auth.GetUserID(r.Context())); err != nil {
		writeModelError(w, err)
		return
	}
	notifications, err := s.models.Notifications.ListForUser(r.Context(), auth.GetUserID(r.Context()), eventID)
	if err != nil {
		s.logger.Error("failed to list notifications", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"notifications": notifications})
}

func (s *Server) handleMarkNotificationsRead(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	if _, err := s.models.Events.GetForHost(r.Context(), eventID, auth.GetUserID(r.Context())); err != nil {
		writeModelError(w, err)
		return
	}
	if err := s.models.Notifications.MarkRead(r.Context(), auth.GetUserID(r.Context()), eventID); err != nil {
		s.logger.Error("failed to mark notifications read", "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "notifications marked read"})
}

func (s *Server) handlePublicEvent(w http.ResponseWriter, r *http.Request) {
	event, err := s.models.Events.GetByShareCode(r.Context(), r.PathValue("shareCode"))
	if err != nil {
		writeModelError(w, err)
		return
	}
	snapshot, err := s.models.Events.Snapshot(r.Context(), event.ID)
	if err != nil {
		writeModelError(w, err)
		return
	}
	sanitizePublicSnapshot(snapshot)
	writeJSON(w, http.StatusOK, map[string]any{
		"snapshot": snapshot,
	})
}

func (s *Server) handlePublicEventStream(w http.ResponseWriter, r *http.Request) {
	event, err := s.models.Events.GetByShareCode(r.Context(), r.PathValue("shareCode"))
	if err != nil {
		writeModelError(w, err)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, "streaming is not supported")
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	send := func() bool {
		snapshot, err := s.models.Events.Snapshot(r.Context(), event.ID)
		if err != nil {
			return false
		}
		sanitizePublicSnapshot(snapshot)
		payload, err := json.Marshal(snapshot)
		if err != nil {
			return false
		}
		_, _ = fmt.Fprintf(w, "event: snapshot\n")
		_, _ = fmt.Fprintf(w, "data: %s\n\n", payload)
		flusher.Flush()
		return true
	}

	if !send() {
		return
	}

	for {
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
			if !send() {
				return
			}
		}
	}
}

func (s *Server) handlePublicRSVP(w http.ResponseWriter, r *http.Request) {
	event, err := s.models.Events.GetByShareCode(r.Context(), r.PathValue("shareCode"))
	if err != nil {
		writeModelError(w, err)
		return
	}

	var req models.GuestInput
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}
	if req.SessionToken == "" {
		req.SessionToken, err = randomCode(32)
		if err != nil {
			s.logger.Error("failed to generate guest token", "error", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	guest, change, err := s.models.Events.UpsertGuest(r.Context(), event.ID, req.SessionToken, req)
	if err != nil {
		writeModelError(w, err)
		return
	}
	s.notifyHostsAsync(event.ID, change)

	writeJSON(w, http.StatusOK, map[string]any{
		"guest": guest,
	})
}

func (s *Server) handlePublicCreateClaim(w http.ResponseWriter, r *http.Request) {
	event, err := s.models.Events.GetByShareCode(r.Context(), r.PathValue("shareCode"))
	if err != nil {
		writeModelError(w, err)
		return
	}

	var req models.ClaimInput
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.ItemID == "" {
		writeError(w, http.StatusBadRequest, "item_id is required")
		return
	}

	claim, change, err := s.models.Events.CreateClaim(r.Context(), event.ID, req.SessionToken, req)
	if err != nil {
		writeModelError(w, err)
		return
	}
	s.notifyHostsAsync(event.ID, change)

	writeJSON(w, http.StatusCreated, map[string]any{"claim": claim})
}

func (s *Server) handlePublicCancelClaim(w http.ResponseWriter, r *http.Request) {
	event, err := s.models.Events.GetByShareCode(r.Context(), r.PathValue("shareCode"))
	if err != nil {
		writeModelError(w, err)
		return
	}

	var req struct {
		SessionToken string `json:"session_token"`
	}
	if err := readJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	change, err := s.models.Events.CancelClaim(r.Context(), event.ID, req.SessionToken, r.PathValue("claimID"))
	if err != nil {
		writeModelError(w, err)
		return
	}
	s.notifyHostsAsync(event.ID, change)

	writeJSON(w, http.StatusOK, map[string]string{"message": "claim removed"})
}

func (s *Server) notifyHostsAsync(eventID string, change *models.ChangeNotification) {
	if change == nil || len(change.HostEmails) == 0 {
		return
	}
	for _, email := range change.HostEmails {
		email := email
		go func() {
			data := map[string]any{
				"EventTitle":   change.EventTitle,
				"Title":        change.Title,
				"Message":      change.Message,
				"DashboardURL": s.cfg.PublicBaseURL + "/#/host?event=" + eventID,
			}
			if id, err := s.mailer.Send(email, "BringIt: "+change.Title, "host_notification.html", data); err != nil {
				s.logger.Error("failed to send host notification", "email", email, "error", err)
			} else if id != "" {
				s.logger.Info("host notification sent", "email", email, "resend_message_id", id)
			}
		}()
	}
}

func sanitizePublicSnapshot(snapshot *models.EventSnapshot) {
	for _, guest := range snapshot.Guests {
		guest.Email = ""
		guest.Phone = ""
		guest.SessionToken = ""
	}
}

func (s *Server) publicEventURL(shareCode string) string {
	return strings.TrimRight(s.cfg.PublicBaseURL, "/") + "/#/event/" + shareCode
}

func (s *Server) setAuthCookie(w http.ResponseWriter, token string) {
	isProd := s.cfg.Env == "production"
	maxAge := 14 * 24 * 60 * 60
	http.SetCookie(w, &http.Cookie{
		Name:     "bringit_auth",
		Value:    token,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
		Domain:   s.cfg.CookieDomain,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "bringit_logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: false,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
		Domain:   s.cfg.CookieDomain,
	})
}

func (s *Server) clearAuthCookie(w http.ResponseWriter) {
	isProd := s.cfg.Env == "production"
	for _, name := range []string{"bringit_auth", "bringit_logged_in"} {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: name == "bringit_auth",
			Secure:   isProd,
			SameSite: http.SameSiteLaxMode,
			Domain:   s.cfg.CookieDomain,
		})
	}
}

func (s *Server) withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && s.originAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) originAllowed(origin string) bool {
	if s.cfg.FrontendOrigin == "*" {
		if s.cfg.Env == "production" {
			s.logger.Warn("CORS wildcard is enabled in production")
		}
		return true
	}
	for _, allowed := range strings.Split(s.cfg.FrontendOrigin, ",") {
		if strings.TrimSpace(allowed) == origin {
			return true
		}
	}
	return false
}

func (s *Server) withSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) withRequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		s.logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func readJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeModelError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrRecordNotFound):
		writeError(w, http.StatusNotFound, "record not found")
	case errors.Is(err, models.ErrNotEventHost):
		writeError(w, http.StatusForbidden, "you do not have access to this event")
	case errors.Is(err, models.ErrQuantityUnavailable):
		writeError(w, http.StatusConflict, "not enough quantity remaining")
	case errors.Is(err, models.ErrGuestSessionRequired):
		writeError(w, http.StatusBadRequest, "RSVP before claiming an item")
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func randomDigits(length int) (string, error) {
	var b strings.Builder
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		b.WriteString(n.String())
	}
	return b.String(), nil
}

func randomCode(length int) (string, error) {
	const alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	var b strings.Builder
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		b.WriteByte(alphabet[n.Int64()])
	}
	return b.String(), nil
}
