package main

import "fmt"

const NMAX int = 1000

type mahasiswa struct {
	nama, jurusan string
	nilai         int
	isDiterima    bool
}

type tabMhs [NMAX]mahasiswa

func main() {
	var masukan int
	var A tabMhs
	var n int
	for {
		fmt.Println("---------------")
		fmt.Println("Selamat Datang!")
		fmt.Println("Anda Login Sebagai: ")
		fmt.Println("1. Administrator")
		fmt.Println("2. Pendaftar")
		fmt.Println("3. Keluar")
		fmt.Print("Menu yang ingin dimasukan? ")
		fmt.Scanln(&masukan)
		fmt.Println("---------------")

		if masukan == 1 {
			password()
		} else if masukan == 2 {
			mainPendaftar(A, n)
		} else {
			fmt.Println("Terima Kasih!")
		}
	}
}

func adminPassword() {
	var pass string
	var A tabMhs
	var n int
	for {
		fmt.Println("---------------")
		fmt.Println("Masukan Password: ")
		fmt.Scan(&pass)
		fmt.Println("---------------")

		if pass == "sayaadmin" {
			mainAdmin(A, n)
		}
	}
}

func mainAdmin(A tabMhs, n int) {
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
			bacaData(&A, &n)
		} else if masukan == 2 {
			ubahData()
		} else if masukan == 3 {
			hapusData()
		} else if masukan == 4 {
			tampilData()
		} else if masukan == 5 {
			cariData()
		} else if masukan == 6 {
			urutData()
		} else {
			fmt.Println("Terima Kasih!")
		}
	}
}

func mainPendaftar() {
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
			tampilData()
		} else if masukan == 2 {
			cariData()
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
		fmt.Scan(&A[i].nama, &A[i].jurusan, &A[i].nilai)
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

func hapusData(A tabMhs) {
	var 
	for i := A; i < JumlahMahasiswa-1; i++ {
		DaftarMahasiswa[i] = DaftarMahasiswa[i+1]
	}
	JumlahMahasiswa--

	for i := A; i < JumlahJurusan-1; i++ {
		DaftarJurusan[i] = DaftarJurusan[i+1]
	}
	JumlahJurusan--
}

func tampilData(A tabMhs, n int) {
	fmt.Printf("%20s %6s %6s %6s\n", "Nama calon mahasiswa", "Jurusan calon mahasiswa", "Nilai calon mahasiswa")
	for i = 0; i < n; i++ {
		fmt.Printf("%20s %6s %6d\n", A[i].nama, A[i].jurusan, A[i].nilai)
	}
}
