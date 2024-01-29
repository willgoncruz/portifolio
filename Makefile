### Makefile to run commands

# Run dev server
run:
	docker-compose --file docker-compose.yml up --build

static:
	go run static.go index.html

