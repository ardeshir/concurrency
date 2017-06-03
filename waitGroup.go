package main

import (
	"fmt"
        "sync"	
)

func main() {
    var wg sync.WaitGroup

    msgPrinter := func(msg string) {
        fmt.Println(msg)
        defer wg.Done()
    }
    wg.Add(2)
    go msgPrinter("First printer")
    go msgPrinter("Second printer...")
    
    wg.Wait()

}


