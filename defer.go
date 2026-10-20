package main

import "fmt"

func logging() {
	fmt.Println("End App")
	message := recover() // ambil pesan error dari panic
	fmt.Println("Panic", message)
}
func runApp() {
	defer logging() // ini dijalankan di akhir function selesai
}

func startApp(error bool) {
	defer logging()
	if error {
		panic("Error bro")
	}
}
func main() {
	fmt.Println("App Start")
	runApp()
	startApp(true)
}
