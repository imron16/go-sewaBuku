package buku

import (
	"errors"

	"github.com/imron16/go-sewaBuku/entity"
	"github.com/imron16/go-sewaBuku/shared/constant"
)

func (s *service) InsertBuku(judulBuku, kategoriBuku, kodeBuku string, jumlahBuku int64) error {
	result, err := s.bukuRepo.FindBukuByJudul(judulBuku)
	if err != nil {
		return errors.New(constant.InternalServerError)
	}

	if len(result.JudulBuku) > 0 {
		return errors.New(constant.BookTitleAlreadyExist)
	}

	bukus := entity.Buku{}
	bukus.JudulBuku = judulBuku
	bukus.KategoriBuku = kategoriBuku
	bukus.KodeBuku = kodeBuku
	bukus.JumlahBuku = jumlahBuku

	err = s.bukuRepo.SaveOrUpdateBuku(bukus)
	if err != nil {
		return err
	}

	return nil

}
