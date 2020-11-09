package where

import (
	"github.com/jinzhu/gorm"
	"github.com/winkb/ddup-go-util/orm/where"
)

func ForGorm(wr *where.Where, db *gorm.DB) *gorm.DB {

	switch wr.Name {
	case "like":
		return db.Where("`"+wr.Name+"` like '%?%'", wr.Val)
	}

	return db.Where("`"+wr.Name+"`"+wr.Con+"?", wr.Val)
}
