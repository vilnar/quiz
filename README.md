# Quiz

## Build examples

```
go build -o ./bin/SERVER.exe quiz/cmd/server
go build -o ./bin/setup.exe quiz/cmd/setup
go build -o ./bin/dump-db.exe quiz/cmd/exportdb
go build -o ./bin/import-db.exe quiz/cmd/importdb
```

## Test examples

```
go test quiz/internal/*
```


## Recursively update packages in any subdirectories

```
cd quiz
go get -u ./...
```

## Dependencies

- MariaDB >= 11.2.1
- golang >= 1.21.0
- MobileHotspot (https://github.com/vilnar/mobile-hotspot)
