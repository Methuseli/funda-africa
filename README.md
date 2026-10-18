# Funda Africa

An African ed-tech platform built as four cooperating services in a single repository: a core REST API, a realtime backend, an AI tutoring gateway, and a web client. It targets African curricula (e.g. the Zimbabwe O-Level ZIMSEC-4004 syllabus) with structured learning content, live lessons, and an AI tutor.

> **Status:** early development. Each service runs on its own, but cross-service integration is still largely pending and several components have known breakages. See [CLAUDE.md](./CLAUDE.md) for the current integration status and a list of things to fix before running.

## Services

| Directory | Stack | Responsibility | Port |
|---|---|---|---|
| [`gin-funda-africa/`](./gin-funda-africa) | Go 1.23, Gin, GORM, Casbin, PostgreSQL | Core REST API — users & auth (issues JWTs), subscription plans, education content (curriculum → subject → topic → sub-topic), organizations. Role-based access control. | 8080 |
| [`funda_africa/`](./funda_africa) | Elixir 1.14, Phoenix 1.7, Ecto, PostgreSQL | Realtime backend — chat & messaging with receipts, live lessons (WebRTC signaling, raise-hand, in-lesson chat), AI-conversation history. Verifies JWTs over HTTP and WebSocket channels. | 4000 |
| [`funda_ai/`](./funda_ai) | Python, FastAPI | AI tutoring gateway — routes student prompts to Google Gemini and Wolfram Alpha (hybrid, LLM-only, or Wolfram-only modes) and returns step-by-step explanations. Persists interactions back to the Phoenix service. | 8000 |
| [`funda-africa-frontend/`](./funda-africa-frontend) | Next.js 15, React 19, Tailwind 4, Zustand | Student web UI — dashboard, subject drill-down, live lessons, progress analytics, floating AI tutor chat (LaTeX-aware). Currently a mock-data prototype. | 3000 |

All HTTP APIs are served under `/api/v1`. The Go and Phoenix services each use their own PostgreSQL database.

## Architecture at a glance

```
                 ┌─────────────────────────┐
                 │  funda-africa-frontend   │  Next.js (3000)
                 └───────────┬─────────────┘
          REST /api/v1       │        Phoenix channels (chat / live lessons)
        ┌────────────────────┼────────────────────────┐
        ▼                    ▼                          ▼
┌───────────────┐   ┌──────────────────┐      (WebRTC signaling,
│ gin-funda-    │   │  funda_africa     │       realtime chat)
│ africa (8080) │   │  Phoenix (4000)   │
│ auth · content│   │ realtime · AI     │◄── interaction history ──┐
│ orgs · JWTs   │   │ conversation log  │                          │
└───────┬───────┘   └──────────────────┘                          │
        │ JWT verified by Phoenix                                  │
        ▼                                                 ┌────────┴────────┐
   PostgreSQL              PostgreSQL                      │   funda_ai      │
                                                           │  FastAPI (8000) │
                                                           │ Gemini · Wolfram│
                                                           └─────────────────┘
```

- **Auth:** the Go service issues JWTs (HS256) and enforces Casbin RBAC (`super_admin`, `admin`, `teacher`, `tutor`, `org_rep`, `student`). The Phoenix service verifies the same tokens for its REST endpoints and WebSocket connections but does not mint them.
- **AI:** `funda_ai` answers tutoring queries and writes each interaction back to Phoenix for persistence.

## Getting started

Each service has its own setup. In brief (see [CLAUDE.md](./CLAUDE.md) for env vars, caveats, and known breakages):

```bash
# Core API (Go)
cd gin-funda-africa && go mod download && air        # needs a local dev.env

# Realtime backend (Phoenix)
cd funda_africa && mix setup && mix phx.server

# AI gateway (FastAPI) — no requirements file yet; install deps manually
cd funda_ai && uvicorn main:app --reload             # needs a .env with API keys

# Web client (Next.js)
cd funda-africa-frontend && npm install && npm run dev
```

You will need a running PostgreSQL instance for the Go and Phoenix services, plus Google Gemini and Wolfram Alpha API keys for the AI gateway.

## Documentation

- **API docs:** Swagger UI at `/swagger` (Go) and `/api/swagger` (Phoenix).
- **Working in this repo with Claude Code:** see [CLAUDE.md](./CLAUDE.md) for build/test commands, per-service architecture conventions, integration status, and the current list of known breakages.
