GOOS=darwin GOARCH=amd64 go build -o mac_knife_panel cmd/server/main.go
GOOS=linux GOARCH=amd64 go build -o linux_knife_panel cmd/server/main.go
