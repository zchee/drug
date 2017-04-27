// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
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
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
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

const (
	exactArgs = iota
	minArgs
	maxArgs
)

func checkArgs(context *cli.Context, expected, checkType int, args ...string) (err error) {
	cmdName := context.Command.Name
	switch checkType {
	case exactArgs:
		if context.NArg() != expected {
			err = fmt.Errorf("%s: %q requires exactly <%s> %d argument(s)", os.Args[0], cmdName, strings.Join(args, " "), expected)
		}
	case minArgs:
		if context.NArg() < expected {
			err = fmt.Errorf("%s: %q requires a minimum of <%s> %d argument(s)", os.Args[0], cmdName, strings.Join(args, " "), expected)
		}
	case maxArgs:
		if context.NArg() > expected {
			err = fmt.Errorf("%s: %q requires a maximum of <%s> %d argument(s)", os.Args[0], cmdName, strings.Join(args, " "), expected)
		}
	}

	if err != nil {
		fmt.Printf("Incorrect Usage.\n\n")
		return err
	}
	return nil
}

type fatalWriter struct {
	cliErrWriter io.Writer
}

// Write implements io.Writer interface.
func (f *fatalWriter) Write(b []byte) (n int, err error) {
	logrus.Error(string(b))
	return f.cliErrWriter.Write(b)
}
