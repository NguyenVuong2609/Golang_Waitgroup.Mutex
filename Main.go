package main

import (
	"Day9-Waitgroup-Mutex/MutexDemo"
	"Day9-Waitgroup-Mutex/WaitgroupDemo"
	"fmt"
)

func main() {
	fmt.Println("------------------ Wait Group ----------------------")
	WaitgroupDemo.TestWaitGroup()
	fmt.Println("------------------ Mutex ----------------------")
	MutexDemo.TestMutex()
}
