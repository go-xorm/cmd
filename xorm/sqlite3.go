// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sqlite3

package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	supportedDrivers["sqlite3"] = "github.com/mattn/go-sqlite3"
}
