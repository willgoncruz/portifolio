### Makefile to run commands

# Run backend server
run:
	docker-compose --file docker-compose.yml up --build

static:
	go run static.go index.html

