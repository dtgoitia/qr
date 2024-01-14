# qr

Small CLI to generate a QR code and print it to your terminal.

## Usage

```shell
qr "hello world"
```

## Development

```shell
git clone ...
cd qr

# run project
go run . "sample text to appear in the QR"

# build release binary
go build -ldflags "-s -w" -o build/qr
```