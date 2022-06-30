/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 09:52:35
 */
package dao

import (
	"github.com/kumataahh/eagle-eye/internal/models"

	"gorm.io/gorm"
)

type userDao struct {
	DB *gorm.DB
}

var UserDao = userDao{DB: models.Db}
