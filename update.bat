@echo off
echo update...
git pull origin master
git log -1 --oneline
echo:
pause
