REM @echo off
echo update...
git pull origin master
git log -1 --oneline
echo build...
go build -o ./bin/SERVER.exe quiz/cmd/server
go build -o ./bin/setup.exe quiz/cmd/setup
go build -o ./bin/dump-db.exe quiz/cmd/exportdb
go build -o ./bin/import-db.exe quiz/cmd/importdb
echo:
echo Done!
pause
