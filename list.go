// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pkgutil/osutil"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var listCommand = cli.Command{
	Name:  "list",
	Usage: "Display the recorded results.",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "c, count",
			Usage: "with count mode.",
		},
	},
	Before: initList,
	Action: runList,
}

var listCount bool

func initList(ctx *cli.Context) error {
	listCount = ctx.Bool("count")
	return nil
}

func runList(ctx *cli.Context) error {
	if osutil.IsNotExist(dataFile()) {
		log.Errorf("Not exist %s file", dataFile())
	}

	f, err := ioutil.ReadFile(dataFile())
	if err != nil {
		log.Fatal(err)
	}
	data := []*format{}
	if err := json.Unmarshal(f, &data); err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	if listCount {
		table.SetHeader([]string{"NAME", "COUNT"})
		m := make(map[string]int)
		for _, d := range data {
			m[d.Name]++
		}
		for k, v := range m {
			table.Append([]string{k, strconv.Itoa(v)})
		}
	} else {
		table.SetHeader([]string{"NAME", "WHEN"})
		for _, d := range data {
			table.Append([]string{d.Name, d.When})
		}
	}
	table.Render()

	return nil
}
