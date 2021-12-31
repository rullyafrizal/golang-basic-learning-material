package concurrent

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Teknik channel timeout digunakan untuk mengontrol penerimaan data dari channel
// ketikaa sudah tidak ada aktivitas penerimaan dalam beberapa tenggat waktu, maka blok timeout bisa dieksekusi

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func receiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func TestChannelTimeout(t *testing.T) {
	rand.Seed(time.Now().Unix())

	var messages = make(chan int)

	go sendData(messages)
	receiveData(messages)
}
