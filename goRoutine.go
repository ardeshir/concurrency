package main

import (
	"fmt"
        "sync"	
)

func main() {
    var wg sync.WaitGroup
    goRoutine := 15
    wg.Add(goRoutine)
    var inxd = goRoutine

    msgPrinter := func(msg string, inx int, goRoutine int) {
        fmt.Printf(" %s, Inxd:%d, goRoutine: %d\n",  msg, inx, goRoutine)
        defer wg.Done()

    }

    for i := 0;  i < goRoutine; i++ {
      inxd = inxd - 1 
      go msgPrinter("First  goRoutine ", inxd, goRoutine)
      go msgPrinter("Second goRoutine ", inxd, goRoutine)
    }
       
    wg.Wait()

}


