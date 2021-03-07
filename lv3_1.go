package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var lock sync.Mutex
// 统计素数个数
var counter = 0

// inputNumbers 产生数据
func inputNumbers(startNumber, endNumber int, inner chan<- int) {
	for i := startNumber; i < endNumber; i++ {
		inner <- i
	}
	close(inner)
	wg.Done()
}

// prime 生产素数，存入通道 primer
func prime(inner <-chan int, primer, flag chan<- int) {
	for i := range inner {
		flag := true
		// 除到一半即可判断出是否为素数
		// 加1是防止生产的数据为单数，如5555等
		for j := 2; j < i/2+1; j++ {
			// 判断是否为素数，不是则跳出内层循环
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag && i != 1 {
			// 将素数发送到通道
			primer <- i
			lock.Lock() // 加锁
			counter++
			lock.Unlock() // 解锁
		}
	}
	flag <- 1
	wg.Done()
}

// printPrime 打印素数
func printPrime(primer <-chan int) {
	for _ = range primer {
		for v := range primer {
			fmt.Print(v, " ")

		}
	}
	fmt.Println("\nTotal primes:", counter)
	wg.Done()
}

func main() {
	start := time.Now()
	// inner 生产者通道，生产原始数据
	inner := make(chan int, 1000)
	// primer 消费原始数据，生产素数
	primer := make(chan int, 1000)
	// flag 设置 primer 结束标记，容量设为开启协程数量
	flag := make(chan int, 10)

	wg.Add(1)
	go inputNumbers(1, 50000, inner) // 1~100000

	// 开启 cap(flag) 个协程，获取素数
	for i := 0; i < cap(flag); i++ {
		wg.Add(1)
		go prime(inner, primer, flag)
	}
	wg.Add(1)
	go printPrime(primer)

	wg.Add(1)
	go func() {
		for {
			if len(flag) == cap(flag) {
				close(primer)
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()
	end := time.Now()
	// 计算耗费时间
	fmt.Println("Total time  :", end.Sub(start))
}