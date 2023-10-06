
# go-ord

Package **ord** provides tools for turning numbers into their human-friendly ordinal representation.

For example:

| Number | (English) Ordinal |
|--------|-------------------|
|      1 |               1st |
|      2 |               2nd |
|      3 |               3rd |
|      4 |               4th |
|      5 |               5th |
|    ... |               ... |
|     20 |              20th |
|     21 |              21st |
|     22 |              22nd |
|     23 |              23rd |
|     24 |              24th |
|     25 |              25th |
|     26 |              26th |
|     27 |              27th |
|    ... |               ... |

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-ord

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-ord?status.svg)](https://godoc.org/sourcecode.social/reiver/go-ord)

## Example:

```go
import "sourcecode.social/reiver/go-ord/en"

// ...

var n int64 = 123

fmt.Printf("You are the %s person in line.", orden.FormatInt64(n))

// Output:
// You are the 123rd person in line.
```

## Import

To import package **opt** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-ord"
```

## Installation

To install package **opt** do the following:
```
GOPROXY=direct https://sourcecode.social/reiver/go-ord
```

## Author

Package **opt** was written by [Charles Iliya Krempeaux](http://changelog.ca)
