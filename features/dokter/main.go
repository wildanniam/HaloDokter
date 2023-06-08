package dokter

import (
	"fmt"
	"projek/common"

	dokterFunc "projek/features/dokter/functions"
	dokterMenu "projek/features/dokter/menu"
	dokterStruct "projek/features/dokter/structs"
	postStruct "projek/features/post/structs"
)

func Main(arrDokter *dokterStruct.TabDokter, arrPost *postStruct.TabPost) {
	var input int
	dokterMenu.ShowAuthDokterMenu()

	fmt.Scan(&input)

	for input != 3 {
		if input == 1 {
			common.ResetConsole()
			dokterFunc.RegisDokter(arrDokter)
			dokterMenu.ShowAuthDokterMenu()
		} else if input == 2 {
			/*dokterFunc.LoginDokter(arrDokter) {
			common.ResetConsole()
			dokterMenu.ShowHomeDokterMenu()
			fmt.Scan(&input)
			for input != 1 {
				dokterFunc.ShowPost(arrPost)
				dokterFunc.RepChat(arrPost) // reply chat masih mau diubah lokasinya kayaknya. rencana mau dibuat di luar aja jadi "projek/features/post" tapi masih butuh briefing dulu setuju atau gak.
				}
			}*/
		} else {
			fmt.Println("Input tidak valid !")
		}
		fmt.Scan(&input)
	}
	Main(arrDokter, arrPost)
}
