package bug

import (
	"github.com/spf13/cobra"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/internal/cobrax"
)

// Cmd describes a bug command.
var Cmd = cobrax.NewCommand("bug", cobrax.WithRunE(cobra.NoArgs), cobrax.WithArgs(cobra.NoArgs))
