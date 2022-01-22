/*
 * (c)  2022 – Example, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package entity

import (
	"api-reference-golang/models"
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	gorm.Model
	Name  string       `json:"name"  gorm:"unique_index" validate:"required"`
	Value int32        `json:"value" gorm:"not null" validate:"gte=0,lte=100"`
	Date  time.Time    `json:"date"  gorm:"not null" validate:"required"`
	Items []EntityItem `json:"items"`
}

type EntityItem struct {
	gorm.Model `json:"-"`
	ItemName   string `json:"item-name" gorm:"unique_index" validate:"required"`
	EntityID   int    `json:"-"`
}

func (u *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	return models.Validate(u)
}

func (u *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	return models.Validate(u)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(Entity{})
	db.AutoMigrate(EntityItem{})
}
