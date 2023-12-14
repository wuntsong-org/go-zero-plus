package tsgen

import (
	_ "embed"
	"path"
	"strings"
	"text/template"

	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/api/spec"
	apiutil "github.com/wuntsong-org/go-zero-plus/tools/goctlwt/api/util"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util/pathx"
)

//go:embed components.tpl
var componentsTemplate string

func genComponents(dir string, api *spec.ApiSpec) error {
	types := api.Types
	if len(types) == 0 {
		return nil
	}

	val, err := buildTypes(types)
	if err != nil {
		return err
	}

	outputFile := apiutil.ComponentName(api) + ".ts"
	filename := path.Join(dir, outputFile)
	if err := pathx.RemoveIfExist(filename); err != nil {
		return err
	}

	fp, created, err := apiutil.MaybeCreateFile(dir, ".", outputFile)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	t := template.Must(template.New("componentsTemplate").Parse(componentsTemplate))
	return t.Execute(fp, map[string]string{
		"componentTypes": val,
	})
}

func buildTypes(types []spec.Type) (string, error) {
	var builder strings.Builder
	first := true
	for _, tp := range types {
		if first {
			first = false
		} else {
			builder.WriteString("\n")
		}
		if err := writeType(&builder, tp); err != nil {
			return "", apiutil.WrapErr(err, "Type "+tp.Name()+" generate error")
		}
	}

	return builder.String(), nil
}
