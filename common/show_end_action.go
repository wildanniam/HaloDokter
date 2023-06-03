package common

import "fmt"

// Berfungsi
func ShowEndAction() {
	var tempInput string

	fmt.Println()
	fmt.Println("Tekan enter untuk selesai...")
	fmt.Scanln(&tempInput)
}
