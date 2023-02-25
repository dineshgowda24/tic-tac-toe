build:
	go build -o play

build-windows:
	GOOS=windows GOARCH=amd64 go build -o play.exe
