package main

// prints local clock time

import (
	"fmt"
	"time"
	)

func main() {
	now := time.Now()
	fmt.Printf("\r%0d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
}
