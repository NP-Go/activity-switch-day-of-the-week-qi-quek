package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	// fmt.Printf("%v", time.Monday)
	// fmt.Printf("%v", today)

	switch today {
	case time.Monday:
		fmt.Println("Today is ", today)
	case time.Tuesday:
		fmt.Println("Today is ", today)
	case time.Wednesday:
		fmt.Println("Today is ", today)
	case time.Thursday:
		fmt.Println("Today is ", today)
	case time.Friday:
		fmt.Println("Today is ", today)
	default:
		fmt.Println("Its the weekend")
	}
}
