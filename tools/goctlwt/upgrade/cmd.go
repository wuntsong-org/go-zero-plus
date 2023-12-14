package upgrade

import "github.com/wuntsong-org/go-zero-plus/tools/goctlwt/internal/cobrax"

// Cmd describes an upgrade command.
var Cmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade))
