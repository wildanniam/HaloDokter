package dokter

import (
	"fmt"
	dokterStruct "projek/features/dokter/structs"
)

/**
	- Fungsi login ini melakukan pengecekan berdasarkan username dan password
	- Jika username dan password benar, maka akan me return index dari dokter tersebut
	- Jika salah, maka akan me return index -1
	- Mengapa harus return index? karna index ini akan dipake untuk disimpan sebagai identitas yang sedang login
**/
func LoginDokter(tDokter *dokterStruct.TabDokter) int {
	var username, password string
	var found int = -1
	var i int
	fmt.Println("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Println("Masukkan password: ")
	fmt.Scan(&password)

	i = 0
	for i < tDokter.N && found == -1 {
		if tDokter.ArrDokter[i].Username == username && tDokter.ArrDokter[i].Password == password {
			found = i
		}
		i = i + 1
	}

	return found

}
