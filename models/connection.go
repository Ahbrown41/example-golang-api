/*
 * (c)  2022 – Example, Inc.
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and shall at all times remain, the sole and exclusive property ofExample, Inc.
 * All such information, including the intellectual and technical concepts contained therein, are proprietary toExample, Inc. and are protected by copyright law.
 * Dissemination or reproduction of this information is strictly forbidden unless prior written permission is obtained from a duly authorized representative ofExample, Inc.
 */

package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"runtime/debug"
)

type Connection struct {
	dialect gorm.Dialector
	db      *gorm.DB
}

func New(dialect gorm.Dialector) *Connection {
	return &Connection{dialect: dialect}
}

// Connect - Connects to the Database
func (c *Connection) Connect() error {
	client, err := gorm.Open(c.dialect, &gorm.Config{})
	if err != nil {
		log.Panic().Msg("Failed to connect database")
	}
	c.db = client
	return nil
}

// DB - Get active DB connection
func (c *Connection) DB() *gorm.DB {
	if c == nil {
		log.Error().Stack().Msg("Connection Null")
		return nil
	}
	if c.db == nil {
		log.Error().Stack().Msg("Connection DB Null")
		debug.PrintStack()
	}
	return c.db
}

// Disconnect - Disconnects database connection
func (c *Connection) Disconnect() {
	sqlDB, err := c.db.DB()
	if err != nil {
		log.Err(err)
	}
	sqlDB.Close()
}
