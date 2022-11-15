server:
	go run main.go

docs:
	swag init -g httpserver/httpserver.go

.PHONY: server docs
