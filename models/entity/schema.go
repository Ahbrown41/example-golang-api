/*
 * (c)  2022 – Wawa, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package entity

import (
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	gorm.Model
	Name  string    `json:"name" gorm:"size:255;index:unique;check:name == ''"`
	Value int32     `json:"value" gorm:"NOT NULL"`
	Date  time.Time `json:"date" gorm:"NOT NULL"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(Entity{})
}
