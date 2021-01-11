docker run --rm -v /d/golang/go-micro-service:/d/golang/go-micro-service -w /d/golang/go-micro-service  micro/micro new user

# mac
cgo_enabled=0 goos=linux goarch=amd64 go build -o user *.go
# win
build -o user .

# docker build -t user:latest .

SET CGO_ENABLED=0 
SET GOOS=linux 
SET GOARCH=amd64 
go build main.go