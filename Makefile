help:
	echo "Use run command to start local aplication"
	echo "Use build command to build your aplication"
run:
	templ generate
	go run main.go
gen:
	templ generate
build:
	go build
build-all:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go