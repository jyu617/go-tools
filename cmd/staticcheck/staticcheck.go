// staticcheck detects a myriad of bugs and inefficiencies in your
// code.
package main // import "github.com/jyu617/go-tools/cmd/staticcheck"

import (
	"os"

	"github.com/jyu617/go-tools/lint/lintutil"
	"github.com/jyu617/go-tools/staticcheck"
)

func main() {
	fs := lintutil.FlagSet("staticcheck")
	gen := fs.Bool("generated", false, "Check generated code")
	fs.Parse(os.Args[1:])
	c := staticcheck.NewChecker()
	c.CheckGenerated = *gen
	cfg := lintutil.CheckerConfig{
		Checker:     c,
		ExitNonZero: true,
	}
	lintutil.ProcessFlagSet([]lintutil.CheckerConfig{cfg}, fs)
}
