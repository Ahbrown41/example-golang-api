/*
 * (c)  2022 – Example, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"reflect"
)

// Global, caches struct validation info
var val = validator.New()

type Errors struct {
	Struct string  `json:"struct"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Field string `json:"struct"`
	Error string `json:"error"`
}

func (m *Errors) Error() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Printf("Cannot serialize error object: %s\n", err)
	}
	return string(bytes)
}

// Validate - validates that a struct is correct
func Validate(obj interface{}) error {
	err := val.Struct(obj)
	if err != nil {
		error := Errors{Struct: reflect.TypeOf(obj).String()}
		for _, err := range err.(validator.ValidationErrors) {
			error.Errors = append(error.Errors, Error{Field: err.Field(), Error: err.Error()})
		}
		return &error
	}
	return nil
}
