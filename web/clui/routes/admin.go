/*
Copyright <holder> All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package routes

import (
	"log"
	"time"

	"github.com/IBM/cloudland/web/clui/model"
	"github.com/IBM/cloudland/web/sca/dbs"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func adminPassword() (password string) {
	time.Sleep(time.Second * 5)
	passwd := viper.GetString("admin.password")
	if password == "" {
		password = "passw0rd"
	}
	return
}

func adminInit() {
	username := "admin"
	dbs.AutoUpgrade("01-admin-upgrade", func(db *gorm.DB) (err error) {
		if err = db.Take(&model.User{}, "username = ?", username).Error; err != nil {
			dbFunc := DB
			defer func() { DB = dbFunc }()
			DB = func() *gorm.DB { return db }
			password := adminPassword()
			_, err = userAdmin.Create(username, password)
			if err != nil {
				return
			}
			_, err = orgAdmin.Create(username, username)
			if err != nil {
				return
			}
		}
		return
	})
}
