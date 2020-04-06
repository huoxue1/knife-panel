go generate
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/mac_knife_panel main.go
CGO_ENABLED=1 GOOS=linux CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" GOARCH=amd64 go build -a -o target/linux_knife_panel main.go
