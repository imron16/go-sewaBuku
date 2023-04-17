package entity

// create entity
type Buku struct {
	Id           int64  `gorm:"primary_key;auto_increment" json:"id"`
	JudulBuku    string `gorm:"column:judul_buku;" json:"judul_buku"`
	KategoriBuku string `gorm:"column:kategori_buku;" json:"kategori_buku"`
	KodeBuku     string `gorm:"column:kode_buku;" json:"kode_buku"`
	JumlahBuku   int64  `gorm:"column:jumlah_buku;" json:"jumlah_buku"`
}
