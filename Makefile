BINARY=app

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

test: 
	@go test -v -cover -covermode=atomic ./...

unittest: 
	@go test -short  ./...

migrate: 
	@go run ./migrations/migrate.go

run: 
	@go run main.go

download:
	@go mod download

build:
	@go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

migrateup:
	@go run migrations/migrate.go up 

migratedown:
	@go run migrations/migrate.go down 

create-secret:
	@kubectl apply -f deployment/mysql-secret.yaml
	@kubectl apply -f deployment/redis-secret.yaml


deploy:
	@kubectl apply -f deployment/mysql-db-pv.yaml
	@kubectl apply -f deployment/mysql-db-pvc.yaml
	@kubectl apply -f deployment/mysql-db-deployment.yaml
	@kubectl apply -f deployment/mysql-db-service.yaml
	@echo "Mysql Successfully Initialized"
	@echo ""
	@kubectl apply -f deployment/redis-pv.yaml
	@kubectl apply -f deployment/redis-pvc.yaml
	@kubectl apply -f deployment/redis-deployment.yaml
	@kubectl apply -f deployment/redis-service.yaml
	@echo "Redis Successfully Initialized"
	@echo ""
	@kubectl apply -f deployment/app-mysql-deployment.yaml
	@kubectl apply -f deployment/app-mysql-service.yaml
	@echo "App Successfully Initialized"
	@echo ""


.PHONY: clean download unittest test run lint-prepare lint build
