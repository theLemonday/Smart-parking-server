release:
	GOOS=windows GOARCH=amd64 go build -o build/iot_server.exe cmd/iot_server/*.go
	GOOS=windows GOARCH=amd64 go build -o build/test_server.exe cmd/test_server/*.go

bundle:
	zip -r build.zip build/