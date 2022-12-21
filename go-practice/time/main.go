package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	fmt.Printf("oneMonthLater: %v\n", oneMonthLater)
	oneMonthBefore := now.AddDate(0, -1, 0)
	fmt.Printf("oneMonthBefore: %v\n", oneMonthBefore)
}
