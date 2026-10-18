# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Funda Africa is an African ed-tech platform split into **four independently-runnable services** in one repo. They are not yet wired together end-to-end — each is at a different maturity level, and the cross-service integration is largely pending (see "Integration status"). Treat each directory as its own project with its own toolchain.

| Directory | Stack | Role | Default port |
|---|---|---|---|
| `gin-funda-africa/` | Go 1.23 + Gin + GORM + Casbin | Core REST API: users/auth, subscriptions, education content (curriculum→subject→topic→sub-topic), organizations. Owns authN/authZ and issues JWTs. | `8080` (via `PORT`) |
| `funda_africa/` | Elixir ~1.14 + Phoenix 1.7 | Realtime backend: chat, live lessons (WebRTC signaling), AI-conversation persistence. Verifies JWTs (does not issue them). | `4000` |
| `funda_ai/` | Python + FastAPI | AI tutoring gateway; routes prompts to Google Gemini and Wolfram Alpha, persists interactions back to Phoenix. | `8000` (uvicorn default) |
| `funda-africa-frontend/` | Next.js 15 + React 19 | Student learning UI (subjects, live lessons, progress, AI tutor). Currently a mock-data prototype. | `3000` |

All HTTP APIs use base path `/api/v1`. Postgres is the database for both Go and Elixir services.

## Commands

### gin-funda-africa (Go)
```bash
cd gin-funda-africa
go mod download                               # install deps
air                                           # hot-reload dev server (.air.toml)
go run ./cmd/app                              # run without hot reload
go build -o ./tmp/main ./cmd/app              # build
go test ./...                                 # tests (none exist yet)
go test ./internal/modules/users/... -run TestName   # single package/test
swag init -g cmd/app/main.go -o internal/docs # regenerate swagger (annotations in main.go + controllers)
```
Config is read by viper from a **`dev.env`** file (NOT `.env`) in the working dir when `GIN_MODE` is empty/`development`. `*.env` is gitignored, so create `dev.env` locally. Swagger UI at `/swagger/*any`. `.air.toml` was authored on Windows (`tmp\main.exe`) — adjust paths on Linux if needed.

### funda_africa (Phoenix)
```bash
cd funda_africa
mix setup                 # deps.get + ecto.setup + assets.setup + assets.build
mix phx.server            # run (or: iex -S mix phx.server)
mix ecto.reset            # drop + recreate + migrate + seed
mix test                  # auto-creates/migrates test DB first
mix test test/path/to/file_test.exs:42   # single test at line 42
mix format
```
Dev binds `127.0.0.1:4000`. Swagger UI at `/api/swagger`; LiveDashboard at `/dev/dashboard` (dev only). Web server is **Bandit**, not Cowboy.

### funda_ai (FastAPI)
```bash
cd funda_ai
# NO requirements.txt/pyproject exists — install manually into a venv:
pip install fastapi uvicorn httpx pydantic pydantic-settings google-generativeai
uvicorn main:app --reload   # default port 8000; app factory also at main:create_app
```
Config via pydantic-settings from a `.env` file. No tests exist.

### funda-africa-frontend (Next.js)
```bash
cd funda-africa-frontend
npm install              # see dependency gap below
npm run dev              # next dev --turbopack, port 3000
npm run build            # next build --turbopack
npm run lint
```

## Architecture conventions

### Go service — layered modules
Structure is `internal/modules/<module>/` with `controllers → services → repository → models` and a `dto/` package. Three modules: `users`, `education`, `organizations`.
- **Repositories are package-level functions using a single global `users.Database *gorm.DB`** (created in `users/database.go`) — there is no dependency injection; services are empty structs that pass through to repositories.
- Routes register via `RegisterAll<Module>...Routes(router, enforcer, jwtKey)` called from `cmd/app/main.go`, which builds a `/api/v1/...` group, applies `AuthMiddleware` then `RequirePermission`, then binds per-resource sub-groups.
- Models embed `gorm.Model` with UUID PKs (`default:gen_random_uuid()`) — Postgres needs pgcrypto / pg13+. `AutoMigrate` runs for all models at startup.
- Model→DTO mapping uses `jinzhu/copier`; partial updates use `utils.StructToMap`. Pagination via `utils.GetPaginatedData` (reflects over `Params` pointer fields → GORM WHERE clauses).
- **RBAC (Casbin):** roles `super_admin` (wildcard), `admin`, `teacher`, `tutor`, `org_rep`, `student` with inheritance (`g,` lines). `obj` = route `FullPath()`, `act` = HTTP method — so route templates in `policy.csv` must match registered routes exactly (including trailing slashes). Policies seed from `internal/pkg/rbac/policy.csv` on first run, then the **DB is the source of truth** (gorm-adapter). Note: despite the filename `rbac_with_domains_model.conf`, it is a plain 3-arg model, not domain/tenant-scoped.
- Module dir naming is inconsistent: `education`/`organizations` use prefixed dirs (`education_models`, `organization_models`), `users` uses plain `models`/`routes`.

### Phoenix service — Ops pattern
Each entity under `lib/funda_africa/{chat,ai,live_lessons}/` is split into **two files**: a schema module (`use Ecto.Schema` + changeset, UUID `binary_id` PKs) and an **`*_ops.ex` module** holding the context functions (Repo CRUD/queries). This replaces the conventional single Phoenix context module. `lib/funda_africa/models.ex` defines `FundaAfrica.Model`, a facade that `defdelegate`s a flat API across all `*_ops` modules. `FundaAfrica.Repo` adds `use Scrivener, page_size: 20` for pagination.
- **Channels** (`user_socket.ex`): `chat:*` → ChatChannel, `ai:*` → AIChannel, `live_lesson:*` → LiveLessonChannel. Chat fans out via `Endpoint.broadcast` to per-user topics `chat:user:<id>` with per-recipient receipts; live lessons relay WebRTC signaling via per-peer topics `live_lesson:peer:<id>`.
- **Auth** (`auth/jwt.ex`, `plugs/auth.ex`): Joken HS256, reads token from cookie → `Authorization: Bearer` → socket `token` param, extracts `sub` as user id. The `:auth` pipeline guards all `/api/v1` routes. This service only verifies tokens minted by the Go service.

### FastAPI service — orchestrator pattern
`routers/ai.py` exposes `POST /ai/query` → `Orchestrator` (`services/orchestration.py`) → provider services. Startup wires one shared `httpx.AsyncClient` and the orchestrator onto `app.state` via a `lifespan` context manager. `BaseLLM` (`services/base_llm.py`) is the interchangeable provider contract (`generate()`), implemented by `GeminiService` and `WolframLLMService`; `WolframService` (Full Results XML API) is separate. `Orchestrator.handle_query` branches on `mode` (`hybrid` default | `llm` | `wolfram`), uses a STEM keyword heuristic to pick Wolfram-LLM vs Gemini, is ZIMSEC-4004 syllabus aware, and persists each interaction to Phoenix via `PhoenixClient` (failures are swallowed so they don't block answers).

### Frontend — App Router + Zustand
`src/app/` (all client components, no `api/` routes). Global state in `src/lib/store.ts` (Zustand + `persist` to localStorage key `funda-africa-store`). Domain content is a hardcoded array in `src/lib/data.ts`. Components split into `ui/` (shadcn/Radix), `layout/`, `content/`, `ai-tutor/`. Always import via the `@/*` → `src/*` alias.

## Integration status (important)

The services are **not yet connected**:
- Phoenix `AIChannel` / audio transcription return hardcoded simulated responses; `AI_GATEWAY_URL` (default `http://localhost:8001`) is configured but never read. No HTTP client to `funda_ai` or the Go service exists.
- `funda_ai`'s `PhoenixClient` *does* POST interactions to Phoenix (`PHOENIX_AI_INTERACTIONS_URL`).
- The frontend has **no** network/WebSocket/API code — all data is mock.

## Known breakages to be aware of

- **funda_ai startup is broken**: `main.py` calls `Orchestrator(wolfram=, llm=, phoenix=)` but `Orchestrator.__init__` expects `wolfram_alpha, wolfram_llm, gemini, phoenix`; `WolframLLMService` is never instantiated. Also `DEFAULT_LLM_MODEL` config (`gemini-2.5-flash`) is unused — `GeminiService` defaults to `gemini-pro`. Reconcile before expecting it to run.
- **Frontend dependency gap**: `package.json` declares only `react`/`react-dom`/`next`, but the code imports zustand, @radix-ui/*, katex, react-katex, recharts, lucide-react, react-syntax-highlighter, next-themes, cva, clsx, tailwind-merge. These must be added before it builds. Several import paths also mismatch actual filenames (`layout/siderbar.tsx`, `progres/`, `ui/seperator.tsx`) and some referenced components (`ui/scroll-area`, `ui/toast`) are missing. `subjects/[subjectId]/[topicId]/[subtopicId]/page.tsx` has corrupted lines near the top.
- **Phoenix**: `FundaAfricaWeb.Presence` is used by channels but not started in `application.ex`; `runtime.exs` configures `FundaRealtime.Finch` (wrong name — app starts `FundaAfrica.Finch`); `:jwt_secret` is only set in prod `runtime.exs` and read at compile time, so JWT verification is not configured in dev/test out of the box.

## Env vars

- **Go** (`dev.env`): `POSTGRES_HOST/USER/PASSWORD/DB/PORT`, `PORT`, `CLIENT_ORIGIN`, `JWT_SECRET_KEY`. DSN uses `sslmode=disable`, `TimeZone=Africa/Johannesburg`.
- **Phoenix** (dev): `POSTGRES_USER/PASSWORD/HOST/DATABASE` (dev DB default `funda_africa_phoenix`), `PORT`, `POOL_SIZE`. Prod/runtime: `DATABASE_URL`, `SECRET_KEY_BASE`, `JWT_SECRET_KEY`, `PHX_HOST`, `CLIENT_ORIGIN`, `AI_GATEWAY_URL`. Note: dev config commits a fallback DB password and secrets.
- **FastAPI** (`.env`): `WOLFRAM_APPID`, `GEMINI_API_KEY`, `PHOENIX_AI_INTERACTIONS_URL`, `PHOENIX_API_TOKEN`; optional `DEFAULT_LLM_MODEL`, `THINKING_BUDGET`.
- **Frontend**: none yet; backend base URLs (e.g. `NEXT_PUBLIC_API_URL`, Phoenix socket URL, AI URL) will need to be introduced when wiring integration.
