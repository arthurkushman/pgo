# pgo
Go library for PHP community with convenient functions

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/pgo)](https://goreportcard.com/report/github.com/arthurkushman/pgo)
[![codecov](https://codecov.io/gh/arthurkushman/pgo/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/pgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/pgo)

#### Installation 

Via go get command:
```bash
go get github.com/arthurkushman/pgo
```

Imagine that you need to write Go code every day and also have a convenient functions in memory from PHP experience

You can use date function with similar formatting for PHP e.g.:

```go
dateStr := pgo.Date("Y-m-d H:i:s")
```

replace sub-strings with StrReplace:
```go
subject := "The quick brown fox jumped over the lazy dog"

str, err := pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)

// and if case-insensitive replace needed - pgo.StrIReplace([]string{"DOG", "QuiCK"}, []string{"fox", "slow"}, subject) 
```

or read files with offset/limit: 
```go
content, err := pgo.FileGetContents("somefile.txt", 0, 1024)
```

reflexively write to files with:
```go
n, err := pgo.FilePutContents("somefile.txt", strToWrite, pgo.FileAppend)
```

Read from context (via http(s)):
```go
content, err := pgo.FileGetContents("http://google.com", pgo.NewContext())
```

Uploading files from web-forms to your server:
```go
ctx := pgo.NewContext()
ctx.Req = YourReq
ctx.UploadMaxFileSize = 10 << 25

uploaded := ctx.MoveUploadedFile("foo", "/srv/images/pic123.png")
```

Bulding a http query string:
```go
queryStr := pgo.HTTPBuildQuery(map[string]string{
	"foo": "bar",
	"bar": "baz",
}) // bar=baz&foo=bar
```

Strip tags with exclusion rules:
```go
html := "<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore <strong>et dolore</strong> magna aliqua.</div>"

str := html.StripTags(html, []string{"a", "span"}) // results in: "Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."
```
UPD: As had been stated here - https://github.com/golang/go/issues/22639
There is a very handy "stripTags" function in html/template, then guys from official team as fast as they got dislike on their negative comment, closed the thread.
That is why libs like `pgo` is appearing and will be move forward/evelove, bypassing strict rules that sometimes looking nonsence.

See more examples from *_test.go files.
