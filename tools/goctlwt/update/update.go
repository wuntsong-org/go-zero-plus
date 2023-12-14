package main

import (
	"flag"
	"net/http"
	"os"
	"path"

	"github.com/wuntsong-org/go-zero-plus/core/conf"
	"github.com/wuntsong-org/go-zero-plus/core/hash"
	"github.com/wuntsong-org/go-zero-plus/core/logx"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/update/config"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/util/pathx"
)

const (
	contentMd5Header = "Content-Md5"
	filename         = "goctl"
)

var configFile = flag.String("f", "etc/update-api.json", "the config file")

func forChksumHandler(file string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !pathx.FileExists(file) {
			logx.Errorf("file %q not exist", file)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			logx.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		chksum := hash.Md5Hex(content)
		if chksum == r.Header.Get(contentMd5Header) {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set(contentMd5Header, chksum)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	fs := http.FileServer(http.Dir(c.FileDir))
	http.Handle(c.FilePath, http.StripPrefix(c.FilePath, forChksumHandler(path.Join(c.FileDir, filename), fs)))
	logx.Must(http.ListenAndServe(c.ListenOn, nil))
}
