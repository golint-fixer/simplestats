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
