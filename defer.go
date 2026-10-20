package main

import "fmt"

func logging() {
	fmt.Println("End App")
}
func runApp() {
	defer logging() // ini dijalankan di akhir function selesai
}
func main() {
	fmt.Println("App Start")
	runApp()
}
