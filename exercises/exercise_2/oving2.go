// Go 1.2
// go run oving2.go

package main

import (
    . "fmt"     
    "runtime"
    "time"
)

var i=0

func add(ic chan int) {
    
	for x := 0; x < 1000000; x++ {
		temp:=<-ic
		temp++
		ic<-temp
	}   
}

func rest(ic chan int) {
   for x := 0; x < 1000000; x++ {
	temp:=<-ic
	temp--
	ic<-temp
	
	} 
}


func main() {
    ic:=make(chan int,1)
    ic<-i
    runtime.GOMAXPROCS(runtime.NumCPU())    
    go add(ic)
    go rest(ic)    
    for x:=0;x<100;x++{
	print := <-ic
        Println(print)
        ic<- print
    }                   

    time.Sleep(100*time.Millisecond)
    Println("Done",i)
}
