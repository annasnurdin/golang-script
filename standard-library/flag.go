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

	// -h untuk cek deskripsi flag yang sudah di daftarkan
	// --nama => langsung ditulis dengan spasi
	// -nama=Annas => kalau 1 dash nya pakai sama dengan tanpa petik/tergantung, ding
}
