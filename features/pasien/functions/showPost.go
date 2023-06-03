package pasien

import (
	"fmt"
	pasienStruct "projek/features/pasien/structs"
)

func showPost(tPost *pasienStruct.TabPost) {
	//Menampilkan isi postingan yang telah diinputkan pengguna
	for i := 0; i < tPost.N; i++ {
		fmt.Printf("Post %d: %s \nTag : %s \n", i+1, tPost.ArrPost[i].TxtAddPost, tPost.ArrPost[i].TagPost)
	}

}
