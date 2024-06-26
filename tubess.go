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
	status        string
}
type tabMhs [NMAX]mahasiswa

func initDummyDataMhs(A *tabMhs, n *int) {
	A[0] = mahasiswa{103012300000, "Muhammad Aulia Khairu", "Informatika", 80, true, "Diterima"}
	A[1] = mahasiswa{103012300001, "Sanubari Nuraulia Legawa", "Informatika", 83, true, "Diterima"}
	A[2] = mahasiswa{103012300002, "Michael Jonatahan", "Teknik Elektro", 85, true, "Diterima"}
	A[3] = mahasiswa{103012300003, "Brian Anindya", "Kedokteran", 78, false, "Ditolak"}
	A[4] = mahasiswa{103012300004, "Nabil Azzamy", "Teknik Sipil", 75, false, "Ditolak"}
	A[5] = mahasiswa{103012300005, "Fairuztsani Kemal", "Informatika", 88, true, "Diterima"}
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
		fmt.Println("========================")
		fmt.Println("Selamat Datang!")
		fmt.Println("Anda Login Sebagai: ")
		fmt.Println("1. Administrator")
		fmt.Println("2. Pendaftar")
		fmt.Println("3. Keluar")
		fmt.Println("========================")
		fmt.Print("Masukan Pilihan Anda: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			adminPassword(&A, &n)
		} else if masukan == 2 {
			mainPendaftar(&A, &n)
		} else if masukan == 3 {
			fmt.Println("Terima Kasih!")
			os.Exit(0)
		} else {

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
		fmt.Println("7. Kembali")
		fmt.Println("8. Keluar")
		fmt.Println("===================================")
		fmt.Print("Masukan Pilihan Anda: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			bacaData(A, n)
		} else if masukan == 2 {
			ubahData(A, n)
		} else if masukan == 3 {
			var noPendaftar int
			fmt.Println("--------------------------------------------")
			fmt.Print("Masukan Nomor pendaftar yang ingin dihapus: ")
			fmt.Scanln(&noPendaftar)
			hapusDataMahasiswa(A, n, noPendaftar)
		} else if masukan == 4 {
			tampilData(*A, *n)
		} else if masukan == 5 {
			selectCariData(A, n)
		} else if masukan == 6 {
			selectUrutData(A, *n)
		} else if masukan == 7 {
			main()
		} else if masukan == 8 {
			fmt.Println("Terima Kasih!")
			os.Exit(0)
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
		fmt.Println("3. Urut data calon mahasiswa")
		fmt.Println("4. Kembali")
		fmt.Println("5. Keluar")
		fmt.Println("===================================")
		fmt.Print("Masukan Pilihan Anda: ")
		fmt.Scanln(&masukan)

		if masukan == 1 {
			tampilData(*A, *n)
		} else if masukan == 2 {
			selectCariData(A, n)
		} else if masukan == 3 {
			selectUrutData(A, *n)
		} else if masukan == 4 {
			main()
		} else if masukan == 5 {
			fmt.Println("Terima Kasih!")
			os.Exit(0)
		} else {

		}
	}
}

func bacaData(A *tabMhs, n *int) {
	if *n == NMAX {
		fmt.Println("-------------------------------------------------")
		fmt.Println("Maaf, kapasitas data calon mahasiswa sudah penuh!")
	} else {
		fmt.Println("-----------------------------")
		fmt.Print("Masukan nomor pendaftar: ")
		fmt.Scanln(&A[*n].noPendaftar)
		fmt.Print("Masukan nama calon mahasiswa: ")
		InputlnString(&A[*n].nama)
		fmt.Print("Masukan nama jurusan: ")
		InputlnString(&A[*n].jurusan)
		fmt.Print("Masukan nilai calon mahasiswa: ")
		fmt.Scanln(&A[*n].nilai)

		for !(A[*n].nilai >= 0 && A[*n].nilai <= 100) {
			fmt.Println("Nilai yang dimasukan tidak valid!")

			fmt.Print("Masukan nilai calon mahasiswa: ")
			fmt.Scanln(&A[*n].nilai)
		}

		if A[*n].nilai >= 80 && A[*n].nilai <= 100 {
			A[*n].isDiterima = true
			A[*n].status = "Diterima"
		} else {
			A[*n].isDiterima = false
			A[*n].status = "Ditolak"
		}
		*n++

		fmt.Println("Data calon mahasiswa berhasil ditambah!")
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
		fmt.Println("Data dengan nomor pendaftar tersebut tidak ditemukan!")
		return
	}

	fmt.Println("--------------------")
	fmt.Println("Masukan data baru:")
	fmt.Print("Nama: ")
	InputlnString(&A[indexToEdit].nama)
	fmt.Print("Jurusan: ")
	InputlnString(&A[indexToEdit].jurusan)
	fmt.Print("Nilai: ")
	fmt.Scanln(&A[indexToEdit].nilai)

	for !(A[indexToEdit].nilai >= 0 && A[indexToEdit].nilai <= 100) {
		fmt.Println("Nilai yang dimasukan tidak valid!")

		fmt.Print("Masukan lagi nilai calon mahasiswa: ")
		fmt.Scanln(&A[indexToEdit].nilai)
	}

	if A[indexToEdit].nilai >= 85 && A[indexToEdit].nilai <= 100 {
		A[indexToEdit].isDiterima = true
		A[indexToEdit].status = "Diterima"
	} else {
		A[indexToEdit].isDiterima = false
		A[indexToEdit].status = "Ditolak"
	}
	fmt.Println("Data calon mahasiswa berhasil diubah!")
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
		fmt.Println("Data calon mahasiswa berhasil dihapus!")

	} else {
		fmt.Println("Data dengan nomor pendaftar tersebut tidak ditemukan!")
	}
}

func tampilData(A tabMhs, n int) {
	fmt.Println("Data Calon Mahasiswa:")
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Printf("%-21s %-23s %-14s %-7s %s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai", "Status")
	fmt.Println("-----------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%-15d %-27s %-17s %-5d %-8s\n", A[i].noPendaftar, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
	}
	fmt.Println("-----------------------------------------------------------------------------")
}

func selectCariData(A *tabMhs, n *int) {
	var ind int

	fmt.Println("=============================")
	fmt.Println("Silahkan Pilih Menu Berikut: ")
	fmt.Println("Cari berdasarkan: ")
	fmt.Println("1. Nomor Pendaftar")
	fmt.Println("2. Nama")
	fmt.Println("3. Jurusan")
	fmt.Println("4. Status")
	fmt.Println("-----------------------------")
	fmt.Println("5. Kembali")
	fmt.Println("6. Keluar")
	fmt.Println("=============================")
	fmt.Print("Masukan Pilihan Anda: ")
	fmt.Scanln(&ind)
	if ind == 1 {
		cariNoPes(A, *n)
	} else if ind == 2 {
		cariNama(*A, *n)
	} else if ind == 3 {
		cariJurusan(*A, *n)
	} else if ind == 4 {
		cariStatus(*A, *n)
	} else if ind == 5 {
		main()
	} else if ind == 6 {
		fmt.Println("Terima Kasih!")
		os.Exit(0)
	} else {

	}

}

func cariNoPes(A *tabMhs, n int) {
	var temp mahasiswa
	var pass, i int
	var no int
	var left int = 0
	var right int = n - 1
	var mid int
	var found bool = false
	var foundIndex int = -1

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[i]
		for i > 0 && temp.noPendaftar > A[i-1].noPendaftar {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}

	fmt.Println("------------------------------------------")
	fmt.Print("Masukan nomor pendaftar yang akan dicari: ")
	fmt.Scanln(&no)

	for left <= right && !found {
		mid = (left + right) / 2
		if no > A[mid].noPendaftar {
			right = mid - 1
		} else if no < A[mid].noPendaftar {
			left = mid + 1
		} else {
			found = true
			foundIndex = mid
		}
	}

	if foundIndex == -1 {
		fmt.Println("Data nomor pendaftar tersebut tidak ditemukan!")
	} else {
		fmt.Println("Data calon mahasiswa ditemukan: ")
		fmt.Println("-----------------------------------------------------------------------------")
		fmt.Printf("%-22s %-25s %-13s %s %-8s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai", "Status")
		fmt.Println("-----------------------------------------------------------------------------")
		fmt.Printf("%-15d %-30s %-15s %3d %-12s\n", A[foundIndex].noPendaftar, A[foundIndex].nama, A[foundIndex].jurusan, A[foundIndex].nilai, A[foundIndex].status)
		fmt.Println("-----------------------------------------------------------------------------")
	}

}

func cariNama(A tabMhs, n int) {
	var nama string
	fmt.Println("-----------------------------------------")
	fmt.Print("Masukan nama pendaftar yang akan dicari: ")
	InputlnString(&nama)
	var found bool

	for i := 0; i < n; i++ {
		if A[i].nama == nama {
			found = true
			fmt.Println("Data calon mahasiswa ditemukan:")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-21s %-23s %-14s %-7s %s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai", "Status")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-15d %-27s %-17s %-5d %-8s\n", A[i].noPendaftar, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
			fmt.Println("-----------------------------------------------------------------------------")
		}
	}
	if !found {
		fmt.Println("Data nama pendaftar tersebut tidak ditemukan!")
	}
}

func cariJurusan(A tabMhs, n int) {
	var jurusan string
	fmt.Println("--------------------------------------------")
	fmt.Print("Masukan jurusan pendaftar yang akan dicari: ")
	InputlnString(&jurusan)
	var found bool

	for i := 0; i < n; i++ {
		if A[i].jurusan == jurusan {
			found = true
			fmt.Println("Data calon mahasiswa ditemukan:")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-21s %-23s %-14s %-7s %s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai", "Status")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-15d %-27s %-17s %-5d %-8s\n", A[i].noPendaftar, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
			fmt.Println("-----------------------------------------------------------------------------")
		}
	}
	if !found {
		fmt.Println("Data jurusan pendaftar tersebut tidak ditemukan!")
	}
}

func cariStatus(A tabMhs, n int) {
	var status bool
	fmt.Println("---------------------------------")
	fmt.Print("Masukan status yang akan dicari: ")
	fmt.Scanln(&status)
	var found bool

	for i := 0; i < n; i++ {
		if A[i].isDiterima == status {
			found = true
			fmt.Println("Data calon mahasiswa ditemukan:")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-21s %-23s %-14s %-7s %s\n", "Nomor Pendaftar", "Nama", "Jurusan", "Nilai", "Status")
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("%-15d %-27s %-17s %-5d %-8s\n", A[i].noPendaftar, A[i].nama, A[i].jurusan, A[i].nilai, A[i].status)
			fmt.Println("-----------------------------------------------------------------------------")
		}
	}
	if !found {
		fmt.Println("Data status pendaftar tersebut tidak ditemukan!")
	}
}

func selectUrutData(A *tabMhs, n int) {
	var inp int

	fmt.Println("=============================")
	fmt.Println("Silahkan Pilih Menu Berikut: ")
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Jurusan")
	fmt.Println("3. Nilai")
	fmt.Println("-----------------------------")
	fmt.Println("4. Kembali")
	fmt.Println("5. Keluar")
	fmt.Println("=============================")
	fmt.Print("Masukkan Pilihan Anda: ")
	fmt.Scanln(&inp)

	if inp == 1 {
		selUrutNama(A, n)
	} else if inp == 2 {
		selUrutJurusan(A, n)
	} else if inp == 3 {
		selUrutNilai(A, n)
	} else if inp == 4 {
		main()
	} else if inp == 5 {
		fmt.Println("Terima Kasih!")
		os.Exit(0)
	} else {

	}
}

func selUrutNama(A *tabMhs, n int) {
	var masuk int

	fmt.Println("=============================")
	fmt.Println("Silahkan Pilih Menu Berikut: ")
	fmt.Println("Urutkan secara:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("-----------------------------")
	fmt.Println("3. Kembali")
	fmt.Println("4. Keluar")
	fmt.Println("=============================")
	fmt.Print("Masukan Pilihan Anda: ")
	fmt.Scanln(&masuk)

	if masuk == 1 {
		namaAsc(A, n)
	} else if masuk == 2 {
		namaDsc(A, n)
	} else if masuk == 3 {
		selectUrutData(A, n)
	} else if masuk == 4 {
		fmt.Println("Terima Kasih!")
		os.Exit(0)
	} else {

	}
}

func selUrutJurusan(A *tabMhs, n int) {
	var masuk int

	fmt.Println("=============================")
	fmt.Println("Silahkan Pilih Menu Berikut: ")
	fmt.Println("Urutkan secara:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("-----------------------------")
	fmt.Println("3. Kembali")
	fmt.Println("4. Keluar")
	fmt.Println("=============================")
	fmt.Print("Masukan Pilihan Anda: ")
	fmt.Scanln(&masuk)

	if masuk == 1 {
		jurusanAsc(A, n)
	} else if masuk == 2 {
		jurusanDsc(A, n)
	} else if masuk == 3 {
		selectUrutData(A, n)
	} else if masuk == 4 {
		fmt.Println("Terima Kasih!")
		os.Exit(0)
	} else {

	}
}

func selUrutNilai(A *tabMhs, n int) {
	var masuk int

	fmt.Println("=============================")
	fmt.Println("Silahkan Pilih Menu Berikut: ")
	fmt.Println("Urutkan secara:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("-----------------------------")
	fmt.Println("3. Kembali")
	fmt.Println("4. Keluar")
	fmt.Println("=============================")
	fmt.Print("Masukan Pilihan Anda: ")
	fmt.Scanln(&masuk)

	if masuk == 1 {
		nilaiAsc(A, n)
	} else if masuk == 2 {
		nilaiDsc(A, n)
	} else if masuk == 3 {
		selectUrutData(A, n)
	} else if masuk == 4 {
		fmt.Println("Terima Kasih!")
		os.Exit(0)
	} else {

	}
}

func namaAsc(A *tabMhs, n int) {
	for i := 0; i < n-1; i++ {
		least := i

		for j := i + 1; j < n; j++ {
			if A[j].nama < A[least].nama {
				least = j
			}
		}

		A[i], A[least] = A[least], A[i]
	}

	tampilData(*A, n)

}

func jurusanAsc(A *tabMhs, n int) {
	for i := 0; i < n-1; i++ {
		least := i

		for j := i + 1; j < n; j++ {
			if A[j].jurusan < A[least].jurusan {
				least = j
			}
		}

		A[i], A[least] = A[least], A[i]
	}

	tampilData(*A, n)

}

func nilaiAsc(A *tabMhs, n int) {
	for i := 0; i < n-1; i++ {
		least := i

		for j := i + 1; j < n; j++ {
			if A[j].nilai < A[least].nilai {
				least = j
			}
		}

		A[i], A[least] = A[least], A[i]
	}

	tampilData(*A, n)

}

func namaDsc(A *tabMhs, n int) {
	for i := 0; i < n; i++ {
		j := i

		for j > 0 {
			if A[j-1].nama < A[j].nama {
				A[j-1], A[j] = A[j], A[j-1]
			}

			j--

		}
	}

	tampilData(*A, n)
}

func jurusanDsc(A *tabMhs, n int) {
	for i := 0; i < n; i++ {
		j := i

		for j > 0 {
			if A[j-1].jurusan < A[j].jurusan {
				A[j-1], A[j] = A[j], A[j-1]
			}

			j--

		}
	}
	tampilData(*A, n)
}

func nilaiDsc(A *tabMhs, n int) {
	for i := 0; i < n; i++ {
		j := i

		for j > 0 {
			if A[j-1].nilai < A[j].nilai {
				A[j-1], A[j] = A[j], A[j-1]
			}

			j--

		}
	}
	tampilData(*A, n)
}
