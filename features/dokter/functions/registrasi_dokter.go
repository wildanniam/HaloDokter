package dokter

import (
	"fmt"
	common "projek/common"
	dokterStruct "projek/features/dokter/structs"
)

func RegisDokter(tDokter *dokterStruct.TabDokter) {
	fmt.Print("Masukkan nama : ")
	common.InputMultipleString(&tDokter.ArrDokter[tDokter.N].Nama)
	fmt.Print("Spesialis anda? : ")
	common.InputMultipleString(&tDokter.ArrDokter[tDokter.N].Spesialis)
	fmt.Print("Masukkan username : ")
	common.InputMultipleString(&tDokter.ArrDokter[tDokter.N].Username)
	fmt.Print("Masukkan password : ")
	common.InputMultipleString(&tDokter.ArrDokter[tDokter.N].Password)

}
