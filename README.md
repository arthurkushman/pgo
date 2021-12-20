# pgo

Go library for PHP community with convenient functions

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/pgo)](https://goreportcard.com/report/github.com/arthurkushman/pgo)
[![Build and test](https://github.com/arthurkushman/pgo/workflows/Build%20and%20test/badge.svg)](https://github.com/arthurkushman/pgo/actions)
[![codecov](https://codecov.io/gh/arthurkushman/pgo/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/pgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/pgo)

* [Installation](#user-content-installation)
* [Core](#user-content-core)
    * [Serialize/Unserialize](#user-content-serializeunserialize)
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
    * [ArrayUdiff](#user-content-arrayudiff)
    * [ArraySum](#user-content-arraysum)
    * [ArrayIntersect](#user-content-arrayintersect)
    * [Range](#user-content-range)
    * [EqualSlices](#user-content-equalslices)
* [Collections](#user-content-collections)
    * [PriorityQueue](#user-content-priority-queue)
* [Network](#user-content-network)
    * [IP2Long/Long2Ip](#user-content-ip2longlong2ip)
    * [GetMxrr](#user-content-getmxrr)
* [Math](#user-content-math)
    * [Rand](#user-content-rand)
* [Crypto](#user-content-crypt)
    * [Md5](#user-content-md5)
    * [Sha1](#user-content-sha1)
    * [Sha2](#user-content-sha2)
    * [HashFile](#user-content-hashfile)
    * [HashHmac](#user-content-hashhmac)
    * [IsValidMac](#user-content-isvalidmac)

#### Installation

Via go get command:

```bash
go get github.com/arthurkushman/pgo
```

Imagine that you need to write Go code every day and also have a convenient functions in memory from PHP experience

### Core

#### Serialize/Unserialize

For instance, to store Go code data in storage engines like rdbms, no-sql, key-value etc, you can use serialization
functions to serialize Go code to string and unserialize it back from string to Go code:

```go
m := make(map[int]string, 0)
m[0] = "abc"

str, err := pgo.Serialize(m) // str -> "Dv+BBAEC/4IAAQQBDAAACf+CAAEAA2FiYw=="

unserMap := make(map[int]string, 0)
err = pgo.Unserialize(str, &unserMap) // unserMap -> map[int]string{0: "abc"}
```

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

Building a http query string:

```go
queryStr := pgo.HTTPBuildQuery(map[string]interface{}{
  "foo":      "bar",
  "bar":      "baz",
  "s":        []string{"1", "foo", "2", "bar", "3", "baz"},
  "num":      123,
  "bigNum":   int64(1238873737737737373),
  "amount":   623.937,
  "isActive": true,
}) // amount=623.937&bar=baz&bigNum=1238873737737737373&foo=bar&isActive=true&num=123&s=1&s=foo&s=2&s=bar&s=3&s=baz
```

#### StripTags

Strip tags with exclusion rules:

```go
html := "<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore <strong>et dolore</strong> magna aliqua.</div>"

str := html.StripTags(html, []string{"a", "span"}) // results in: "Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."
```

UPD: As had been stated here - https://github.com/golang/go/issues/22639
There is a very handy "stripTags" function in html/template, then guys from official team as fast as they got dislike on
their negative comment, closed the thread. That is why libs like `pgo` is appearing and will be move forward/evelove,
bypassing strict rules that sometimes looking nonsence.

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
pgo.ArrayMap([]string{"foo", "bar", "baz"}, func (v string) string {
return strings.ToUpper(v)
}) // []string{"FOO", "BAR", "BAZ"}

pgo.ArrayMap([]float64{1, 2, 3, 4, 5}, func (v float64) float64 {
return math.Pow(v, 2)
}) // []float64{1, 4, 9, 16, 25}
```

#### ArrayFilter

filters elements of an array using a callback function

```go
pgo.ArrayFilter([]float64{1, 2, 3, 4, 5}, func (v float64) bool {
return v > 2.718
}) // []float64{3, 4, 5}
```

#### ArrayDiff

returns the values in array1 that are not present in any of the other arrays

```go
pgo.ArrayDiff([]string{"foo", "bar", "fizz", "baz"}, []string{"foo", "bar"}) // []string{"fizz", "baz"}
pgo.ArrayDiff([]int{3, 43, 8, 4, 9}, []int{3, 8, 9, 4}) // []int{43}
```

#### ArrayUdiff

computes the difference of arrays by using a callback function for data comparison

```go
pgo.ArrayUdiff(func (a interface{}, b interface{}) int {
if a.(string) > b.(string) {
return 1
} else if a.(string) < b.(string) {
return -1
}

return 0
}, []string{"foo", "bar", "fizz", "baz"}, []string{"foo", "bar"}) // []string{"fizz", "baz"}

pgo.ArrayUdiff(func (a interface{}, b interface{}) int {
if a.(int) > b.(int) {
return 1
} else if a.(int) < b.(int) {
return -1
}

return 0
}, []int{3, 43, 8, 4, 9}, []int{3, 8, 9, 4}) // []int{43}
```

```go
type TestUdiffComparing interface {ComparingValue() string}

type A struct {valueA string}
func (a A) ComparingValue() string  {return a.valueA}

type B struct {valueB string}
func (b B) ComparingValue() string  {return b.valueB}

type C struct {valueC string}
func (c C) ComparingValue() string  {return c.valueC}

func main() {
a := []A {{valueA: "q"}, {valueA: "e"}, {valueA: "t"}, {valueA: "k"}, {valueA: "g"}, {valueA: "o"}, {valueA: "j"}}
b := []B {{valueB: "q"}, {valueB: "g"}, {valueB: "b"}, {valueB: "h"}, {valueB: "j"}, {valueB: "k"}, {valueB: "l"}}
c := []C {{valueC: "o"}, {valueC: "q"}, {valueC: "e"}, {valueC: "x"}, {valueC: "c"}, {valueC: "v"}, {valueC: "b"}}

s1 := pgo.ArrayUdiff(func (arr1 interface{}, arr2 interface{}) int {
if arr1.(TestUdiffComparing).ComparingValue() > arr2.(TestUdiffComparing).ComparingValue() {
return 1
} else if arr1.(TestUdiffComparing).ComparingValue() < arr2.(TestUdiffComparing).ComparingValue() {
return -1
}

return 0
}, a, b, c)

fmt.Print(s1) // [{t}]

a2 := []string {"a", "b", "c"}
b2 := []string {"a", "d", "g"}
c2 := []string {"b", "x", "y"}

s2 := pgo.ArrayUdiff(func (arr1 interface{}, arr2 interface{}) int {
if arr1.(string) > arr2.(string) {
return 1
} else if arr1.(string) < arr2.(string) {
return -1
}

return 0
}, a2, b2, c2)

fmt.Print(s2) // [c]
}

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

#### EqualSlices

Compares two slices and returns true if they are equal, false otherwise (any type of slices support)

```go
res, err := pgo.EqualSlices([]int{1, 2, 3}, []int{1, 2, 3}) // true

res, err := pgo.EqualSlices([]string{"foo"}, []string{"bar"}) // false
```

See more examples in *_test.go files.

### Collections

#### Priority Queue

```go
    // Some items and their priorities.
items := map[string]int{
"banana": 3, "apple": 2, "pear": 4, "peach": 1, "plum": 6,
}

// Create a Priority queue, put the items in it, and
// establish the Priority queue (heap) invariants.
pq := make(pgo.PriorityQueue, len(items))
i := 0
for value, priority := range items {
pq[i] = &pgo.Item{
Value:    value,
Priority: priority,
Index:    i,
}
i++
}
pq.Init()

// Insert a new item and then modify its Priority.
item := &pgo.Item{
Value:    "orange",
Priority: 1,
}
pq.Push(item)
pq.Update(item, item.Value, 5)

item := pq.Pop().(*pgo.Item) // 06:plum
item := pq.Pop().(*pgo.Item) // 05:orange
item := pq.Pop().(*pgo.Item) // 04:pear
item := pq.Pop().(*pgo.Item) // 03:banana
// ... 
```

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

### Math

#### Rand

```go
rand := pgo.Rand(1, 100)
```

### Crypto

#### Md5

```go
pgo.Md5("abc123") // e99a18c428cb38d5f260853678922e03
```

#### Sha1

```go
pgo.Sha1("abc123") // 6367c48dd193d56ea7b0baad25b19455e529f5ee
```

#### Sha2

```go
pgo.Sha2("abc123") // 6ca13d52ca70c883e0f0bb101e425a89e8624de51db2d2392593af6a84118090
```

#### HashFile

```go
hex, err := pgo.HashFile("sha1", "example.txt") // 6367c48dd193d56ea7b0baad25b19455e529f5ee
```

#### HashHmac

```go
hmac := HashHmac("foo bar baz", "secret", sha256.New) // 9efc4f86917b454deae37c869521f88dee79305303fa2283df0b480e3cc8104c
```

#### IsValidMac

```go
IsValidMac("foo bar baz", hmac, "secret", sha256.New) // true/false
```

Supporters gratitude:

<img src="https://github.com/SoliDry/laravel-api/blob/master/tests/images/jetbrains-logo.png" alt="JetBrains logo" width="200" height="166" />