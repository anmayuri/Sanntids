// Go 1.2
// go run oving1.go

package main

import (
    . "fmt"     
    "runtime"
    "time"
)

var i=0

func add() {
	for x := 0; x < 1000000; x++ {
			i ++
	}   
}

func rest() {
   for x := 0; x < 1000000; x++ {
			i --
	} 
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    
    go add()
    go rest()    
    for x:=0;x<100;x++{
      Println(i)
    }                   

    time.Sleep(100*time.Millisecond)
    Println("Done",i)
}
