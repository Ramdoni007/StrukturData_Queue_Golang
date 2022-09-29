package main

import "fmt"

// Queue adalah struktur data antrian dimana data yang pertama masuk
// akan keluar pertama juga.

// 1. Saya akan membuat suatu struct yang di dalam nya mempunyai slice items.
type Queue struct {
	items []int
}

// 2. Disini Saya Juga membuat suatu function Enqueue memasukan data ke dalam elemen sisi belakang (rear)
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

// 3.Disini Saya Juga suatu Function Dequeu yaitu mengambil data yang sudah kita masukan ke sisi depan (front).
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}

	item, items := q.items[0], q.items[1:]
	q.items = items
	return item

}

// Dan di Function main ini saya akan menjalankan dan memmanggil masing masing function tadi yang saya buat.
// Dan Struktur data Queue berjalan sesuai.
func main() {
	// Memasukan Data 1 - 5
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	println(q.Dequeue())
	fmt.Printf("%v", q.items)

	println(q.Dequeue())
	fmt.Printf("%v", q.items)

	println(q.Dequeue())
	fmt.Printf("%v", q.items)

	println(q.Dequeue())
	fmt.Printf("%v", q.items)

	println(q.Dequeue())
	fmt.Printf("%v", q.items)

}
