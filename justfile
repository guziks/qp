# justfile

build:
    go build -o bin/qp cmd/main.go

release:
    GOOS=darwin GOARCH=amd64 go build -o bin/mac-intel/qp ./cmd
    GOOS=darwin GOARCH=arm64 go build -o bin/mac-arm/qp ./cmd
    GOOS=windows GOARCH=amd64 go build -o bin/win-intel/qp.exe ./cmd
    GOOS=windows GOARCH=arm64 go build -o bin/win-arm/qp.exe ./cmd
    GOOS=linux GOARCH=amd64 go build -o bin/linux-intel/qp ./cmd
    GOOS=linux GOARCH=arm64 go build -o bin/linux-arm/qp ./cmd
    mkdir -p release
    zip -j release/qp-mac-intel.zip bin/mac-intel/qp
    zip -j release/qp-mac-arm.zip bin/mac-arm/qp
    zip -j release/qp-win-intel.zip bin/win-intel/qp.exe
    zip -j release/qp-win-arm.zip bin/win-arm/qp.exe
    zip -j release/qp-linux-intel.zip bin/linux-intel/qp
    zip -j release/qp-linux-arm.zip bin/linux-arm/qp

clean:
    rm -rf bin release
