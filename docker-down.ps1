#!/usr/bin/env pwsh
# Docker down script
Write-Host "Stopping Docker services..." -ForegroundColor Cyan
docker compose down -v
Write-Host "✓ Services stopped" -ForegroundColor Green
