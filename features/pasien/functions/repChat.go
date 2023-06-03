package pasien

import (
	"fmt"
	common "projek/common"
	pasienStruct "projek/features/pasien/structs"
)

func repChat(tPost *pasienStruct.TabPost) {
	//Pengguna menginputkan balasan pada postingan
	var noPost int
	fmt.Println("Pilih postingan yang ingin dibalas: ")
	fmt.Scan(&noPost)
	fmt.Println("Balasan anda: ")
	common.InputMultipleString(&tPost.ArrPost[noPost+1].TxtReply)
	fmt.Println("Balasan anda telah diposting!")
}
