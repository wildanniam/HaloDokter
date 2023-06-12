package pasien

const MAXUSER int = 100

type Pasien struct {
	Nama, Username, Password string
	ID                       int
}

type TabPasien struct {
	ArrPasien [MAXUSER]Pasien
	N         int
}
