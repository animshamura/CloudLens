#!/usr/bin/env pwsh
# Build script
Write-Host "Building api-gateway..." -ForegroundColor Cyan
go build -o api-gateway ./cmd/api-gateway
if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Build successful" -ForegroundColor Green
} else {
    Write-Host "✗ Build failed" -ForegroundColor Red
    exit 1
}
