# Mellow

Modern, full‑stack social network.

Backend (Go + SQLite + sessions + WebSocket) and Frontend (Next.js + Tailwind). Features include authentication, profiles, posts, comments, groups, messages, notifications, and basic real‑time updates.

## Features
- Authentication via secure cookies (no JWT)
- User profiles and privacy
- Posts and comments
- Groups (membership, invites, events)
- Private and group messages
- Notifications (follows, invites, group actions)
- WebSocket broadcast channels

## Tech Stack
- Backend: Go 1.22+, net/http, SQLite, Gorilla WebSocket
- Frontend: Next.js (App Router), React 18, Tailwind CSS
- Orchestration: Docker, Docker Compose

## Project Structure
- `backend/` — Go API + WebSocket
  - `controllers/` HTTP handlers per domain
  - `services/` interfaces + `servimpl/` implementations
  - `repositories/` interfaces + `repoimpl/` SQL (SQLite)
  - `database/migration/sqlite/` SQL migrations
  - `middlewares/` auth, CORS, logger
  - `utils/` responses, errors, helpers
  - `websocket/` rooms and broadcasting
- `frontend/` — Next.js application
  - `src/app/` routes, components and UI
  - `public/` assets
- `assets/`, `doc/`, `.github/`

## Getting Started

Prerequisites
- Go 1.22+
- Node.js 18+
- npm 9+

Backend
- Run migrations automatically and start server:
  - `cd backend && go run ./`
- Defaults: `PORT=3225`, `DB_PATH=backend/data/social.db`, `MIGRATIONS_PATH=backend/database/migration/sqlite`
- Format code: `cd backend && go fmt ./...`

Frontend
- Start dev server with backend proxy:
  - `cd frontend && BACKEND_ORIGIN=http://localhost:3225 npm run dev`
- Lint: `cd frontend && npm run lint`

## Docker
- Dev (hot reload): `docker compose --profile dev up --build`
- Prod (detached): `docker compose --profile prod up -d`

Volumes persist `backend/data/` and `backend/uploads/`.

## Configuration
- Environment file `.env` (not committed). Important variables:
  - `PORT` — Backend HTTP port (default 3225)
  - `DB_PATH` — SQLite path (default `backend/data/social.db`)
  - `MIGRATIONS_PATH` — SQL migrations directory
  - `COOKIE_*` — Cookie/session settings
- Frontend proxy: `/api/*` → `BACKEND_ORIGIN` (see `frontend/next.config.mjs`).

## Database & Migrations
- SQLite database lives at `DB_PATH`
- Migrations in `backend/database/migration/sqlite`
- Applied automatically at backend start

## API (High‑Level)
- Auth: login/logout, session, current user
- Users: profile, follow/unfollow, search
- Posts: CRUD, likes, comments
- Groups: CRUD, membership, events, invites
- Messages: private and group conversations
- Notifications: list, read/unread

## WebSocket
- Endpoint: `/ws/chat?room=<room-id>`
- Rooms include `private:<u1>:<u2>` and `group:<groupId>`
- Broadcast is used for lightweight real‑time updates

## Development Guidelines
- Go code formatted with `gofmt`; keep packages idiomatic
- React components under `frontend/src/app/components` (PascalCase filenames)
- Avoid unused code and stray console logs

## Testing
- Backend: `go test ./backend/...` (use in‑memory SQLite for repo/service tests)
- Frontend: `npm run lint`; if needed, Jest + React Testing Library colocated tests

## Troubleshooting
- Database permissions: ensure `DB_PATH` is writable
- CORS/dev proxy: confirm `BACKEND_ORIGIN` points to backend URL
- Route conflicts: only one `mux.Handle("/something/")` per prefix; dispatch inside

## License
This project is licensed under the MIT License. See `LICENSE` for details.
