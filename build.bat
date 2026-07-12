@echo off
echo Building api-gateway...
go build -o api-gateway ./cmd/api-gateway
if %ERRORLEVEL% EQU 0 (
    echo Build successful
) else (
    echo Build failed
    exit /b 1
)
