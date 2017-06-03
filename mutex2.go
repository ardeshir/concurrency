package main

import (
        "sync"
        "time"
)

type Counter struct {
  sync.Mutex
  value int
}

func main() {
   counter := Counter{}

   for i := 0; i < 2; i++ {
       go func(i int) {
          counter.value++
       }(i)
   }
   time.Sleep(time.Second)
   counter.Lock()
   println(counter.value)
   counter.Unlock()
   
}

