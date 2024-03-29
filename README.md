# Psychological tests

## Build examples

```sh
go build -o ./bin/RUN.exe quiz/cmd/server
go build -o ./bin/setup.exe quiz/cmd/setup
```

## run

```sh
./bin/RUN.exe

# setup
./bin/setup.exe
# db drop
./bin/setup.exe -drop
```

## Test examples

```sh
go test quiz/internal/*
```


## Recursively update packages in any subdirectories

```sh
cd quiz
go get -u ./...
```

## Dependencies

- golang >= 1.21.0
- sqlite3
- required for building: cgo, gcc

Example install gcc in mingw
```
pacman -S base-devel mingw-w64-x86_64-toolchain mingw-w64-x86_64-libtool
```
- MobileHotspot (https://github.com/vilnar/mobile-hotspot)
