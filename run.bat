@echo off
echo Building and running api-gateway...
go build -o api-gateway ./cmd/api-gateway
if %ERRORLEVEL% EQU 0 (
    echo Starting server...
    api-gateway.exe
) else (
    echo Build failed
    exit /b 1
)
