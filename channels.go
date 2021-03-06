package main

import (
     "time"
)

func main() {
     hellCh := make(chan string, 1)
     goodbyeCh := make(chan string, 1)
     quitCh := make(chan bool)
     
     go receiver(hellCh, goodbyeCh, quitCh)
     go sendString(hellCh, "hello!")
     time.Sleep(time.Second)
     go sendString(goodbyeCh, "goodbye!")
     <- quitCh
}


func sendString(ch chan<- string, s string) {
     ch <- s
}

func receiver(hellCh, goodbyeCh <-chan string, quitCh chan<- bool) {
    for {
       select {
         case msg := <- hellCh:
              println(msg)
         case mgs := <- goodbyeCh:
              println(mgs)
         case <- time.After(time.Second * 2):
              println("Nothing received in 2 seconds. Exiting")
              quitCh <- true
              break
       } // end select
    }  // end for
}
