package pasien

import (
	"fmt"
	common "projek/common"
	pasienStruct "projek/features/pasien/structs"
)

func RegisPasien(tPasien *pasienStruct.TabPasien) {
	fmt.Print("Masukkan nama : ")
	common.InputMultipleString(&tPasien.ArrPasien[tPasien.N].Nama)
	fmt.Print("Masukkan username : ")
	common.InputMultipleString(&tPasien.ArrPasien[tPasien.N].Username)
	fmt.Print("Masukkan password : ")
	common.InputMultipleString(&tPasien.ArrPasien[tPasien.N].Password)

	tPasien.N++
}
