package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("hello 这是我的工程")
	fmt.Println(now)
	fmt.Println("2006-01-02 15:04:05             output:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("2006-01-02                      output:", now.Format("2006-01-02"))
	fmt.Println("01-02-2006                      output:", now.Format("01-02-2006"))
	fmt.Println("15:03:04                        output:", now.Format("15:03:04"))
	fmt.Println("2006/01/02 15:04                output:", now.Format("2006/01/02 15:04"))
	fmt.Println("15:04 2006/01/02                output:", now.Format("15:04 2006/01/02"))
	fmt.Println("2006#01#02                      output:", now.Format("2006#01#02"))
	fmt.Println("2006$01$02                      output:", now.Format("2006$01$02"))
	fmt.Println("2006-01-02 15:04:05.000         output:", now.Format("2006-01-02 15:04:05.000"))
	fmt.Println("2006-01-02 15:04:05.000000000   output:", now.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println("2006-January-02 15:04:05 Monday output:", now.Format("2006-January-02 15:04:05 Monday"))
	fmt.Println("2006-Jan-02 15:04:05 Mon        output:", now.Format("2006-Jan-02 15:04:05 Mon"))
	fmt.Println("2006-1-2 3:4:5 PM               output:", now.Format("2006-1-2 3:4:5 PM"))

	fmt.Println("这是develop version")
}
