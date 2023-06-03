package pasien

import (
	pasienFunc "projek/features/pasien/functions"
	pasienStruct "projek/features/pasien/structs"
)

func Main(pasien *pasienStruct.TabPasien) {
	pasienFunc.LoginPasien(pasien)
}
