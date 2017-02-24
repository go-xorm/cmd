// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

var CmdDriver = &Command{
	UsageLine: "driver",
	Short:     "list all supported drivers",
	Long: `
list all supported drivers
`,
}

func init() {
	CmdDriver.Run = runDriver
	CmdDriver.Flags = map[string]bool{}
}

func runDriver(cmd *Command, args []string) {
	for n, d := range supportedDrivers {
		fmt.Println(n, "\t", d)
	}
}
