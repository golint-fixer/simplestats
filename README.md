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

### License

Copyright (c) 2016, linkosmos
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of simplestats nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
