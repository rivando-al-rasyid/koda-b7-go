package main

type dataSekolah struct {
	Nama    string
	Jurusan string
}

type dataDiri struct {
	Nama        string
	Foto        string
	Email       string
	Umur        uint8
	NoTel       string
	SPernikahan bool
	Pendidikan  []dataSekolah
}
