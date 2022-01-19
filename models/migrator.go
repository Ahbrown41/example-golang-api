/*
 * (c)  2022 – Wawa, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package models

import "reference-golang-svc/models/entity"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(entity.Entity{})
}
