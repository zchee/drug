// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"

	"github.com/pkgutil/osutil"
	"github.com/urfave/cli"
)

var intervalCommand = cli.Command{
	Name:      "interval",
	Usage:     "Calculate the ingestion interval of the drug.",
	ArgsUsage: "<drug name>",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "n, now",
			Usage: "until now",
		},
	},
	Before: initInterval,
	Action: runInterval,
}

var (
	intervalDrugName string
	intervalNow      bool
)

func initInterval(ctx *cli.Context) error {
	intervalDrugName = ctx.Args().First()
	intervalNow = ctx.Bool("now")
	return nil
}

func runInterval(ctx *cli.Context) error {
	if err := checkArgs(ctx, 1, minArgs, "drug name"); err != nil {
		return err
	}

	if osutil.IsNotExist(dataFile) {
		return fmt.Errorf("Not exist %s file", dataFile)
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}
	list := []*format{}
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	intervals := []string{}
	for _, d := range list {
		if d.Name == intervalDrugName {
			intervals = append(intervals, d.When)
		}
	}
	if len(intervals) <= 1 {
		return fmt.Errorf("%s was ingested only once", intervalDrugName)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(intervals)))
	last, err := time.Parse(time.Stamp, intervals[0])
	if err != nil {
		return err
	}
	var interval time.Duration
	if intervalNow {
		now, err := time.Parse(time.Stamp, time.Now().Format(time.Stamp))
		if err != nil {
			return err
		}
		interval = now.Sub(last)
	} else {
		prev, err := time.Parse(time.Stamp, intervals[1])
		if err != nil {
			return err
		}
		interval = last.Sub(prev)
	}
	fmt.Print(interval.String())

	return nil
}
