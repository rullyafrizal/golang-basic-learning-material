package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

// Channel digunakan untuk emnghubungkan Goroutine satu dengan lainnya.
// Fungsi sebuah Channel adalah menjadi pengirim dan penerima antara 2 Goroutine.
// Pengirim dan penerima channel bersifat blocking atau syncrhonous

func TestChannel(t *testing.T) {
	runtime.GOMAXPROCS(2)

	message := make(chan string) // deklarasi channle menggunakan make(chan {tipe-data})

	// 3 function yang dijalankan bersama Goroutine secara asynchronous

	go func(who string) {
		data := fmt.Sprintf("hello %s", who)
		message <- data // mengirim data ke channel
	}("John Doe 1")

	go func(who string) {
		data := fmt.Sprintf("hello %s", who)
		message <- data // mengirim data ke channel
	}("John Doe 2")

	go func(who string) {
		data := fmt.Sprintf("hello %s", who)
		message <- data // mengirim data ke channel
	}("John Doe 3")

	message1Receiver := <-message // variabel penerima
	fmt.Println(message1Receiver)

	message2Receiver := <-message // variabel penerima
	fmt.Println(message2Receiver)

	message3Receiver := <-message // variabel penerima
	fmt.Println(message3Receiver)
}

func TestChannelAsParam(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"John", "Wick", "Doe"} {
		// eksekusi IIFE (Immediately Invoked Function Expression) secara langsung dan menggunakan Goroutine
		go func(who string) {
			data := fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

// channel as a parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}

func TestBufferedChannel(t *testing.T) {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3) // jumlah buffer 3 karena dihitung dari 0

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)

		messages <- i
	}
}
