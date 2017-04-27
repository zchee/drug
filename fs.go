// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"

	xdgbasedir "github.com/zchee/go-xdgbasedir"
)

var (
	dataFile   = filepath.Join(dataDir(), "drug.json")
	configFile = filepath.Join(dataDir(), "config.json")
)

func dataDir() string {
	return filepath.Join(xdgbasedir.DataHome(), "drug")
}
