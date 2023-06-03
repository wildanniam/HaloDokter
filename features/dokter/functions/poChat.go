package dokter

import (
	"fmt"
	common "projek/common"
	dokterStruct "projek/features/dokter/structs"
)

func poChat(tPost *dokterStruct.TabPost) {
	//Pengguna menginputkan postingan kedalam aplikasi
	if tPost.N < dokterStruct.NMAX {
		fmt.Print("Posting pertanyaan anda: ")
		common.InputMultipleString(&tPost.ArrPost[tPost.N].TxtAddPost)
		fmt.Print("Tag: ")
		fmt.Scan(&tPost.ArrPost[tPost.N].TagPost)
		tPost.N++
		fmt.Print("Pertanyaan anda telah diposting!")
	}
}
