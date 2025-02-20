// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows
// +build windows

package main

import (
	"os"
)

var ignoreSignals = []os.Signal{os.Interrupt}
var forwardSignals []os.Signal
