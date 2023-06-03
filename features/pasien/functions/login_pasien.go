package pasien

import (
	"fmt"
	pasienStruct "projek/features/pasien/structs"
)

func LoginPasien(tPasien *pasienStruct.TabPasien) {
	var username, password string
	var found bool
	var i int
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	found = false
	i = 0
	for i < tPasien.N && !found {
		found = tPasien.ArrPasien[i].Username == username && tPasien.ArrPasien[i].Password == password
		i = i + 1

	}

	if found {
		fmt.Println("Selamat anda berhasil login")
	} else {
		fmt.Println("Username atau password yang anda masukkan salah")
	}
}
