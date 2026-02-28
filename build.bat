@echo off
echo Building port-radar...
go mod tidy
go build -ldflags="-s -w" -o port-radar.exe .
if %errorlevel% neq 0 (
    echo Build failed!
    pause
    exit /b 1
)
echo Build success!
echo.
echo Run: port-radar.exe
pause
