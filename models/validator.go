/*
 * (c)  2022 – Wawa, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func validate(scope *gorm.DB.) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if result, ok := scope.DB().Get(skipValidations); !(ok && result.(bool)) {
			if !scope.HasError() {
				scope.CallMethod("Validate")
				if scope.Value != nil {
					resource := scope.IndirectValue().Interface()
					_, validatorErrors := govalidator.ValidateStruct(resource)
					if validatorErrors != nil {
						if errors, ok := validatorErrors.(govalidator.Errors); ok {
							for _, err := range flatValidatorErrors(errors) {
								scope.DB().AddError(formattedError(err, resource))
							}
						} else {
							scope.DB().AddError(validatorErrors)
						}
					}
				}
			}
		}
	}
}