# justfile

build:
    go build -o bin/qp cmd/main.go

release:
    GOOS=darwin GOARCH=amd64 go build -o bin/mac-amd64/qp ./cmd
    GOOS=darwin GOARCH=arm64 go build -o bin/mac-arm64/qp ./cmd
    GOOS=windows GOARCH=amd64 go build -o bin/win-amd64/qp.exe ./cmd
    GOOS=windows GOARCH=arm64 go build -o bin/win-arm64/qp.exe ./cmd
    GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/qp ./cmd
    GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/qp ./cmd
    mkdir -p release
    zip -j release/qp-mac-amd64.zip bin/mac-amd64/qp
    zip -j release/qp-mac-arm64.zip bin/mac-arm64/qp
    zip -j release/qp-win-amd64.zip bin/win-amd64/qp.exe
    zip -j release/qp-win-arm64.zip bin/win-arm64/qp.exe
    zip -j release/qp-linux-amd64.zip bin/linux-amd64/qp
    zip -j release/qp-linux-arm64.zip bin/linux-arm64/qp

clean:
    rm -rf bin release
