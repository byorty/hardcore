package configs

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/test/models"
)

var (
	DAOs = []types.ModelDAO{ 
		models.UserDaoInst(),
		models.PostDaoInst(),
	}
)
