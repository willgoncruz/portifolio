### Makefile to run commands

# Run backend server
run:
	docker-compose --file docker-compose.yml up --build

## Generate typescript types from graphql schema
## cd src/backend && npm run generate
generate:
	go run github.com/99designs/gqlgen generate
