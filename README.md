# pgo
Go library for PHP community with convenient functions

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/pgo)](https://goreportcard.com/report/github.com/arthurkushman/pgo)
[![codecov](https://codecov.io/gh/arthurkushman/pgo/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/pgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/pgo)

* [Installation](#user-content-installation)
* [Date](#user-content-date)
    * [Milli/Micro](#user-content-millimicro)
* [Strings](#user-content-strings)
	* [StrReplace/StrIReplace](#user-content-strreplace)
	* [HTTPBuildQuery](#user-content-httpbuildquery)
	* [StripTags](#user-content-striptags)
* [Files](#user-content-files)
	* [FileGetContents](#user-content-filegetcontents)
	* [FilePutContents](#user-content-fileputcontents)
	* [MoveUploadedFile](#user-content-moveuploadedfile)
	* [FileExists](#user-content-fileexists)
	* [IsFile/IsDir/IsLink](#user-content-isfileisdirislink)
* [Arrays](#user-content-arrays)
	* [InArray](#user-content-inarray)
	* [ArrayChunk](#user-content-arraychunk)
	* [ArrayCombine](#user-content-arraycombine)
	* [ArrayCountValues](#user-content-arraycountvalues)
	* [ArrayMap](#user-content-arraymap)
	* [ArrayFilter](#user-content-arrayfilter)
	* [ArrayDiff](#user-content-arraydiff)
	* [ArrayKeys](#user-content-arraykeys)
	* [ArraySum](#user-content-arraysum)
	* [ArrayIntersect](#user-content-arrayintersect)
	* [Range](#user-content-range)
* [Network](#user-content-network)
    * [IP2Long/Long2Ip](#user-content-ip2longlong2ip)
    * [GetMxrr](#user-content-getmxrr)

#### Installation 

Via go get command:
```bash
go get github.com/arthurkushman/pgo
```

Imagine that you need to write Go code every day and also have a convenient functions in memory from PHP experience

### Date
You can use date function with similar formatting for PHP e.g.:

```go
dateStr := pgo.Date("Y-m-d H:i:s") // 2019-03-28 12:23:03

pgo.Date("j D, M") // 27 Wed, Mar

pgo.Date("Q") // 3 (of 1,2,3,4 quarters)
```

#### Milli/Micro
```go
nowMicro := pgo.UnixMicro() // get current unix microseconds
nowMilli := pgo.UnixMilli() // get current unix milliseconds

// get current millis + 3ms 
nowMillisPlusThree := pgo.Time(time.Now().Add(time.Millisecond * 3)).Milliseconds()
// get current microseconds + 7μs 
nowMicroPlusSeven := pgo.Time(now.Add(time.Microsecond * 7)).Microseconds()
```

### Strings

#### StrReplace
replace sub-strings with StrReplace:
```go
subject := "The quick brown fox jumped over the lazy dog"

str, err := pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)

// and if case-insensitive replace needed - pgo.StrIReplace([]string{"DOG", "QuiCK"}, []string{"fox", "slow"}, subject) 
```

#### HTTPBuildQuery
Bulding a http query string:
```go
queryStr := pgo.HTTPBuildQuery(map[string]string{
	"foo": "bar",
	"bar": "baz",
}) // bar=baz&foo=bar
```

#### StripTags
Strip tags with exclusion rules:
```go
html := "<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore <strong>et dolore</strong> magna aliqua.</div>"

str := html.StripTags(html, []string{"a", "span"}) // results in: "Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."
```
UPD: As had been stated here - https://github.com/golang/go/issues/22639
There is a very handy "stripTags" function in html/template, then guys from official team as fast as they got dislike on their negative comment, closed the thread.
That is why libs like `pgo` is appearing and will be move forward/evelove, bypassing strict rules that sometimes looking nonsence.

### Files

#### FileGetContents
Read files with offset/limit:
```go
content, err := pgo.FileGetContents("somefile.txt", 0, 1024)
```

#### FilePutContents
reflexively write to files with:
```go
n, err := pgo.FilePutContents("somefile.txt", strToWrite, pgo.FileAppend)
```

Read from context (via http(s)):
```go
content, err := pgo.FileGetContents("http://google.com", pgo.NewContext())
```

#### MoveUploadedFile
Uploading files from web-forms to your server:
```go
ctx := pgo.NewContext()
ctx.Req = YourReq
ctx.UploadMaxFileSize = 10 << 25

uploaded := ctx.MoveUploadedFile("foo", "/srv/images/pic123.png")
```

#### FileExists
Checking for file existence
```go
if pgo.FileExists("file1.txt") == true {
	// do something with existent file
}
```

#### IsFile/IsDir/IsLink
Check if it is file/dir/symlink 
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

#### InArray
Check if an array contains an element
```go
pgo.InArray(3, []int{1, 2, 3}) // true
pgo.InArray("bar33", []string{"foo", "bar", "baz"}) // false
pgo.InArray(3.14159, []float64{33.12, 12.333, 3.14159, 78.4429}) // true
```

#### ArrayChunk
Split an array by chunks (with auto-tailing)
```go
pgo.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2) // [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}

pgo.ArrayChunk([]string{"foo", "bar", "baz", "fizz", "buzz"}, 3) // [][]string{[]string{"foo", "bar", "baz"}, []string{"fizz", "buzz"}}
```

#### ArrayCombine 
Create an array by using one array for keys and another for its values
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
pgo.ArrayCombine([]string{"foo", "bar", "baz", "fizz", "buzz"}, []float64{11.32, 32.42, 13.246, 14.41, 51.98})
/*
map[string]float64{
			"foo":  11.32,
			"bar":  32.42,
			"baz":  13.246,
			"fizz": 14.41,
			"buzz": 51.98,
		}
*/
```

#### ArrayCountValues
Count all the values of an array/slice
```go
pgo.ArrayCountValues([]string{"foo", "bar", "foo", "baz", "bar", "bar"}) // map[string]int{"foo": 2, "bar": 3, "baz": 1}

pgo.ArrayCountValues([]float64{3.14159, 43.03, 8, 3.14159, 43.02, 8}) // map[float64]int{3.14159: 2, 8: 2, 43.03: 1, 43.02: 1}
```

#### ArrayMap
Apply the callback to the elements of the given arrays
```go
pgo.ArrayMap([]string{"foo", "bar", "baz"}, func(v string) string {
		return strings.ToUpper(v)
	}) // []string{"FOO", "BAR", "BAZ"}

pgo.ArrayMap([]float64{1, 2, 3, 4, 5}, func(v float64) float64 {
		return math.Pow(v, 2)
	}) // []float64{1, 4, 9, 16, 25}
```

#### ArrayFilter
filters elements of an array using a callback function
```go
pgo.ArrayFilter([]float64{1, 2, 3, 4, 5}, func(v float64) bool {
		return v > 2.718
	}) // []float64{3, 4, 5}
```

#### ArrayDiff
returns the values in array1 that are not present in any of the other arrays
```go
pgo.ArrayDiff([]string{"foo", "bar", "fizz", "baz"}, []string{"foo", "bar"}) // []string{"fizz", "baz"}
pgo.ArrayDiff([]int{3, 43, 8, 4, 9}, []int{3, 8, 9, 4}) // []int{43}
```

#### ArrayKeys
returns an array using the keys of another array
```go
pgo.ArrayKeys(map[string]int{"foo": 1, "bar": 8, "fizz": 12, "baz": 0}) // []string{"foo", "bar", "fizz", "baz"}

pgo.ArrayKeys(map[interface{}]int{3.45: 32, "foo": 33, 8: 53, "bar": 1, 9: 1}) // []interface{}{3.45, "foo", 8, "bar", 9}
```

#### ArraySum
calculate the sum of values in an array
```go
pgo.ArraySum([]int{12, 54, 32, 12, 33}) // int: 143
```

#### ArrayIntersect
computes the intersection of arrays
```go
pgo.ArrayIntersect([]int{12, 54, 32, 12, 33}, []int{3, 12, 54, 9}, []int{12, 33, 9}) // []int{12, 54, 33}

pgo.ArrayIntersect([]string{"foo", "bar", "baz", "fizz", "bazz", "fizz", "fizz"}, []string{"bar", "fizz"}, []string{"foo", "bar", "hey"}) // []string{"foo", "bar", "fizz"}
```

#### Range
creates an int slice of min to max range
```go
pgo.Range(3, 9) // []int{3, 4, 5, 6, 7, 8, 9}

// If a step value is given, it will be used as the increment between elements in the sequence.
pgo.Range(-3, 7, 5) // []int{-3, 2, 7}
```
See more examples in *_test.go files.

### Network

#### IP2Long/Long2Ip

```go
long, _ := pgo.IP2long("176.59.34.117") // 2956665461

ip := pgo.Long2ip(2956665461) // "176.59.34.117"
```

#### GetMxrr
```go
isMx, mxs, _ := pgo.GetMxrr("google.com") // e.g.: true, n
```
