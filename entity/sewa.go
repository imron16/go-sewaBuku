package entity

// create entity
type Sewa struct {
	Id        int64  `gorm:"primary_key;auto_increment" json:"id"`
	Anggota   string `gorm:"column:anggota;" json:"anggota"`
	JudulBuku string `gorm:"column:judul_buku;" json:"judul_buku"`
}
