package migrate

import (
	"net/url"
	"os"
	"strings"

	"github.com/wuntsong-org/go-zero-plus/core/stringx"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/rpc/execx"
)

var (
	defaultProxy   = "https://goproxy.cn"
	defaultProxies = []string{defaultProxy}
)

func goProxy() []string {
	wd, err := os.Getwd()
	if err != nil {
		return defaultProxies
	}

	proxy, err := execx.Run("go env GOPROXY", wd)
	if err != nil {
		return defaultProxies
	}
	list := strings.FieldsFunc(proxy, func(r rune) bool {
		return r == '|' || r == ','
	})
	var ret []string
	for _, item := range list {
		if len(item) == 0 {
			continue
		}
		_, err = url.Parse(item)
		if err == nil && !stringx.Contains(ret, item) {
			ret = append(ret, item)
		}
	}
	if !stringx.Contains(ret, defaultProxy) {
		ret = append(ret, defaultProxy)
	}
	return ret
}
