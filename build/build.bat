go test -timeout 30s -v ./...
if %errorlevel% equ 0 (
	echo Tests passed, compiling code...
	go build -tags editor -o ./bin/kaiju.exe -ldflags="-s -w" main.go
) else (
	echo Tests failed, skipping code compile
)