package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1000

type mahasiswa struct {
	noPendaftar   int
	nama, jurusan string
	nilai         int
	isDiterima    bool
}
type tabMhs [NMAX]mahasiswa

func initDummyDataMhs(A *tabMhs, n *int) {
	A[0] = mahasiswa{103012300000, "Muhammad Aulia Khairu", "Informatika", 80, true}
	A[1] = mahasiswa{103012300001, "Sanubari Nuraulia Legawa", "Informatika", 83, true}
	A[2] = mahasiswa{103012300002, "Michael Jonatahan", "Teknik Elektro", 85, true}
	A[3] = mahasiswa{103012300003, "Brian Anindya", "Kedokteran", 78, false}
	A[4] = mahasiswa{103012300004, "Nabil Azzamy", "Teknik Sipil", 75, false}
	A[5] = mahasiswa{103012300005, "Fairuztsani Kemal", "Informatika", 88, true}
	*n = 6
}

func InputlnString(ptr *string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	*ptr = strings.TrimSpace(input)
}

func main() {
	var masukan int
	var A tabMhs
	var n int

	initDummyDataMhs(&A, &n)
	for {
		fmt.Println("=========================")
		fmt.Println("Selamat Datang!")
		fmt.Println("Anda Login Sebagai: ")
		fmt.Println("1. Administrator")
		fmt.Println("2. Pendaftar")
		fmt.Println("=========================")
		fmt.Print("Menu yang ingin dimasukan: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			adminPassword(&A, &n)
		} else if masukan == 2 {
			mainPendaftar(&A, &n)
		} else {
			fmt.Println("Terima kasih")
		}
	}
}

func adminPassword(A *tabMhs, n *int) {
	var pass string
	for {
		fmt.Print("Password: ")
		fmt.Scanln(&pass)

		if pass == "sayaadmin" {
			mainAdmin(A, n)
			return
		}
		fmt.Println("Maaf password yang anda masukan salah, silahkan coba lagi!")
		fmt.Println("==========================================================")
	}
}

func mainAdmin(A *tabMhs, n *int) {
	var masukan int
	for {
		fmt.Println("===================================")
		fmt.Println("Selamat Datang Admin!")
		fmt.Println("Silahkan Pilih Menu Berikut: ")
		fmt.Println("1. Tambah data calon mahasiswa")
		fmt.Println("2. Ubah data calon mahasiswa")
		fmt.Println("3. Hapus data calon mahasiswa")
		fmt.Println("4. Tampilkan data calon mahasiswa")
		fmt.Println("5. Cari data calon mahasiswa")
		fmt.Println("6. Urutkan data calon mahasiswa")
		fmt.Println("7. Keluar")
		fmt.Println("===================================")
		fmt.Print("Menu yang ingin dimasukan: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			bacaData(A, n)
		} else if masukan == 2 {
			ubahData(A, n)
		} else if masukan == 3 {
			var noPendaftar int
			fmt.Print("Nomor pendaftar calon mahasiswa yang ingin dihapus: ")
			fmt.Scanln(&noPendaftar)
			hapusDataMahasiswa(A, n, noPendaftar)
		} else if masukan == 4 {
			tampilData(*A, *n)
		} else if masukan == 5 {
			// cariData()
		} else if masukan == 6 {
			// urutData()
		} else if masukan == 7 {
			fmt.Println("Terima kasih!")
			main()
		} else {

		}
	}
}

func mainPendaftar(A *tabMhs, n *int) {
	var masukan int
	for {
		fmt.Println("===================================")
		fmt.Println("Selamat Datang Calon Mahasiswa!")
		fmt.Println("Silahkan Pilih Menu Berikut: ")
		fmt.Println("1. Tampilkan data calon mahasiswa")
		fmt.Println("2. Cari data calon mahasiswa")
		fmt.Println("3. Keluar")
		fmt.Println("===================================")
		fmt.Print("Menu yang ingin dimasukan: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			tampilData(*A, *n)
		} else if masukan == 2 {
			// cariData()
		} else if masukan == 3 {
			fmt.Println("Terima Kasih!")
			main()
		} else {

		}
	}
}

func bacaData(A *tabMhs, n *int) {
	if *n == NMAX {
		fmt.Println("Maaf, kapasitas data calon mahasiswa sudah penuh.")
	} else {
		fmt.Print("Masukkan nomor pendaftar: ")
		fmt.Scanln(&A[*n].noPendaftar)
		fmt.Print("Masukkan nama calon mahasiswa: ")
		InputlnString(&A[*n].nama)
		fmt.Print("Masukkan nama jurusan: ")
		InputlnString(&A[*n].jurusan)
		fmt.Print("Masukkan nilai calon mahasiswa: ")
		fmt.Scanln(&A[*n].nilai)

		for !(A[*n].nilai >= 0 && A[*n].nilai <= 100) {
			fmt.Println("Nilai yang dimasukan tidak valid")

			fmt.Print("Masukkan nilai calon mahasiswa: ")
			fmt.Scanln(&A[*n].nilai)
		}

		if A[*n].nilai >= 80 && A[*n].nilai <= 100 {
			A[*n].isDiterima = true
		} else {
			A[*n].isDiterima = false
		}
		*n++

		fmt.Println("Data berhasil ditambah!")
	}
}

func doSequentialSearch(A *tabMhs, n *int, noPendaftarToFind int) int {
	for i := 0; i < *n; i++ {
		if A[i].noPendaftar == noPendaftarToFind {
			return i
		}
	}
	return -1
}

func ubahData(A *tabMhs, n *int) {
	var noPendaftar int
	fmt.Print("Masukan nomor pendaftar calon mahasiswa yang ingin diubah: ")
	fmt.Scanln(&noPendaftar)

	var indexToEdit int = doSequentialSearch(A, n, noPendaftar)
	if indexToEdit == -1 {
		fmt.Println("Data dengan nomor pendaftar tersebut tidak ditemukan.")
		return
	}

	fmt.Println("------------------------------")
	fmt.Println("Masukkan data baru:")
	fmt.Print("Nama: ")
	InputlnString(&A[indexToEdit].nama)
	fmt.Print("Jurusan: ")
	InputlnString(&A[indexToEdit].jurusan)
	fmt.Print("Nilai: ")
	fmt.Scanln(&A[indexToEdit].nilai)

	for !(A[indexToEdit].nilai >= 0 && A[indexToEdit].nilai <= 100) {
		fmt.Println("Nilai yang dimasukan tidak valid!")

		fmt.Print("Masukkan lagi nilai calon mahasiswa: ")
		fmt.Scanln(&A[indexToEdit].nilai)
	}

	if A[indexToEdit].nilai >= 85 && A[indexToEdit].nilai <= 100 {
		A[indexToEdit].isDiterima = true
	} else {
		A[indexToEdit].isDiterima = false
	}
	fmt.Println("Data berhasil diubah.")
}

func hapusDataMahasiswa(A *tabMhs, nMhs *int, noPendaftar int) {
	var index int
	found := false

	for i := 0; i < *nMhs-1 || !found; i++ {
		if A[i].noPendaftar == noPendaftar {
			index = i
			found = true
		}
	}

	if found {
		for i := index; i < *nMhs-1; i++ {
			A[i] = A[i+1]
		}
		*nMhs--
		fmt.Println("Data calon mahasiswa berhasil dihapus.")

	} else {
		fmt.Println("Data dengan nomor pendaftar tersebut tidak ditemukan.")
	}
}

func tampilData(A tabMhs, n int) {
	fmt.Println("Data Calon Mahasiswa:")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%-22s %-25s %-13s %s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai")
	fmt.Println("---------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%-15d %-30s %-15s %3d\n", A[i].noPendaftar, A[i].nama, A[i].jurusan, A[i].nilai)
	}
	fmt.Println("---------------------------------------------------------------------")
}
