REM @echo off
echo update...
git pull origin master || echo ERROR && pause && exit /b
git log -1 --oneline
echo build...
go build -o ./bin/RUN.exe quiz/cmd/server || echo ERROR && pause && exit /b
go build -o ./bin/setup.exe quiz/cmd/setup || echo ERROR && pause && exit /b
echo:
echo Done!
pause
