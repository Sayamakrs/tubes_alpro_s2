package main

import "fmt"

const NMAX int = 1000

type mahasiswa struct {
	nim           string
	nama, jurusan string
	nilai         int
	isDiterima    bool
}

type tabMhs [NMAX]mahasiswa

func initDummyDataMhs(A *tabMhs, n *int) {
	A[0] = mahasiswa{"103012300000", "Aulia Khairu Rizqy Sembiring", "Informatika", 80, true}
	A[1] = mahasiswa{"103012300001", "Mikel Jonatahan", "Teknik Elektro", 95, true}
	A[2] = mahasiswa{"103012300002", "Brian Anindya", "Kedokteran", 65, false}

	*n = 3
}

func main() {
	var masukan int
	var A tabMhs
	var n int

	initDummyDataMhs(&A, &n)

	for {
		fmt.Println("---------------")
		fmt.Println("Selamat Datang!")
		fmt.Println("Anda Login Sebagai: ")
		fmt.Println("1. Administrator")
		fmt.Println("2. Pendaftar")
		fmt.Println("3. Keluar")
		fmt.Println("---------------")
		fmt.Print("Menu yang ingin dimasukan? ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			adminPassword(&A, &n)
		} else if masukan == 2 {
			mainPendaftar(&A, &n)
		} else {
			fmt.Println("Terima Kasih!")
		}
	}
}

func adminPassword(A *tabMhs, n *int) {
	var pass string
	for {
		fmt.Println("---------------")
		fmt.Println("Masukan Password: ")
		fmt.Scanln(&pass)
		fmt.Println("---------------")

		if pass == "sayaadmin" {
			mainAdmin(A, n)
		}
	}
}

func mainAdmin(A *tabMhs, n *int) {
	var masukan int
	for {
		fmt.Println("---------------")
		fmt.Println("Selamat Datang Admin!")
		fmt.Println("Silahkan Pilih Menu Berikut: ")
		fmt.Println("1. Tambah data calon mahasiswa")
		fmt.Println("2. Ubah data calon mahasiswa")
		fmt.Println("3. Hapus data calon mahasiswa")
		fmt.Println("4. Tampilkan data calon mahasiswa")
		fmt.Println("5. Cari data calon mahasiswa")
		fmt.Println("6. Urutkan data calon mahasiswa")
		fmt.Println("7. Keluar")
		fmt.Print("Menu yang ingin dimasukan: ")
		fmt.Scanln(&masukan)
		fmt.Println("---------------")

		if masukan == 1 {
			bacaData(A, n)
		} else if masukan == 2 {
			ubahData()
		} else if masukan == 3 {
			var nim string
			fmt.Print("NIM yang akan dihapus: ")
			fmt.Scanln(&nim)

			hapusDataMahasiswa(A, n, nim)
		} else if masukan == 4 {
			tampilData(*A, *n)
		} else if masukan == 5 {
			// cariData()
		} else if masukan == 6 {
			// urutData()
		} else {
			fmt.Println("Terima Kasih!")
		}
	}
}

func mainPendaftar(A *tabMhs, n *int) {
	var masukan int

	for {
		fmt.Println("---------------")
		fmt.Println("Selamat Datang Calon Mahasiswa!")
		fmt.Println("Silahkan Pilih Menu Berikut: ")
		fmt.Println("1. Tampilkan data calon mahasiswa")
		fmt.Println("2. Cari data calon mahasiswa")
		fmt.Println("3. Keluar")
		fmt.Print("Menu yang ingin dimasukan: ")
		fmt.Scanln(&masukan)
		fmt.Println("---------------")

		if masukan == 1 {
			tampilData(*A, *n)
		} else if masukan == 2 {
			// cariData()
		} else {
			fmt.Println("Terima Kasih!")
		}
	}
}

func bacaData(A *tabMhs, n *int) {
	if *n > NMAX {
		*n = NMAX
	}
	fmt.Println("Masukan data calon mahasiswa")
	for i := 0; i < *n; i++ {
		fmt.Scanln(&A[i].nama, &A[i].jurusan, &A[i].nilai)
		if A[i].nilai >= 85 && A[i].nilai <= 100 {
			A[i].isDiterima = true
		} else if A[i].nilai < 85 {
			A[i].isDiterima = false
		} else {
			A[i].isDiterima = false
		}
	}
}

func ubahData() {

}

func hapusDataMahasiswa(A *tabMhs, nMhs *int, nim string) {
	var index int

	// Bisa dijadiin searching sequentially
	for i := 0; i < *nMhs-1; i++ {
		if A[i].nim == nim {
			index = i
		}
	}

	for i := index; i < *nMhs-1; i++ {
		A[i] = A[i+1]
	}

	*nMhs--
}

// Fungsi tampilData untuk menampilkan data calon mahasiswa
func tampilData(A tabMhs, n int) {
	fmt.Println("Data Calon Mahasiswa")
	fmt.Printf("%-13s %-30s %-20s %3s\n", "NIM", "Nama", "Jurusan", "Nilai")
	for i := 0; i < n; i++ {
		fmt.Printf("%-13s %-30s %-20s %3d\n", A[i].nim, A[i].nama, A[i].jurusan, A[i].nilai)
	}
}