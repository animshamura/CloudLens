@echo off
echo Starting Docker services...
docker compose up --build -d
echo Services started at http://localhost:8080
