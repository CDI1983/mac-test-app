package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	//i := 0
	ch := make(chan int , 10)
	done := make(chan int , 10)
	line := [100]int{}
	for i:=0;i<10;i++ {
		ch <- i
	}

	for i:=0;i<100;i++ {
		line[i] = i+1
	}

	maptest := make(map[int]int)

	//var i int
	j:= 0
	for {
		if j>98 {
			break
		}
		i :=<- ch
		done <- i
		maptest[i] = line[j]
		wg.Add(1)
		go func() {
			l:=<-done
			if i == l{
				fmt.Println("goroutine -> ", l)
				//fmt.Println("test分支")
				fmt.Println("data = ", maptest[l])
				ch <- l
				wg.Done()
			}
			}()
		j++
	}
wg.Wait()
}
