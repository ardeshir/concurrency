package main

import (
	"fmt"
        "sync"
        "strings"	
)

var wg sync.WaitGroup

func toUpperAsync(word string, f func(string)) {
  go func(){
      f(strings.ToUpper(word))
  }()
}

func main() {
    wg.Add(1)

    toUpperAsync("callback to upper case", func(v string) {
       fmt.Printf("String passed: %s\n", v)
       wg.Done()
    })

    fmt.Println("Waiting for async response...")
    wg.Wait()
   
}


