#!/bin/bash

echo "Compilando binarios..."

mkdir -p releases

APP_NAME="ia_tp2"

echo "Compilando para Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o releases/${APP_NAME}_windows_amd64.exe ./main.go

echo "Compilando para Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o releases/${APP_NAME}_linux_amd64 ./main.go

echo "Compilando para macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o releases/${APP_NAME}_macos_amd64 ./main.go

echo "Listo la compilación"
echo "Binarios generados en la carpeta 'releases/':"
ls -la releases/