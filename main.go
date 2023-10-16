package main

import (
	"fmt"
	"time"
)

// TASK 1
// printNumbers mencetak numbers dari 1 hingga 10
func printNumbers() {
	// Mengulangi dari 1 hingga 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100) // Mensimulasikan beberapa pekerjaan
	}
}

// printLetters mencetak huruf dari 'a' hingga 'j'
func printLetters() {
	// Mengulangi dari 'a' hingga 'j'
	for j := 'a'; j <= 'j'; j++ {
		fmt.Printf("%c\n", j)
		time.Sleep(time.Millisecond * 100) // Mensimulasikan beberapa pekerjaan
	}
}

func main() {
	// Start a goroutine to execute printNumbers concurrently
	go printNumbers()

	// Start another goroutine to execute printLetters concurrently
	go printLetters()

	// Tunggu selama satu detik untuk memungkinkan gorutin selesai
	time.Sleep(time.Second)
}

// // TASK 2
// // prints numbers from 1 to 10
// func printNumbers() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// // prints letters from 'a' to 'j'
// func printLetters() {
// 	for j := 'a'; j <= 'j'; j++ {
// 		fmt.Printf("%c\n", j)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup // Membuat WaitGroup untuk menunggu goroutine

// 	// Menambahkan 2 ke WaitGroup karena memiliki 2 goroutine
// 	wg.Add(2)

// 	// Memulai goroutine untuk printNumbers
// 	go func() {
// 		defer wg.Done() // Mengurangi counter WaitGroup saat selesai, wg.Done() akan selalu dipanggil bahkan jika terjadi error atau panic di dalam goroutine
// 		printNumbers()  // Memanggil fungsi untuk mencetak angka
// 	}()

// 	// Memulai goroutine untuk printLetters
// 	go func() {
// 		defer wg.Done() // Mengurangi counter WaitGroup saat selesai
// 		printLetters()  // Memanggil fungsi untuk mencetak huruf
// 	}()

// 	// Menunggu kedua goroutine selesai
// 	wg.Wait()
// }

// // TASK 3
// // Fungsi untuk mengirimkan angka dari 1 sampai 10 melalui channel
// func produce(ch chan int, wg *sync.WaitGroup) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i // Mengirimkan 'i' melalui channel
// 	}
// 	close(ch) // Menutup channel setelah selesai mengirim
// 	wg.Done() // Memberi tahu WaitGroup bahwa goroutine ini telah selesai
// }

// // Fungsi untuk membaca angka dari channel dan mencetaknya
// func consume(ch chan int, wg *sync.WaitGroup) {
// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// 	wg.Done() // Memberi tahu WaitGroup bahwa goroutine ini telah selesai
// }

// func main() {
// 	ch := make(chan int)  // Membuat channel tipe int
// 	var wg sync.WaitGroup // Membuat WaitGroup untuk sinkronisasi goroutine

// 	// Menambahkan 2 ke WaitGroup karena memiliki 2 goroutine
// 	wg.Add(2)

// 	// Memulai goroutine untuk menghasilkan angka
// 	go produce(ch, &wg)

// 	// Memulai goroutine lain untuk mengonsumsi angka
// 	go consume(ch, &wg)

// 	// Menunggu kedua goroutine selesai
// 	wg.Wait()
// }

// // TASK 4
// // Fungsi untuk mengirimkan angka dari 1 sampai 10 melalui channel
// func produce(ch chan int, wg *sync.WaitGroup) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i // Mengirimkan 'i' melalui channel
// 	}
// 	close(ch) // Menutup channel setelah selesai mengirim
// 	wg.Done() // Memberi tahu WaitGroup bahwa goroutine ini telah selesai
// }

// // Fungsi untuk membaca angka dari channel dan mencetaknya
// func consume(ch chan int, wg *sync.WaitGroup) {
// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// 	wg.Done() // Memberi tahu WaitGroup bahwa goroutine ini telah selesai
// }

// func main() {
// 	ch := make(chan int, 5) // Membuat channel tipe int dengan buffer berukuran 5
// 	var wg sync.WaitGroup   // Membuat WaitGroup untuk sinkronisasi goroutine

// 	// Menambahkan 2 ke WaitGroup karena kita memiliki 2 goroutine
// 	wg.Add(2)

// 	// Memulai goroutine untuk menghasilkan angka
// 	go produce(ch, &wg)

// 	// Memulai goroutine lain untuk mengonsumsi angka
// 	go consume(ch, &wg)

// 	// Menunggu kedua goroutine selesai
// 	wg.Wait()
// }

// // Behavior difference when using a buffered channel versus an unbuffered one
// // Channel Tanpa Buffer:
// // - Ketika menulis data ke channel (menggunakan <-), goroutine pengirim akan terblokir hingga ada goroutine penerima yang siap untuk membaca dari channel tersebut.

// // Channel dengan Buffer:
// // - Ketika membuat channel dengan buffer (seperti make(chan int, 5)), maka channel tersebut memiliki kapasitas buffer sebanyak 5 elemen.
// // - Jika buffer tidak penuh, goroutine pengirim dapat menulis data ke dalam buffer tanpa menunggu goroutine penerima.
// // - Jika buffer penuh, pengirim akan tetap berjalan dan tidak terblokir, kecuali ada goroutine penerima yang membaca dari channel.

// // TASK 5
// func main() {
// 	// Membuat dua channel untuk angka ganjil dan genap
// 	oddCh := make(chan int)
// 	evenCh := make(chan int)

// 	// Memulai goroutine untuk mengirim angka ke channel
// 	go numbers(oddCh, evenCh)

// 	// Iterasi dari 0 hingga 19
// 	for i := 0; i < 20; i++ {
// 		select {
// 		// Jika ada data di evenCh, cetak sebagai angka genap
// 		case even := <-evenCh:
// 			fmt.Printf("Received an even number\t: %v\n", even)

// 		// Jika ada data di oddCh, cetak sebagai angka ganjil
// 		case odd := <-oddCh:
// 			fmt.Printf("Received an odd number\t: %v\n", odd)
// 		}
// 	}
// }

// // Fungsi numbers mengirim angka genap ke evenCh dan angka ganjil ke oddCh
// func numbers(oddCh chan<- int, evenCh chan<- int) {
// 	// Iterasi dari 1 hingga 20
// 	for i := 1; i <= 20; i++ {
// 		if i%2 == 0 {
// 			// Kirim angka genap ke evenCh
// 			evenCh <- i
// 		} else {
// 			// Kirim angka ganjil ke oddCh
// 			oddCh <- i
// 		}
// 	}

// 	// Tutup channel setelah selesai mengirim
// 	close(oddCh)
// 	close(evenCh)
// }

// // TASK 6
// func main() {
// 	// Membuat dua channel untuk angka ganjil, genap, dan error
// 	oddCh := make(chan int)
// 	evenCh := make(chan int)
// 	errorCh := make(chan string)

// 	// Memulai goroutine untuk mengirim angka ke channel
// 	go numbers(oddCh, evenCh, errorCh)

// 	// Iterasi dari 1 hingga 22 (termasuk angka 21 dan 22)
// 	for i := 1; i <= 22; i++ {
// 		select {
// 		// Jika ada data di evenCh, cetak sebagai angka genap
// 		case even := <-evenCh:
// 			fmt.Printf("Received an even number\t: %v\n", even)

// 		// Jika ada data di oddCh, cetak sebagai angka ganjil
// 		case odd := <-oddCh:
// 			fmt.Printf("Received an odd number\t: %v\n", odd)

// 		// Jika ada pesan error di errorCh, cetak pesan error
// 		case err := <-errorCh:
// 			fmt.Println(err)
// 		}
// 	}
// }

// // Fungsi numbers mengirim angka genap ke evenCh dan angka ganjil ke oddCh
// func numbers(oddCh chan<- int, evenCh chan<- int, errorCh chan<- string) {
// 	// Iterasi dari 1 hingga 22 (termasuk angka 21 dan 22)
// 	for i := 1; i <= 22; i++ {
// 		if i <= 20 {
// 			if i%2 == 0 {
// 				// Kirim angka genap ke evenCh
// 				evenCh <- i
// 			} else {
// 				// Kirim angka ganjil ke oddCh
// 				oddCh <- i
// 			}
// 		} else {
// 			// Kirim pesan error jika angka lebih dari 20
// 			errorCh <- fmt.Sprintf("error: number %d is greater than 20", i)
// 		}
// 	}

// 	// Tutup channel setelah selesai mengirim
// 	close(oddCh)
// 	close(evenCh)
// 	close(errorCh)
// }
