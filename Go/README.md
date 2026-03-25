CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webapp main.go

ops run ./webapp -p 8080

ops run ./webapp -p 8080 -m 128

ops run ./webapp -p 8080 -v

ops run ./webapp -p 8080 -a

time ops run ./webapp -p 8080