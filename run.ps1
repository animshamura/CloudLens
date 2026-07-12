#!/usr/bin/env pwsh
# Run script
Write-Host "Building and running api-gateway..." -ForegroundColor Cyan
go build -o api-gateway ./cmd/api-gateway
if ($LASTEXITCODE -eq 0) {
    Write-Host "Starting server..." -ForegroundColor Green
    .\api-gateway
} else {
    Write-Host "Build failed" -ForegroundColor Red
    exit 1
}
