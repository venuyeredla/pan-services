package learngo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type A struct {
	Mutex sync.Mutex
	Value int
}

type B struct {
	Mutex sync.Mutex
	Value int
}

func DeadLock() {
	var wg sync.WaitGroup
	wg.Add(2)
	a := &A{Value: 1}
	b := &B{Value: 1}

	var Worker1 = func(a *A, b *B) {
		defer wg.Done()
		a.Mutex.Lock()
		fmt.Println("goroutine 1 acquired lock 1")
		time.Sleep(100 * time.Microsecond)
		b.Mutex.Lock()
		fmt.Println("goroutine 1 acquired lock 2")
		//a.Mutex.Unlock()
	}

	var Worker2 = func(a *A, b *B) {
		defer wg.Done()
		b.Mutex.Lock()
		fmt.Println("goroutine 2 acquired lock 2")
		time.Sleep(100 * time.Microsecond)
		a.Mutex.Lock()
		fmt.Println("goroutine 2 acquired lock 1")
		//	b.Mutex.Unlock()
	}
	go Worker1(a, b)
	go Worker2(a, b)
}

func PubSub() {
	wg.Add(2)
	intChannel := make(chan int)
	var Publisher = func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Publisher adding : ", i)
			intChannel <- i
			time.Sleep(time.Duration(5 * time.Millisecond))
		}
		close(intChannel)
		wg.Done()
	}
	var Consumer = func() {
		for val := range intChannel {
			fmt.Println("Consumer : ", val)
		}
		wg.Done()
	}
	go Publisher()
	go Consumer()
	wg.Wait()

}

var wg = sync.WaitGroup{}

func TestGoRoutine() {

}

var wait sync.WaitGroup
var count int

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)

	}
	wait.Done()

}
func RaceTest() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

var mutex sync.Mutex

func Mutexincrement(s string) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
		mutex.Unlock()

		//  Atmoic variable    atomic.AddInt64(&count,1)

	}
	wait.Done()

}
func MutexTest() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func ChannelTest() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
