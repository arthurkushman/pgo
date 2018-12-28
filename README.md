# pgo
Go library for PHP community with convenient functions

[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/pgo)](https://goreportcard.com/report/github.com/arthurkushman/pgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

#### Installation 

Via go get command:
```bash
go get github.com/arthurkushman/pgo
```

Imagine that you need to write Go code every day and also have a convenient functions in memory from PHP experience

You can use date function with similar formatting for PHP e.g.:

```go
pgo.Date("Y-m-d H:i:s")
```

replace sub-strings with StrReplace:
```go
subject := "The quick brown fox jumped over the lazy dog"

pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)
```

or write to files with offset/limit: 
```go
pgo.FileGetContents("somefile.txt", 0, 1024)
```

See more examples from *_test.go files.