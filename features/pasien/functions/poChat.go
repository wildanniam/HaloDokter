package pasien

import (
	"fmt"
	common "projek/common"
	pasienStruct "projek/features/pasien/structs"
)

func poChat(tPost *pasienStruct.TabPost) {
	//Pengguna menginputkan postingan kedalam aplikasi
	if tPost.N < pasienStruct.NMAX {
		fmt.Print("Posting pertanyaan anda: ")
		common.InputMultipleString(&tPost.ArrPost[tPost.N].TxtAddPost)
		fmt.Print("Tag: ")
		fmt.Scan(&tPost.ArrPost[tPost.N].TagPost)
		tPost.N++
		fmt.Print("Pertanyaan anda telah diposting!")
	}
}
