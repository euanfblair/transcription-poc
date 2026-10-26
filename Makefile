.PHONY: test test-backend test-frontend typecheck run-backend run-frontend

test: test-backend test-frontend typecheck

test-backend:
	cd backend && go vet ./... && go test ./...

test-frontend:
	cd frontend && npm test

typecheck:
	cd frontend && npm run typecheck

run-backend:
	cd backend && go run ./cmd/server

run-frontend:
	cd frontend && npm run dev
