package sewa

import "github.com/imron16/go-sewaBuku/entity"

// create interface
type Repository interface {
	All() (out []entity.Sewa, err error)
	SaveOrUpdateSewa(req entity.Sewa) error
	FindSewaById(id int64) (out entity.Sewa, err error)
	FindSewaByAnggota(anggota string) (out entity.Sewa, err error)
	DeleteSewa(id int64) error
}
