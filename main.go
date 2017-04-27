// Copyright 2016 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// version will be increased when upgrading release version.
	version = "0.0.1"
)

type format struct {
	Name string `json:"name"`
	When string `json:"when"`
}

func main() {
	app := cli.NewApp()
	app.Name = "drug"
	app.Usage = "Records ingestion time of the drugs."
	app.Version = version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug mode",
		},
	}

	app.Commands = []cli.Command{
		listCommand,
		takeCommand,
	}

	cli.ErrWriter = &fatalWriter{cli.ErrWriter}
	if err := app.Run(os.Args); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
