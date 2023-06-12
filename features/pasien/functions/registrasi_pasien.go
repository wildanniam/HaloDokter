package pasien

import (
	"fmt"
	common "projek/common"
	pasienStruct "projek/features/pasien/structs"
)

func RegistrasiPasien(arrPasien *pasienStruct.TabPasien) {
	fmt.Print("Masukkan nama : ")
	common.InputMultipleString(&arrPasien.ArrPasien[arrPasien.N].Nama)
	fmt.Print("Masukkan username : ")
	common.InputMultipleString(&arrPasien.ArrPasien[arrPasien.N].Username)
	fmt.Print("Masukkan password : ")
	common.InputMultipleString(&arrPasien.ArrPasien[arrPasien.N].Password)
	arrPasien.ArrPasien[arrPasien.N].ID = arrPasien.N
	arrPasien.N++

	common.ResetConsole()

	fmt.Println("=======================================================================================")
	fmt.Println("                           Berhasil Registrasi Sebagai Pasien                          ")
	fmt.Println("=======================================================================================")
	fmt.Println()

	common.ShowEndAction()
}
