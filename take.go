// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/pkgutil/osutil"
	"github.com/urfave/cli"
)

var takeCommand = cli.Command{
	Name:      "take",
	Usage:     "TODO",
	Flags:     []cli.Flag{},
	ArgsUsage: "<drug name>",
	Before:    initTake,
	Action:    runTake,
}

var (
	runDrugName string
)

func initTake(ctx *cli.Context) error {
	runDrugName = ctx.Args().First()
	return nil
}

func runTake(ctx *cli.Context) error {
	if err := checkArgs(ctx, 1, minArgs, "drug name"); err != nil {
		return err
	}

	var data []*format
	if osutil.IsExist(dataFile) {
		oldData, err := ioutil.ReadFile(dataFile)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(oldData, &data); err != nil {
			return err
		}
	} else {
		if err := osutil.MkdirAll(dataDir(), 0700); err != nil {
			return err
		}
	}

	d := &format{
		Name: runDrugName,
		When: time.Now().Format(time.Stamp),
	}
	data = append(data, d)
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dataFile, out, 0644); err != nil {
		return err
	}

	return nil
}
