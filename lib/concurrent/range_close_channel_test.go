package concurrent

import (
	"fmt"
	"testing"
)

// for - range bisa diterapkan pada channel guna handle penerimaan data
// setiap kali ada pengiriman data via channel, for-range akan berjalan
// looping hanya akan berhenti jika channel di-close

func sendMessage(ch chan<- string) {
	for i := 1; i <= 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMsg(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func TestRangeClosChannel(t *testing.T) {
	var messages = make(chan string)

	go sendMessage(messages)

	printMsg(messages)
}

// Channel direction
// Adalah konsep level akses channel yang ada di parameter function, apakah sebagai penerima, pengirim, atau keduanya

// ch chan string (bisa untuk menerima dan mengirim data)
// ch chan<- string (bisa untuk mengirim data saja)
// ch <-chan string (bisa untuk menerima data)
