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
	"gorm.io/gorm"
	"log"
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
		panic("Failed to connect database")
	}
	c.db = client
	return nil
}

// DB - Get active DB connection
func (c *Connection) DB() *gorm.DB {
	return c.db
}

// Disconnect - Disconnects database connection
func (c *Connection) Disconnect() {
	sqlDB, err := c.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
