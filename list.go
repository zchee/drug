// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pkgutil/osutil"
	"github.com/urfave/cli"
)

var listCommand = cli.Command{
	Name:  "list",
	Usage: "Display the recorded ingestion time results.",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "c, count",
			Usage: "with count mode.",
		},
		cli.BoolFlag{
			Name:  "n, name",
			Usage: "list the drug names mode.",
		},
	},
	Before: initList,
	Action: runList,
}

var (
	listCount bool
	listNames bool
)

func initList(ctx *cli.Context) error {
	listCount = ctx.Bool("count")
	listNames = ctx.Bool("name")
	return nil
}

func runList(ctx *cli.Context) error {
	if listCount && listNames {
		return fmt.Errorf("connot use -count(-c) and -name(-n) at the same time")
	}

	if osutil.IsNotExist(dataFile) {
		return fmt.Errorf("Not exist %s file", dataFile)
	}

	f, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}
	data := []*format{}
	if err := json.Unmarshal(f, &data); err != nil {
		return err
	}

	if listNames {
		var buf bytes.Buffer
		seen := make(map[string]bool)
		for _, d := range data {
			if !seen[d.Name] {
				buf.WriteString(d.Name + "\n")
				seen[d.Name] = true
			}
		}
		fmt.Print(buf.String())
		return nil
	}
	table := tablewriter.NewWriter(os.Stdout)
	switch {
	case listCount:
		table.SetHeader([]string{"NAME", "COUNT"})
		m := make(map[string]int)
		for _, d := range data {
			m[d.Name]++
		}
		for k, v := range m {
			table.Append([]string{k, strconv.Itoa(v)})
		}
	default:
		table.SetHeader([]string{"NAME", "WHEN"})
		for _, d := range data {
			table.Append([]string{d.Name, d.When})
		}
	}
	table.Render()

	return nil
}
