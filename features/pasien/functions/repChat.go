package pasien

import (
	"fmt"
	postStruct "projek/features/post/structs"
)

func RepChat(tPost *postStruct.TabPost) {
	//Pengguna menginputkan balasan pada postingan
	var noPost int
	fmt.Println("Pilih postingan yang ingin dibalas: ")
	fmt.Scan(&noPost)
	fmt.Println("Balasan anda: ")
	// common.InputMultipleString(&tPost.ArrPost[noPost-1].ArrReply[tPost.ArrPost[noPost-1].ArrReply])
	fmt.Println("Balasan anda telah diposting!")
}
