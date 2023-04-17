package buku

import "github.com/imron16/go-sewaBuku/entity"

type Service interface {
	InsertBuku(judulBuku, kategoriBuku, kodeBuku string, jumlahBuku int64) error
	UpdateBuku(judulBuku, kategoriBuku, kodeBuku string, id, jumlahBuku int64) error
	GetAllBuku() ([]entity.Buku, error)
	GetBukuById(id int64) (entity.Buku, error)
	DeleteBuku(id int64) error
}
