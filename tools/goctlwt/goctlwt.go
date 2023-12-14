package main

import (
	"github.com/wuntsong-org/go-zero-plus/core/load"
	"github.com/wuntsong-org/go-zero-plus/core/logx"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwtwt/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
