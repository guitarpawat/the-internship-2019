# Hangman

*By Pawat Nakpiphatkul  [(More Info)](../README.md)*

## Installation

1. Install Golang (using go1.11.2).
1. Get the repository into `GOPATH` by typing command `go get github.com/guitarpawat/the-internship-2019`.
1. Go to `%GOPATH%\src\github.com\guitarpawat\the-internship-2019\hangman` (Windows) or `$GOPATH/src/github.com/guitarpawat/the-internship-2019/hangman`.
1. Use `go run main.go` to compile and run the program.

## Resource file

The list of words and categories will be found on `/words` in json format.

### Example
```json
{
    "category": "Anime",
    "words": [
        {
            "word": "Silver Spoon", 
            "hint": "Farmer's Life"
        },
        {
            "word": "Tamako Market", 
            "hint": "Mochi Shop"
        },
        {
            "word": "Bakuman",
            "hint": "Manga"
        }
    ]
}
        
```

## References

* [Go rand.Intn same number/value](https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value)
* [How to read input from console line?](https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line)
* [How to replace a letter at a specific index in a string in Go?](https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go)
* [List directory in Go](https://stackoverflow.com/questions/14668850/list-directory-in-go)
* [Package path - The Go Programming Language](https://golang.org/pkg/path/)
* [Package rand - The Go Programming Language](https://golang.org/pkg/math/rand/)
* [Package time - The Go Programming Language](https://golang.org/pkg/time/)
