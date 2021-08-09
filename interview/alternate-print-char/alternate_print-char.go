package interview

import (
	"fmt"
	"sync"
)

/*
实现：用到了两个`channel`负责通知，letter负责通知打印字母的goroutine来打印字母，number用来通知打印数字的goroutine打印数字。
wait用来等待字母打印完成后退出循环。
*/

func alternatePrintChar() {
	var wg sync.WaitGroup
	numChannel := make(chan bool)
	letterChannel := make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-numChannel:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letterChannel <- true
			}
		}
	}()
	wg.Add(1)
	go func() {
		i := 'A'
		for {
			select {
			case <-letterChannel:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				numChannel <- true
			}
		}
	}()
	numChannel <- true
	wg.Wait()
	return
}
