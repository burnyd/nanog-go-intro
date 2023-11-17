GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64
GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe
GOOS=linux GOARCH=arm go build -o myapp-linux-arm
GOOS=linux GOARCH=arm GOARM=7 go build -o myapp-rpi