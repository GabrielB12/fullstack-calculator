# Full-Stack Calculator

React (TypeScript) frontend + Go REST API backend.

## Setup

### Backend
    cd backend
    go run main.go handlers.go calculator.go
Runs on http://localhost:8080

### Frontend
    cd frontend
    npm install
    npm run dev
Runs on http://localhost:5173

### Tests
    cd backend && go test ./... -v -cover
    cd frontend && npm run test

### Docker (optional)
    docker compose up --build

## API Examples

POST /api/add
    { "a": 2, "b": 3 } -> { "result": 5 }

POST /api/divide
    { "a": 10, "b": 0 } -> 400 { "error": "division by zero" }

POST /api/sqrt
    { "a": 9 } -> { "result": 3 }

Endpoints: /api/add, /api/subtract, /api/multiply, /api/divide,
/api/power, /api/sqrt, /api/percentage, /api/health

## Design Decisions
- Separated pure calculation logic (calculator.go) from HTTP handlers for testability.
- Used net/http standard library only — no framework needed for this scope.
- Frontend uses a single reusable Calculator component with a services/api.ts layer,
  decoupling UI from network logic.
- Errors from the backend (division by zero, negative sqrt, zero-base percentage)
  are returned as 400 with a JSON error message, surfaced directly in the UI.
- Vitest + React Testing Library chosen for frontend tests (fast, native ESM support with Vite).