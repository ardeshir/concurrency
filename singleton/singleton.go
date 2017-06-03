package singleton

import(
    "testing"
    "time"
    "fmt"
)

func TestStartInstance(t *testing.T){
     singleton := GetInstance()
     singleton := GetInstance()

    n := 5000

    for i := 0; i < n ; i++ {
        go singleton.AddOne()
        go singleton2.AddOne()

    }

    fmt.Println("Before loop, current count is %d\n", singleton.GetCount())

    var val int 
    for val != n*2 {
        val = singleton.GetCount()
        time.Sleep(10 * time.Millisecond)
    }
    singleton.Stop()
}

