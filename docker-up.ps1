#!/usr/bin/env pwsh
# Docker up script
Write-Host "Starting Docker services..." -ForegroundColor Cyan
docker compose up --build -d
Write-Host "✓ Services started" -ForegroundColor Green
Write-Host "API Gateway running at http://localhost:8080" -ForegroundColor Yellow
