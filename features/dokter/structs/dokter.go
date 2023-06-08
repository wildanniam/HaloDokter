package dokter

const MAXDOKTER int = 100

type Dokter struct {
	ID, Nama, Spesialis, Username, Password string
}

type TabDokter struct {
	ArrDokter [MAXDOKTER]Dokter
	N         int
}
