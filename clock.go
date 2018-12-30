package main

import (
	"fmt"
	"time"
	)

func main() {
	// tick every 1/10 second
	ticker := time.NewTicker(100*time.Millisecond)

	for {
		now := <- ticker.C
		fmt.Printf("\r%0d:%02d:%02d ", now.Hour(), now.Minute(), now.Second())
	}
}
