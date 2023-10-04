client:
	@go build -o build/client game_client/main.go

server:
	@go build -o build/server game_server/main.go