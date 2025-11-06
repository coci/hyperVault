build :
	@go build -o ./bin/HyperVault

run : build
	@./bin/HyperVault

test :
	@go test ./... -v