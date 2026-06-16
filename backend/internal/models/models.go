package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrNotEventHost         = errors.New("user is not a host for this event")
	ErrQuantityUnavailable  = errors.New("not enough quantity remaining")
	ErrGuestSessionRequired = errors.New("guest session is required")
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) UpsertEmailUser(ctx context.Context, email string, displayName string) (*User, error) {
	stmt := `
		INSERT INTO users (email, display_name)
		VALUES ($1, $2)
		ON CONFLICT (email) DO UPDATE SET
			display_name = COALESCE(NULLIF(EXCLUDED.display_name, ''), users.display_name),
			updated_at = NOW()
		RETURNING id, email, display_name, avatar_url, created_at, updated_at`

	var user User
	err := m.DB.QueryRowContext(ctx, stmt, email, displayName).Scan(
		&user.ID,
		&user.Email,
		&user.DisplayName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Get(ctx context.Context, id string) (*User, error) {
	stmt := `
		SELECT id, email, display_name, avatar_url, created_at, updated_at
		FROM users
		WHERE id = $1`

	var user User
	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(
		&user.ID,
		&user.Email,
		&user.DisplayName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

type OTP struct {
	ID       string
	Email    string
	CodeHash string
	Attempts int
}

type OTPModel struct {
	DB *sql.DB
}

func (m *OTPModel) Insert(ctx context.Context, email string, codeHash string, expiresAt time.Time) error {
	_, err := m.DB.ExecContext(ctx, `
		INSERT INTO email_otps (email, code_hash, expires_at)
		VALUES ($1, $2, $3)`, email, codeHash, expiresAt)
	return err
}

func (m *OTPModel) LatestValid(ctx context.Context, email string) (*OTP, error) {
	stmt := `
		SELECT id, email, code_hash, attempts
		FROM email_otps
		WHERE email = $1 AND used_at IS NULL AND expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 1`

	var otp OTP
	err := m.DB.QueryRowContext(ctx, stmt, email).Scan(&otp.ID, &otp.Email, &otp.CodeHash, &otp.Attempts)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &otp, nil
}

func (m *OTPModel) IncrementAttempts(ctx context.Context, id string) error {
	_, err := m.DB.ExecContext(ctx, `UPDATE email_otps SET attempts = attempts + 1 WHERE id = $1`, id)
	return err
}

func (m *OTPModel) MarkUsed(ctx context.Context, id string) error {
	_, err := m.DB.ExecContext(ctx, `UPDATE email_otps SET used_at = NOW() WHERE id = $1`, id)
	return err
}

type Event struct {
	ID           string    `json:"id"`
	HostID       string    `json:"host_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartsAt     time.Time `json:"starts_at"`
	Timezone     string    `json:"timezone"`
	LocationName string    `json:"location_name"`
	LocationNote string    `json:"location_note"`
	ShareCode    string    `json:"share_code"`
	Status       string    `json:"status"`
	Theme        string    `json:"theme"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EventStats struct {
	Attending      int `json:"attending"`
	Maybe          int `json:"maybe"`
	NotAttending   int `json:"not_attending"`
	PartySizeTotal int `json:"party_size_total"`
	ItemsNeeded    int `json:"items_needed"`
	ItemsClaimed   int `json:"items_claimed"`
	ItemsRemaining int `json:"items_remaining"`
	ClaimsCount    int `json:"claims_count"`
}

type EventSummary struct {
	Event
	Stats EventStats `json:"stats"`
}

type ItemInput struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	NeededQty int    `json:"needed_qty"`
	Unit      string `json:"unit"`
	Priority  string `json:"priority"`
	Notes     string `json:"notes"`
	SortOrder int    `json:"sort_order"`
}

type EventCreateInput struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	StartsAt     time.Time   `json:"starts_at"`
	Timezone     string      `json:"timezone"`
	LocationName string      `json:"location_name"`
	LocationNote string      `json:"location_note"`
	Items        []ItemInput `json:"items"`
}

type EventItem struct {
	ID          string    `json:"id"`
	EventID     string    `json:"event_id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	NeededQty   int       `json:"needed_qty"`
	ClaimedQty  int       `json:"claimed_qty"`
	Unit        string    `json:"unit"`
	Priority    string    `json:"priority"`
	Notes       string    `json:"notes"`
	SortOrder   int       `json:"sort_order"`
	IsPurchased bool      `json:"is_purchased"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Guest struct {
	ID           string    `json:"id"`
	EventID      string    `json:"event_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	RSVPStatus   string    `json:"rsvp_status"`
	PartySize    int       `json:"party_size"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	SessionToken string    `json:"session_token,omitempty"`
}

type GuestInput struct {
	SessionToken string `json:"session_token"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	RSVPStatus   string `json:"rsvp_status"`
	PartySize    int    `json:"party_size"`
	Note         string `json:"note"`
}

type ItemClaim struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	ItemID    string    `json:"item_id"`
	GuestID   string    `json:"guest_id"`
	ItemName  string    `json:"item_name"`
	GuestName string    `json:"guest_name"`
	Quantity  int       `json:"quantity"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClaimInput struct {
	SessionToken string `json:"session_token"`
	ItemID       string `json:"item_id"`
	Quantity     int    `json:"quantity"`
	Note         string `json:"note"`
}

type Activity struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	ActorName string    `json:"actor_name"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Notification struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type EventSnapshot struct {
	Event      *Event       `json:"event"`
	Stats      EventStats   `json:"stats"`
	Items      []*EventItem `json:"items"`
	Guests     []*Guest     `json:"guests"`
	Claims     []*ItemClaim `json:"claims"`
	Activities []*Activity  `json:"activities"`
}

type ChangeNotification struct {
	EventTitle string
	Title      string
	Message    string
	HostEmails []string
}

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) Create(ctx context.Context, hostID string, shareCode string, input EventCreateInput) (*Event, error) {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	timezone := input.Timezone
	if timezone == "" {
		timezone = "Europe/Berlin"
	}
	startsAt := input.StartsAt
	if startsAt.IsZero() {
		startsAt = time.Now().Add(72 * time.Hour)
	}

	stmt := `
		INSERT INTO events (host_id, title, description, starts_at, timezone, location_name, location_note, share_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, host_id, title, description, starts_at, timezone, location_name, location_note, share_code, status, theme, created_at, updated_at`

	var event Event
	err = tx.QueryRowContext(ctx, stmt,
		hostID,
		input.Title,
		input.Description,
		startsAt,
		timezone,
		input.LocationName,
		input.LocationNote,
		shareCode,
	).Scan(
		&event.ID,
		&event.HostID,
		&event.Title,
		&event.Description,
		&event.StartsAt,
		&event.Timezone,
		&event.LocationName,
		&event.LocationNote,
		&event.ShareCode,
		&event.Status,
		&event.Theme,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO event_hosts (event_id, user_id, role)
		VALUES ($1, $2, 'owner')`, event.ID, hostID)
	if err != nil {
		return nil, err
	}

	if len(input.Items) == 0 {
		input.Items = defaultItems()
	}

	for index, item := range input.Items {
		if item.Name == "" {
			continue
		}
		if item.NeededQty <= 0 {
			item.NeededQty = 1
		}
		if item.Category == "" {
			item.Category = "Other"
		}
		if item.Priority == "" {
			item.Priority = "normal"
		}
		if item.SortOrder == 0 {
			item.SortOrder = index
		}
		_, err = tx.ExecContext(ctx, `
			INSERT INTO event_items (event_id, name, category, needed_qty, unit, priority, notes, sort_order)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			event.ID, item.Name, item.Category, item.NeededQty, item.Unit, item.Priority, item.Notes, item.SortOrder,
		)
		if err != nil {
			return nil, err
		}
	}

	_, _ = tx.ExecContext(ctx, `
		INSERT INTO activity_events (event_id, actor_name, type, message)
		VALUES ($1, 'BringIt', 'event_created', 'Event created and ready to share')`, event.ID)

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &event, nil
}

func defaultItems() []ItemInput {
	return []ItemInput{
		{Name: "Grill meat", Category: "Grill", NeededQty: 4, Unit: "packs", Priority: "high"},
		{Name: "Jollof rice", Category: "Sides", NeededQty: 2, Unit: "trays", Priority: "high"},
		{Name: "Soft drinks", Category: "Drinks", NeededQty: 8, Unit: "bottles", Priority: "normal"},
		{Name: "Water", Category: "Drinks", NeededQty: 12, Unit: "bottles", Priority: "normal"},
		{Name: "Charcoal", Category: "Supplies", NeededQty: 2, Unit: "bags", Priority: "high"},
		{Name: "Disposable plates", Category: "Supplies", NeededQty: 1, Unit: "pack", Priority: "normal"},
	}
}

func (m *EventModel) ListByHost(ctx context.Context, hostID string) ([]*EventSummary, error) {
	rows, err := m.DB.QueryContext(ctx, `
		SELECT e.id, e.host_id, e.title, e.description, e.starts_at, e.timezone, e.location_name, e.location_note,
		       e.share_code, e.status, e.theme, e.created_at, e.updated_at,
		       COUNT(g.id) FILTER (WHERE g.rsvp_status = 'attending')::int AS attending,
		       COUNT(g.id) FILTER (WHERE g.rsvp_status = 'maybe')::int AS maybe,
		       COUNT(g.id) FILTER (WHERE g.rsvp_status = 'not_attending')::int AS not_attending,
		       COALESCE(SUM(g.party_size) FILTER (WHERE g.rsvp_status IN ('attending', 'maybe')), 0)::int AS party_size_total,
		       COALESCE((SELECT SUM(needed_qty)::int FROM event_items WHERE event_id = e.id), 0) AS items_needed,
		       COALESCE((SELECT SUM(claimed_qty)::int FROM event_items WHERE event_id = e.id), 0) AS items_claimed,
		       COALESCE((SELECT SUM(needed_qty - claimed_qty)::int FROM event_items WHERE event_id = e.id), 0) AS items_remaining,
		       COALESCE((SELECT COUNT(*)::int FROM item_claims WHERE event_id = e.id AND status = 'active'), 0) AS claims_count
		FROM events e
		JOIN event_hosts eh ON eh.event_id = e.id
		LEFT JOIN event_guests g ON g.event_id = e.id
		WHERE eh.user_id = $1
		GROUP BY e.id
		ORDER BY e.starts_at ASC`, hostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*EventSummary
	for rows.Next() {
		var summary EventSummary
		if err := rows.Scan(
			&summary.ID,
			&summary.HostID,
			&summary.Title,
			&summary.Description,
			&summary.StartsAt,
			&summary.Timezone,
			&summary.LocationName,
			&summary.LocationNote,
			&summary.ShareCode,
			&summary.Status,
			&summary.Theme,
			&summary.CreatedAt,
			&summary.UpdatedAt,
			&summary.Stats.Attending,
			&summary.Stats.Maybe,
			&summary.Stats.NotAttending,
			&summary.Stats.PartySizeTotal,
			&summary.Stats.ItemsNeeded,
			&summary.Stats.ItemsClaimed,
			&summary.Stats.ItemsRemaining,
			&summary.Stats.ClaimsCount,
		); err != nil {
			return nil, err
		}
		events = append(events, &summary)
	}
	return events, rows.Err()
}

func (m *EventModel) GetForHost(ctx context.Context, id string, hostID string) (*Event, error) {
	var event Event
	err := m.DB.QueryRowContext(ctx, `
		SELECT e.id, e.host_id, e.title, e.description, e.starts_at, e.timezone, e.location_name, e.location_note,
		       e.share_code, e.status, e.theme, e.created_at, e.updated_at
		FROM events e
		JOIN event_hosts eh ON eh.event_id = e.id
		WHERE e.id = $1 AND eh.user_id = $2`, id, hostID).Scan(
		&event.ID,
		&event.HostID,
		&event.Title,
		&event.Description,
		&event.StartsAt,
		&event.Timezone,
		&event.LocationName,
		&event.LocationNote,
		&event.ShareCode,
		&event.Status,
		&event.Theme,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &event, nil
}

func (m *EventModel) GetByShareCode(ctx context.Context, shareCode string) (*Event, error) {
	var event Event
	err := m.DB.QueryRowContext(ctx, `
		SELECT id, host_id, title, description, starts_at, timezone, location_name, location_note,
		       share_code, status, theme, created_at, updated_at
		FROM events
		WHERE share_code = $1 AND status = 'published'`, shareCode).Scan(
		&event.ID,
		&event.HostID,
		&event.Title,
		&event.Description,
		&event.StartsAt,
		&event.Timezone,
		&event.LocationName,
		&event.LocationNote,
		&event.ShareCode,
		&event.Status,
		&event.Theme,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &event, nil
}

func (m *EventModel) Snapshot(ctx context.Context, eventID string) (*EventSnapshot, error) {
	event, err := m.getByID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	items, err := m.ListItems(ctx, eventID)
	if err != nil {
		return nil, err
	}
	guests, err := m.ListGuests(ctx, eventID)
	if err != nil {
		return nil, err
	}
	claims, err := m.ListClaims(ctx, eventID)
	if err != nil {
		return nil, err
	}
	activities, err := m.ListActivities(ctx, eventID)
	if err != nil {
		return nil, err
	}
	stats, err := m.Stats(ctx, eventID)
	if err != nil {
		return nil, err
	}

	return &EventSnapshot{
		Event:      event,
		Stats:      stats,
		Items:      items,
		Guests:     guests,
		Claims:     claims,
		Activities: activities,
	}, nil
}

func (m *EventModel) getByID(ctx context.Context, id string) (*Event, error) {
	var event Event
	err := m.DB.QueryRowContext(ctx, `
		SELECT id, host_id, title, description, starts_at, timezone, location_name, location_note,
		       share_code, status, theme, created_at, updated_at
		FROM events
		WHERE id = $1`, id).Scan(
		&event.ID,
		&event.HostID,
		&event.Title,
		&event.Description,
		&event.StartsAt,
		&event.Timezone,
		&event.LocationName,
		&event.LocationNote,
		&event.ShareCode,
		&event.Status,
		&event.Theme,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &event, nil
}

func (m *EventModel) Stats(ctx context.Context, eventID string) (EventStats, error) {
	var stats EventStats
	err := m.DB.QueryRowContext(ctx, `
		SELECT COUNT(g.id) FILTER (WHERE g.rsvp_status = 'attending')::int AS attending,
		       COUNT(g.id) FILTER (WHERE g.rsvp_status = 'maybe')::int AS maybe,
		       COUNT(g.id) FILTER (WHERE g.rsvp_status = 'not_attending')::int AS not_attending,
		       COALESCE(SUM(g.party_size) FILTER (WHERE g.rsvp_status IN ('attending', 'maybe')), 0)::int AS party_size_total,
		       COALESCE((SELECT SUM(needed_qty)::int FROM event_items WHERE event_id = $1), 0) AS items_needed,
		       COALESCE((SELECT SUM(claimed_qty)::int FROM event_items WHERE event_id = $1), 0) AS items_claimed,
		       COALESCE((SELECT SUM(needed_qty - claimed_qty)::int FROM event_items WHERE event_id = $1), 0) AS items_remaining,
		       COALESCE((SELECT COUNT(*)::int FROM item_claims WHERE event_id = $1 AND status = 'active'), 0) AS claims_count
		FROM event_guests g
		WHERE g.event_id = $1`, eventID).Scan(
		&stats.Attending,
		&stats.Maybe,
		&stats.NotAttending,
		&stats.PartySizeTotal,
		&stats.ItemsNeeded,
		&stats.ItemsClaimed,
		&stats.ItemsRemaining,
		&stats.ClaimsCount,
	)
	return stats, err
}

func (m *EventModel) ListItems(ctx context.Context, eventID string) ([]*EventItem, error) {
	rows, err := m.DB.QueryContext(ctx, `
		SELECT id, event_id, name, category, needed_qty, claimed_qty, unit, priority, notes, sort_order, is_purchased, created_at, updated_at
		FROM event_items
		WHERE event_id = $1
		ORDER BY category, sort_order, created_at`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*EventItem
	for rows.Next() {
		var item EventItem
		if err := rows.Scan(
			&item.ID,
			&item.EventID,
			&item.Name,
			&item.Category,
			&item.NeededQty,
			&item.ClaimedQty,
			&item.Unit,
			&item.Priority,
			&item.Notes,
			&item.SortOrder,
			&item.IsPurchased,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, rows.Err()
}

func (m *EventModel) ListGuests(ctx context.Context, eventID string) ([]*Guest, error) {
	rows, err := m.DB.QueryContext(ctx, `
		SELECT id, event_id, name, email, phone, rsvp_status, party_size, note, created_at, updated_at
		FROM event_guests
		WHERE event_id = $1
		ORDER BY updated_at DESC`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []*Guest
	for rows.Next() {
		var guest Guest
		if err := rows.Scan(
			&guest.ID,
			&guest.EventID,
			&guest.Name,
			&guest.Email,
			&guest.Phone,
			&guest.RSVPStatus,
			&guest.PartySize,
			&guest.Note,
			&guest.CreatedAt,
			&guest.UpdatedAt,
		); err != nil {
			return nil, err
		}
		guests = append(guests, &guest)
	}
	return guests, rows.Err()
}

func (m *EventModel) ListClaims(ctx context.Context, eventID string) ([]*ItemClaim, error) {
	rows, err := m.DB.QueryContext(ctx, `
		SELECT c.id, c.event_id, c.item_id, c.guest_id, i.name, g.name, c.quantity, c.note, c.status, c.created_at, c.updated_at
		FROM item_claims c
		JOIN event_items i ON i.id = c.item_id
		JOIN event_guests g ON g.id = c.guest_id
		WHERE c.event_id = $1 AND c.status = 'active'
		ORDER BY c.created_at DESC`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var claims []*ItemClaim
	for rows.Next() {
		var claim ItemClaim
		if err := rows.Scan(
			&claim.ID,
			&claim.EventID,
			&claim.ItemID,
			&claim.GuestID,
			&claim.ItemName,
			&claim.GuestName,
			&claim.Quantity,
			&claim.Note,
			&claim.Status,
			&claim.CreatedAt,
			&claim.UpdatedAt,
		); err != nil {
			return nil, err
		}
		claims = append(claims, &claim)
	}
	return claims, rows.Err()
}

func (m *EventModel) ListActivities(ctx context.Context, eventID string) ([]*Activity, error) {
	rows, err := m.DB.QueryContext(ctx, `
		SELECT id, event_id, actor_name, type, message, created_at
		FROM activity_events
		WHERE event_id = $1
		ORDER BY created_at DESC
		LIMIT 40`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*Activity
	for rows.Next() {
		var activity Activity
		if err := rows.Scan(
			&activity.ID,
			&activity.EventID,
			&activity.ActorName,
			&activity.Type,
			&activity.Message,
			&activity.CreatedAt,
		); err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, rows.Err()
}

func (m *EventModel) UpsertGuest(ctx context.Context, eventID string, token string, input GuestInput) (*Guest, *ChangeNotification, error) {
	if token == "" {
		return nil, nil, ErrGuestSessionRequired
	}
	if input.PartySize <= 0 {
		input.PartySize = 1
	}
	if input.RSVPStatus == "" {
		input.RSVPStatus = "attending"
	}

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	eventTitle, hostEmails, err := hostNotificationContext(ctx, tx, eventID)
	if err != nil {
		return nil, nil, err
	}

	var previousStatus string
	_ = tx.QueryRowContext(ctx, `
		SELECT rsvp_status
		FROM event_guests
		WHERE event_id = $1 AND session_token = $2`, eventID, token).Scan(&previousStatus)

	stmt := `
		INSERT INTO event_guests (event_id, session_token, name, email, phone, rsvp_status, party_size, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (event_id, session_token) DO UPDATE SET
			name = EXCLUDED.name,
			email = EXCLUDED.email,
			phone = EXCLUDED.phone,
			rsvp_status = EXCLUDED.rsvp_status,
			party_size = EXCLUDED.party_size,
			note = EXCLUDED.note
		RETURNING id, event_id, name, email, phone, rsvp_status, party_size, note, created_at, updated_at`

	var guest Guest
	err = tx.QueryRowContext(ctx, stmt,
		eventID,
		token,
		input.Name,
		input.Email,
		input.Phone,
		input.RSVPStatus,
		input.PartySize,
		input.Note,
	).Scan(
		&guest.ID,
		&guest.EventID,
		&guest.Name,
		&guest.Email,
		&guest.Phone,
		&guest.RSVPStatus,
		&guest.PartySize,
		&guest.Note,
		&guest.CreatedAt,
		&guest.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}
	guest.SessionToken = token

	action := "joined"
	if previousStatus != "" {
		action = "updated RSVP"
	}
	message := fmt.Sprintf("%s %s: %s", guest.Name, action, guest.RSVPStatus)
	if guest.PartySize > 1 {
		message = fmt.Sprintf("%s with %d people", message, guest.PartySize)
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO activity_events (event_id, actor_name, type, message)
		VALUES ($1, $2, 'rsvp_updated', $3)`, eventID, guest.Name, message); err != nil {
		return nil, nil, err
	}

	title := "RSVP updated"
	notificationMessage := message
	if err := insertHostNotifications(ctx, tx, eventID, title, notificationMessage, "rsvp"); err != nil {
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &guest, &ChangeNotification{
		EventTitle: eventTitle,
		Title:      title,
		Message:    notificationMessage,
		HostEmails: hostEmails,
	}, nil
}

func (m *EventModel) CreateClaim(ctx context.Context, eventID string, token string, input ClaimInput) (*ItemClaim, *ChangeNotification, error) {
	if token == "" {
		return nil, nil, ErrGuestSessionRequired
	}
	if input.Quantity <= 0 {
		input.Quantity = 1
	}

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	eventTitle, hostEmails, err := hostNotificationContext(ctx, tx, eventID)
	if err != nil {
		return nil, nil, err
	}

	var guestID, guestName string
	err = tx.QueryRowContext(ctx, `
		SELECT id, name
		FROM event_guests
		WHERE event_id = $1 AND session_token = $2`, eventID, token).Scan(&guestID, &guestName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrGuestSessionRequired
		}
		return nil, nil, err
	}

	var itemName string
	var neededQty, claimedQty int
	err = tx.QueryRowContext(ctx, `
		SELECT name, needed_qty, claimed_qty
		FROM event_items
		WHERE id = $1 AND event_id = $2
		FOR UPDATE`, input.ItemID, eventID).Scan(&itemName, &neededQty, &claimedQty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrRecordNotFound
		}
		return nil, nil, err
	}

	if claimedQty+input.Quantity > neededQty {
		return nil, nil, ErrQuantityUnavailable
	}

	var claim ItemClaim
	err = tx.QueryRowContext(ctx, `
		INSERT INTO item_claims (event_id, item_id, guest_id, quantity, note)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, event_id, item_id, guest_id, quantity, note, status, created_at, updated_at`,
		eventID,
		input.ItemID,
		guestID,
		input.Quantity,
		input.Note,
	).Scan(
		&claim.ID,
		&claim.EventID,
		&claim.ItemID,
		&claim.GuestID,
		&claim.Quantity,
		&claim.Note,
		&claim.Status,
		&claim.CreatedAt,
		&claim.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}
	claim.ItemName = itemName
	claim.GuestName = guestName

	if _, err := tx.ExecContext(ctx, `
		UPDATE event_items
		SET claimed_qty = claimed_qty + $1
		WHERE id = $2`, input.Quantity, input.ItemID); err != nil {
		return nil, nil, err
	}

	message := fmt.Sprintf("%s is bringing %d %s", guestName, input.Quantity, itemName)
	if _, err := tx.ExecContext(ctx, `
		INSERT INTO activity_events (event_id, actor_name, type, message)
		VALUES ($1, $2, 'item_claimed', $3)`, eventID, guestName, message); err != nil {
		return nil, nil, err
	}

	title := "Item claimed"
	if err := insertHostNotifications(ctx, tx, eventID, title, message, "claim"); err != nil {
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return &claim, &ChangeNotification{
		EventTitle: eventTitle,
		Title:      title,
		Message:    message,
		HostEmails: hostEmails,
	}, nil
}

func (m *EventModel) CancelClaim(ctx context.Context, eventID string, token string, claimID string) (*ChangeNotification, error) {
	if token == "" {
		return nil, ErrGuestSessionRequired
	}

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	eventTitle, hostEmails, err := hostNotificationContext(ctx, tx, eventID)
	if err != nil {
		return nil, err
	}

	var itemID, itemName, guestName string
	var quantity int
	err = tx.QueryRowContext(ctx, `
		SELECT c.item_id, i.name, g.name, c.quantity
		FROM item_claims c
		JOIN event_items i ON i.id = c.item_id
		JOIN event_guests g ON g.id = c.guest_id
		WHERE c.id = $1 AND c.event_id = $2 AND g.session_token = $3 AND c.status = 'active'
		FOR UPDATE OF c`, claimID, eventID, token).Scan(&itemID, &itemName, &guestName, &quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, `
		SELECT id
		FROM event_items
		WHERE id = $1
		FOR UPDATE`, itemID); err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE item_claims
		SET status = 'cancelled'
		WHERE id = $1`, claimID); err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE event_items
		SET claimed_qty = GREATEST(0, claimed_qty - $1)
		WHERE id = $2`, quantity, itemID); err != nil {
		return nil, err
	}

	message := fmt.Sprintf("%s removed %d %s from their list", guestName, quantity, itemName)
	if _, err := tx.ExecContext(ctx, `
		INSERT INTO activity_events (event_id, actor_name, type, message)
		VALUES ($1, $2, 'claim_cancelled', $3)`, eventID, guestName, message); err != nil {
		return nil, err
	}

	title := "Item claim removed"
	if err := insertHostNotifications(ctx, tx, eventID, title, message, "claim_cancelled"); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &ChangeNotification{
		EventTitle: eventTitle,
		Title:      title,
		Message:    message,
		HostEmails: hostEmails,
	}, nil
}

func (m *EventModel) AddItem(ctx context.Context, eventID string, hostID string, input ItemInput) (*EventItem, error) {
	if err := m.requireHost(ctx, eventID, hostID); err != nil {
		return nil, err
	}
	if input.NeededQty <= 0 {
		input.NeededQty = 1
	}
	if input.Category == "" {
		input.Category = "Other"
	}
	if input.Priority == "" {
		input.Priority = "normal"
	}

	var item EventItem
	err := m.DB.QueryRowContext(ctx, `
		INSERT INTO event_items (event_id, name, category, needed_qty, unit, priority, notes, sort_order)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, event_id, name, category, needed_qty, claimed_qty, unit, priority, notes, sort_order, is_purchased, created_at, updated_at`,
		eventID,
		input.Name,
		input.Category,
		input.NeededQty,
		input.Unit,
		input.Priority,
		input.Notes,
		input.SortOrder,
	).Scan(
		&item.ID,
		&item.EventID,
		&item.Name,
		&item.Category,
		&item.NeededQty,
		&item.ClaimedQty,
		&item.Unit,
		&item.Priority,
		&item.Notes,
		&item.SortOrder,
		&item.IsPurchased,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (m *EventModel) requireHost(ctx context.Context, eventID string, hostID string) error {
	var exists bool
	err := m.DB.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM event_hosts
			WHERE event_id = $1 AND user_id = $2
		)`, eventID, hostID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return ErrNotEventHost
	}
	return nil
}

type NotificationModel struct {
	DB *sql.DB
}

func (m *NotificationModel) ListForUser(ctx context.Context, userID string, eventID string) ([]*Notification, error) {
	args := []any{userID}
	where := "WHERE user_id = $1"
	if eventID != "" {
		where += " AND event_id = $2"
		args = append(args, eventID)
	}

	rows, err := m.DB.QueryContext(ctx, `
		SELECT id, event_id, user_id, title, message, type, is_read, created_at
		FROM notifications
		`+where+`
		ORDER BY created_at DESC
		LIMIT 80`, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		var notification Notification
		if err := rows.Scan(
			&notification.ID,
			&notification.EventID,
			&notification.UserID,
			&notification.Title,
			&notification.Message,
			&notification.Type,
			&notification.IsRead,
			&notification.CreatedAt,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}
	return notifications, rows.Err()
}

func (m *NotificationModel) MarkRead(ctx context.Context, userID string, eventID string) error {
	args := []any{userID}
	where := "user_id = $1"
	if eventID != "" {
		where += " AND event_id = $2"
		args = append(args, eventID)
	}

	_, err := m.DB.ExecContext(ctx, `UPDATE notifications SET is_read = TRUE WHERE `+where, args...)
	return err
}

func hostNotificationContext(ctx context.Context, tx *sql.Tx, eventID string) (string, []string, error) {
	var eventTitle string
	err := tx.QueryRowContext(ctx, `SELECT title FROM events WHERE id = $1`, eventID).Scan(&eventTitle)
	if err != nil {
		return "", nil, err
	}

	rows, err := tx.QueryContext(ctx, `
		SELECT u.email
		FROM event_hosts eh
		JOIN users u ON u.id = eh.user_id
		WHERE eh.event_id = $1 AND eh.notify_email = TRUE`, eventID)
	if err != nil {
		return "", nil, err
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return "", nil, err
		}
		emails = append(emails, email)
	}
	return eventTitle, emails, rows.Err()
}

func insertHostNotifications(ctx context.Context, tx *sql.Tx, eventID string, title string, message string, notificationType string) error {
	_, err := tx.ExecContext(ctx, `
		INSERT INTO notifications (event_id, user_id, title, message, type)
		SELECT $1, user_id, $2, $3, $4
		FROM event_hosts
		WHERE event_id = $1`, eventID, title, message, notificationType)
	return err
}
