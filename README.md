# simplestats
Simple stats for tracking internal stats

[![Build Status](https://travis-ci.org/linkosmos/simplestats.svg?branch=master)](https://travis-ci.org/linkosmos/simplestats)
[![Coverage Status](https://coveralls.io/repos/github/linkosmos/simplestats/badge.svg?branch=master)](https://coveralls.io/github/linkosmos/simplestats?branch=master)
[![GoDoc](http://godoc.org/github.com/linkosmos/simplestats?status.svg)](http://godoc.org/github.com/linkosmos/simplestats)
[![Go Report Card](http://goreportcard.com/badge/linkosmos/simplestats)](http://goreportcard.com/report/linkosmos/simplestats)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

### Methods

 - New() *Stats
 - Increment(key string)
 - IncrementBy(key string, value int64)
 - Get(key string) int64

### Usage

```go
  package main

  import (
    "fmt"

    "github.com/linkosmos/simplestats"
  )

  func main() {
    s := simplestats.New()
    s.Increment("one")
    s.Increment("two")
    s.Increment("one")

    fmt.Println(s.Get("one"))

    // 2
  }
```
