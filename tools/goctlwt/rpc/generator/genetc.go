package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	conf "github.com/wuntsong-org/go-zero-plus/tools/goctlwt/config"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/rpc/parser"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util/format"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util/pathx"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util/stringx"
)

//go:embed etc.tpl
var etcTemplate string

// GenEtc generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEtc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetEtc()
	etcFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, fmt.Sprintf("%v.yaml", etcFilename))

	text, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	return util.With("etc").Parse(text).SaveTo(map[string]any{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}
