package dokter

import (
	"fmt"
	dokterStruct "projek/features/dokter/structs"
)

func LoginDokter(tDokter *dokterStruct.TabDokter) {
	var username, password string
	var found bool
	var i int
	fmt.Println("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Println("Masukkan password: ")
	fmt.Scan(&password)

	found = false
	i = 0
	for i < tDokter.N && !found {
		found = tDokter.ArrDokter[i].Username == username && tDokter.ArrDokter[i].Password == password
		i = i + 1

	}

	if found {
		fmt.Println("Selamat anda berhasil login")
	} else {
		fmt.Println("Username atau password yang anda masukkan salah")
		LoginDokter(tDokter)
	}
}
