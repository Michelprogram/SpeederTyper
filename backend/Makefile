run-frontend:
	cd frontend && bun run dev

install-dependencies:
	cd frontend && bun install

run-backend:
	go run cmd/main.go -env prod -port 3000

build:
	go build -o build/speeder-typer-webscoket cmd/main.go