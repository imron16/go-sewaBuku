package mysql

import (
	"errors"

	"github.com/imron16/go-sewaBuku/domain/sewa"
	"github.com/imron16/go-sewaBuku/entity"
	"github.com/imron16/go-sewaBuku/shared/constant"
	"github.com/jinzhu/gorm"
)

type sewaRepo struct {
	db *gorm.DB
}

func SetupSewaRepo(db *gorm.DB) sewa.Repository {
	return &sewaRepo{db: db}
}

func (a *sewaRepo) All() (out []entity.Sewa, err error) {
	err = a.db.Model(&entity.Sewa{}).Find(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.NoFindAllTermCondition)
	}
	return out, err
}

func (a *sewaRepo) SaveOrUpdateSewa(req entity.Sewa) error {
	operation := a.db.Model(entity.Sewa{}).Save(&req).Error
	if operation != nil {
		return errors.New(constant.CantSaveUpdateData)
	}
	return operation
}

func (a *sewaRepo) FindSewaById(id int64) (out entity.Sewa, err error) {
	err = a.db.Model(&entity.Sewa{}).Where("id = ? ", id).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *sewaRepo) FindSewaByAnggota(anggota string) (out entity.Sewa, err error) {
	err = a.db.Model(&entity.Sewa{}).Where("anggota = ? ", anggota).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New(constant.CantFindTermCondition)
	}
	return out, err
}

func (a *sewaRepo) DeleteSewa(id int64) error {
	err := a.db.Where("id = ?", id).Delete(entity.Buku{}).Error
	if err != nil {
		return errors.New(constant.CantDeleteData)
	}
	return err
}
