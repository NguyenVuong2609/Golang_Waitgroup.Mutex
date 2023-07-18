package MutexDemo

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex // Khai báo một biến Mutex

func increment() {
	mutex.Lock() // Khóa Mutex trước khi thay đổi biến counter
	counter++
	mutex.Unlock() // Mở khóa Mutex sau khi hoàn thành công việc
}

func TestMutex() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
