package main

import (
	"fmt"
	"time"
)

func main() {
	go count("heechang") //run in background after main()
	count("kang")

	//fmt.Scanln() //wait till input
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
