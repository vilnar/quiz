# Quiz

## Build examples

```
go build -o ./bin/RUN.exe quiz/cmd/server
go build -o ./bin/setup.exe quiz/cmd/setup
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

- sqlite3
- golang >= 1.21.0
- MobileHotspot (https://github.com/vilnar/mobile-hotspot)
