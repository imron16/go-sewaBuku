package mysql

import (
	"errors"

	"github.com/imron16/go-sewaBuku/shared/constant"
	"github.com/jinzhu/gorm"

	"github.com/imron16/go-sewaBuku/domain/buku"
	"github.com/imron16/go-sewaBuku/entity"
)

type bukuRepo struct {
	db *gorm.DB
}

func SetupBukuRepo(db *gorm.DB) buku.Repository {
	return &bukuRepo{db: db}
}

func (a *bukuRepo) All() (out []entity.Buku, err error) {
	err = a.db.Model(&entity.Buku{}).Find(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.NoFindAllTermCondition)
	}
	return out, err
}

func (a *bukuRepo) SaveOrUpdateBuku(req entity.Buku) error {
	operation := a.db.Model(entity.Buku{}).Save(&req).Error
	if operation != nil {
		return errors.New(constant.CantSaveUpdateData)
	}
	return operation
}

func (a *bukuRepo) FindBukuById(id int64) (out entity.Buku, err error) {
	err = a.db.Model(&entity.Buku{}).Where("id = ? ", id).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *bukuRepo) FindBukuByJudul(title string) (out entity.Buku, err error) {
	err = a.db.Model(&entity.Buku{}).Where("judul_buku = ? ", title).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *bukuRepo) DeleteBuku(id int64) error {
	err := a.db.Where("id = ?", id).Delete(entity.Buku{}).Error
	if err != nil {
		return errors.New(constant.CantDeleteData)
	}
	return err
}
