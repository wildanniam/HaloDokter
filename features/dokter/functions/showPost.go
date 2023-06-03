package dokter

import (
	"fmt"
	dokterStruct "projek/features/dokter/structs"
)

func showPost(tPost *dokterStruct.TabPost) {
	//Menampilkan isi postingan yang telah diinputkan pengguna
	for i := 0; i < tPost.N; i++ {
		fmt.Printf("Post %d: %s \nTag : %s \n", i+1, tPost.ArrPost[i].TxtAddPost, tPost.ArrPost[i].TagPost)
	}

}
