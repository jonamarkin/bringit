# BringIt

BringIt is a mobile-first event coordination MVP for RSVPs and shared contribution lists.

## Stack

- Go API with `net/http`
- PostgreSQL
- Nuxt 3 PWA frontend with shadcn-vue primitives
- Docker Compose
- Resend-compatible email delivery

## Local Development

Create a local env file:

```sh
cp .env.example .env
```

Start the stack:

```sh
make dev
```

Local URLs:

- Frontend: `http://localhost:3000`
- API: `http://localhost:8080`
- API health: `http://localhost:8080/healthz`

The first login code is returned in the API response when `APP_ENV=development`.

## MVP Flows

- Host signs in with email OTP.
- Host creates an event and item checklist.
- Host shares the public guest link.
- Guest RSVPs without creating an account.
- Guest claims item quantities.
- Host receives in-app notifications and email notifications when configured.

## Production Notes

Update `Caddyfile`, `PUBLIC_BASE_URL`, `FRONTEND_ORIGIN`, `COOKIE_DOMAIN`, and `JWT_SECRET` before deploying.
