# User Service

This is the User service

Generated with

```
micro new user
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```


# mac
 # cgo_enabled=0 goos=linux goarch=amd64 go build -o user *.go
# win
 # o build -o user .

# docker build -t user:latest .