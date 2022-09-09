run: sdkv2

helloworld:
	go run cmd/helloworld/main.go
getting-started:
	go run cmd/getting-started/main.go
sdkv2:
	go run cmd/sdkv2/main.go

tidy:
	go mod tidy
	go mod vendor