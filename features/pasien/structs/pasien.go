package pasien

const MAXUSER int = 100

type Pasien struct {
	ID, Nama, Username, Password string
}

type TabPasien struct {
	ArrPasien [MAXUSER]Pasien
	N         int
}
