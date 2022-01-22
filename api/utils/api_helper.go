/*
 * (c)  2022 – Example, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property of Wawa, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary to Wawa, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative of Wawa, Inc.
 */

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func ParamInt64(c *gin.Context, param string) int64 {
	int, err := strconv.ParseInt(c.Param(param), 10, 64)
	if err != nil {
		log.Warn().Msgf("Could not decode param: %s, error: %s", param, err)
		return 0
	} else {
		return int
	}
}

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	var res ResponseData
	res.Status = status
	res.Data = payload
	c.JSON(200, res)
}
