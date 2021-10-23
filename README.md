# Read and Write Goroutine
This project will show you how file read and write operation can be done asynchronous. The concurrent task can be archive by using `goroutine` and `channel`. let dive in! 🚀🚀🚀

---
## Prerequisite
- [Golang v1.16 or later](https://go.dev/)

## Get Started
### Run With Golang
1. you will need to install the dependencies that used in this package by running below command
```bash
go mod tidy
```
2. to run the app you will need to run it by running below command
```bash
go run main.go
```
The default digit lenght is 2048 if you want to specific the lenght use
```bash
go run main.go 30
```
3. to run test, simply run
```bash
go test ./...
```
or with coverage
```bash
go test -cover ./...
```
this will run the test in current directory and it's subdirectory
### Run With Docker
📌 This option you will need to install Docker in your machine first 📌
1. you will need to build the docker image, before you can run it
```bash
docker build -t read-digit .
```
2. you can run the image with
```bash
docker run --rm read-digit
```
or with digits count option
```bash
docker run --rm -e DIGITS_COUNT=30 read-digit
```
## Build The Binary
1. to build the app to binary file you can simply use below command
```bash
go build -o read-digit .
```
1. to run binary file
```bash
./read-digit
```
The default digit lenght is 2048 if you want to specific the lenght use
```bash
./read-digit 30
```
