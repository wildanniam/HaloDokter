package dokter

import (
	common "projek/common"
	dokterStruct "projek/features/dokter/structs"
)

func InputDoctor(arrDokter *dokterStruct.TabDokter) {
	arrDokter.ArrDokter[0].Nama = "Indra Mahesa"
	arrDokter.ArrDokter[0].ID = common.GenerateRandomString(3)
	arrDokter.ArrDokter[0].Username = "indramahesa"
	arrDokter.ArrDokter[0].Password = "indramahesa"
	arrDokter.ArrDokter[0].Spesialis = "Dokter Umum"
	arrDokter.N++

	arrDokter.ArrDokter[1].Nama = "Zahra Amiera Putri"
	arrDokter.ArrDokter[1].ID = common.GenerateRandomString(3)
	arrDokter.ArrDokter[1].Username = "zahra"
	arrDokter.ArrDokter[1].Password = "zahra"
	arrDokter.ArrDokter[1].Spesialis = "Dokter Anak"
	arrDokter.N++

	arrDokter.ArrDokter[2].Nama = "Yazid Al Ghozali"
	arrDokter.ArrDokter[2].ID = common.GenerateRandomString(3)
	arrDokter.ArrDokter[2].Username = "yazidal"
	arrDokter.ArrDokter[2].Password = "yazidal"
	arrDokter.ArrDokter[2].Spesialis = "THT"
	arrDokter.N++

	arrDokter.ArrDokter[3].Nama = "Aaron Joseph"
	arrDokter.ArrDokter[3].ID = common.GenerateRandomString(3)
	arrDokter.ArrDokter[3].Username = "aaron"
	arrDokter.ArrDokter[3].Password = "aaron"
	arrDokter.ArrDokter[3].Spesialis = "Anestesi"

	arrDokter.N++
}
