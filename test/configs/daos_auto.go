package configs

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	DAOs = []types.ModelDAO{
		models.UserDaoInst(),
		models.PostDaoInst(),
	}
)
