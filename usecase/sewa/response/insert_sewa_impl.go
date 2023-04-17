package sewa

import (
	"errors"

	"github.com/imron16/go-sewaBuku/entity"
	"github.com/imron16/go-sewaBuku/shared/constant"
)

func (s *service) InsertSewa(anggota, judulBuku string) error {
	result, err := s.sewaRepo.FindSewaByAnggota(anggota)

	if err != nil {
		return errors.New(constant.InternalServerError)
	}

	if len(result.Anggota) > 0 {
		return errors.New(constant.BookTitleAlreadyExist)
	}

	sewas := entity.Sewa{}
	sewas.Anggota = anggota
	sewas.JudulBuku = judulBuku

	err = s.sewaRepo.SaveOrUpdateSewa(sewas)

	if err != nil {
		return err
	}

	return nil

}
