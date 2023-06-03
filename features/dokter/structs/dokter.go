package dokter

const MAXDOKTER int = 100

type Dokter struct {
	Nama, Spesialis, Username, Password string
}

type TabDokter struct {
	ArrDokter [MAXDOKTER]Dokter
	N         int
}
