migrate-db:
	@echo "Running migrations"
	go run cmd/migrate/migrate.go

run-cqlsh:
	@echo "Running cqlsh"
	docker exec -it cassandra cqlsh

run-dockers:
	@echo "Running dockers"
	docker-compose up -d