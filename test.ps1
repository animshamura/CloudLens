#!/usr/bin/env pwsh
# Test script
Write-Host "Running tests..." -ForegroundColor Cyan
go test ./... -v
