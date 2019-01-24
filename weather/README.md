# Weather

*By Pawat Nakpiphatkul  [(More Info)](../README.md)*

## Installation

1. Install Golang (using go1.11.2).
1. Get the repository into `GOPATH` by typing command `go get github.com/guitarpawat/the-internship-2019`.
1. Go to `%GOPATH%\src\github.com\guitarpawat\the-internship-2019\weather` (Windows) or `$GOPATH/src/github.com/guitarpawat/the-internship-2019/weather`.
1. Use `go run main.go` to compile and run the program.

## Usage
```
go run main.go -help
Usage of C:\Users\HP\AppData\Local\Temp\go-build537904090\b001\exe\main.exe:
  -json string
        JSON file location, if not specified, output will use the same name as xml file
  -xml string
        XML file location (default "weather.xml")
```

### Example

* `go run main.go`: Format weather.xml to weather.json
* `go run main.go -json=test.json`: Format weather.xml to test.json
* `go run main.go -xml=test.xml`: Format test.xml to test.json
* `go run main.go -xml=test.xml -json=test2.json`: Format test.xml to test2.json

## References

* [Documentation: encoding/xml - The Go Programming Language](https://golang.org/src/encoding/xml/example_test.go)
* [Package os - The Go Programming Language](https://golang.org/pkg/os/)
* [Trim string's suffix or extension?](https://stackoverflow.com/questions/13027912/trim-strings-suffix-or-extension)