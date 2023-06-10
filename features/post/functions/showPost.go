package post

import (
	"fmt"
	"projek/common"
	doctorStruct "projek/features/dokter/structs"
	patientStruct "projek/features/pasien/structs"
	postStruct "projek/features/post/structs"
)

func ShowPost(tPost *postStruct.TabPost, tDoctor *doctorStruct.TabDokter, tPatient *patientStruct.TabPasien) {
	//Menampilkan isi postingan yang telah diinputkan pengguna
	common.ResetConsole()

	for i := 0; i < tPost.N; i++ {
		fmt.Printf("ID Postingan :%s\n%s \nTag : %s \n", tPost.ArrPost[i].ID, tPost.ArrPost[i].TxtAddPost, tPost.ArrPost[i].TagPost)
		for j := 0; j < tPost.ArrPost[i].Nreply; j++ {
			// Pengecekan tipe user
			if tPost.ArrPost[i].ArrReply[j].User == "dokter" {

				doctor := tDoctor.ArrDokter[tPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("Nama Dokter: ", doctor.Nama)
				fmt.Println("Spesialis: ", doctor.Spesialis)

			} else if tPost.ArrPost[i].ArrReply[j].User == "pasien" {

				patient := tPatient.ArrPasien[tPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("Nama Pasien: ", patient.Nama)

			}
			fmt.Println("Balasan : ", tPost.ArrPost[i].ArrReply[j].Message)
		}
		fmt.Println()

	}

}
