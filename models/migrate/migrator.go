/*
 * (c)  2022 – Example, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property ofExample, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary toExample, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative ofExample, Inc.
 */

package migrate

import (
	"api-reference-golang/models/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	entity.Migrate(db)
}
