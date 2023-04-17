package buku

import "github.com/imron16/go-sewaBuku/entity"

// create interface
type Repository interface {
	All() (out []entity.Buku, err error)
	SaveOrUpdateBuku(req entity.Buku) error
	FindBukuById(id int64) (out entity.Buku, err error)
	FindBukuByJudul(title string) (out entity.Buku, err error)
	DeleteBuku(id int64) error
}
