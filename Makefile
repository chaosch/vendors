all:
	go test ./...
	env GOOS=linux GOARCH=386 go test ./...
	go test ./... -short -race
	go vet
