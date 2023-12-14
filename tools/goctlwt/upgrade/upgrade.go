package upgrade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/rpc/execx"
)

// upgrade gets the latest goctlwt by
// go install github.com/zeromicro/go-zero/tools/goctlwt@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/wuntsong-org/go-zero-plus/tools/goctlwtwt@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/wuntsong-org/go-zero-plus/tools/goctlwtwt@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
