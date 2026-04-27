package models

type DataSekolah struct {
	Nama    string
	Jurusan string
}

type DataDiri struct {
	Nama        string
	Foto        string
	Email       string
	Umur        uint8
	NoTel       string
	SPernikahan bool
	Pendidikan  []DataSekolah
}
