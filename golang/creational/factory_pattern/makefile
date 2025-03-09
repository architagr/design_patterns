ServiceName := JobAPI

lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o $(ServiceName) main.go 
	zip main.zip $(ServiceName)
	rm -f $(ServiceName)

clear:
	rm -f $(ServiceName) main.zip

tidy:
	go mod tidy
