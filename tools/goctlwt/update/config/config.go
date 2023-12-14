package config

import "github.com/wuntsong-org/go-zero-plus/core/logx"

// Config defines a service configure for goctlwt update
type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
