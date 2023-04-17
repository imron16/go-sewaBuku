package sewa

import "github.com/imron16/go-sewaBuku/entity"

type Service interface {
	InsertSewa(anggota, judulBuku string) error
	UpdateSewa(anggota, judulBuku string, id int64) error
	GetAllSewa() ([]entity.Sewa, error)
	GetSewaById(id int64) (entity.Sewa, error)
	DeleteSewa(id int64) error
}
