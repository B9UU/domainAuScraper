package main

import (
	"fmt"
	"time"
)

func Timer(name int) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%v took %v \n", name, time.Since(start))
	}
}
