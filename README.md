# pgo
Go library for PHP community with convenient functions

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/pgo)](https://goreportcard.com/report/github.com/arthurkushman/pgo)
[![codecov](https://codecov.io/gh/arthurkushman/pgo/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/pgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/pgo)

* [Installation](#user-content-installation)
* [Date](#user-content-date)
* [Strings](#user-content-strings)
* [Files](#user-content-files)
* [Arrays](#user-content-arrays)

#### Installation 

Via go get command:
```bash
go get github.com/arthurkushman/pgo
```

Imagine that you need to write Go code every day and also have a convenient functions in memory from PHP experience

### Date
You can use date function with similar formatting for PHP e.g.:

```go
dateStr := pgo.Date("Y-m-d H:i:s")
```

### Strings

#### replace sub-strings with StrReplace:
```go
subject := "The quick brown fox jumped over the lazy dog"

str, err := pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)

// and if case-insensitive replace needed - pgo.StrIReplace([]string{"DOG", "QuiCK"}, []string{"fox", "slow"}, subject) 
```

#### Bulding a http query string:
```go
queryStr := pgo.HTTPBuildQuery(map[string]string{
	"foo": "bar",
	"bar": "baz",
}) // bar=baz&foo=bar
```

#### Strip tags with exclusion rules:
```go
html := "<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore <strong>et dolore</strong> magna aliqua.</div>"

str := html.StripTags(html, []string{"a", "span"}) // results in: "Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."
```
UPD: As had been stated here - https://github.com/golang/go/issues/22639
There is a very handy "stripTags" function in html/template, then guys from official team as fast as they got dislike on their negative comment, closed the thread.
That is why libs like `pgo` is appearing and will be move forward/evelove, bypassing strict rules that sometimes looking nonsence.

### Files

Read files with offset/limit:
```go
content, err := pgo.FileGetContents("somefile.txt", 0, 1024)
```

reflexively write to files with:
```go
n, err := pgo.FilePutContents("somefile.txt", strToWrite, pgo.FileAppend)
```

#### Read from context (via http(s)):
```go
content, err := pgo.FileGetContents("http://google.com", pgo.NewContext())
```

#### Uploading files from web-forms to your server:
```go
ctx := pgo.NewContext()
ctx.Req = YourReq
ctx.UploadMaxFileSize = 10 << 25

uploaded := ctx.MoveUploadedFile("foo", "/srv/images/pic123.png")
```

#### Checking for file existence
```go
if pgo.FileExists("file1.txt") == true {
	// do something with existent file
}
```

#### Check if it is file/dir/symlink 
```go 
if pgo.IsFile("someFile.txt") {
	// do something with file
}

if pgo.IsDir("someDir/") {
	// do something with dir
}

if pgo.IsLink("someLink") {
	// do somthing with symlink
}
```

### Arrays

#### Check if an array contains an element
```go
pgo.InArray(3, []int{1, 2, 3}) // true
pgo.InArray("bar33", []string{"foo", "bar", "baz"}) // false
pgo.InArray(3.14159, []float64{33.12, 12.333, 3.14159, 78.4429}) // true
```

#### Split an array by chunks (with auto-tailing)
```go
pgo.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2) // [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}

pgo.ArrayChunk([]string{"foo", "bar", "baz", "fizz", "buzz"}, 3) // [][]string{[]string{"foo", "bar", "baz"}, []string{"fizz", "buzz"}}
```

#### Create an array by using one array for keys and another for its values
```go
pgo.ArrayCombine([]int{11, 32, 13, 14, 51, 46, 17, 88}, []string{"foo", "bar", "baz", "fizz", "buzz", "mazz", "freez", "lorum"})
/*
map[int]string{
		11: "foo",
		32: "bar",
		13: "baz",
		14: "fizz",
		51: "buzz",
		46: "mazz",
		17: "freez",
		88: "lorum",
	}	
*/
```

See more examples in *_test.go files.
