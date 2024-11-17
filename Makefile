run-python:
	cd backend-servers && python3 main.py
run-go:
	go run main.go
run-kong:
	cd kong-gateway && docker-compose up
