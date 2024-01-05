# Quiz

## Build examples

```
go build -o ./bin/SERVER.exe quiz/cmd/server
go build -o ./bin/setup.exe quiz/cmd/setup
go build -o ./bin/dump-db.exe quiz/cmd/exportdb
go build -o ./bin/import-db.exe quiz/cmd/exportdb
```

## Recursively update packages in any subdirectories

```
cd quiz
go get -u ./...
```
