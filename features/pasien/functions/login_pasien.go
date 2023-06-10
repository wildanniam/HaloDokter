package pasien

import (
	"fmt"
	pasienStruct "projek/features/pasien/structs"
)

func LoginPasien(tPasien *pasienStruct.TabPasien) int {
	var username, password string
	var found int = -1
	var i int
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	i = 0
	for i < tPasien.N && found == -1 {
		if tPasien.ArrPasien[i].Username == username && tPasien.ArrPasien[i].Password == password {
			found = i
		}
		i = i + 1
	}

	return found
}
