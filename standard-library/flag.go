package main

import (
	"flag"
	"fmt"
)

func main() {
	// Format: flag.Tipe(nama_flag, nilai_default, deskripsi)
	namaPtr := flag.String("nama", "Tamu", "Nama pengguna")
	umurPtr := flag.Int("umur", 20, "Umur pengguna")
	menikahPtr := flag.Bool("menikah", false, "Status pernikahan")

	flag.Parse()

	//(karena pointer, gunakan tanda bintang *)
	fmt.Printf("Halo %s, umur Anda %d. Menikah: %t\n", *namaPtr, *umurPtr, *menikahPtr)
}
