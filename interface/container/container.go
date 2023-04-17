package container

import (
	"github.com/imron16/go-sewaBuku/infrastructure/mysql"
	"github.com/imron16/go-sewaBuku/shared/database"
	buku "github.com/imron16/go-sewaBuku/usecase/buku/response"
	sewa "github.com/imron16/go-sewaBuku/usecase/sewa/response"
	"github.com/jinzhu/gorm"
)

type Container struct {
	BukuService buku.Service
	SewaService sewa.Service
}

func SetupContainer(tx *gorm.DB) Container {

	database.Migrate(tx)
	bukuService := buku.NewBukuService(mysql.SetupBukuRepo(tx))
	sewaService := sewa.NewSewaService(mysql.SetupSewaRepo(tx))
	return Container{BukuService: bukuService,
		SewaService: sewaService}
}
