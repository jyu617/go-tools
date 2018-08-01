package main // import "github.com/jyu617/go-tools/cmd/errcheck-ng"

import (
	"os"

	"github.com/jyu617/go-tools/errcheck"
	"github.com/jyu617/go-tools/lint/lintutil"
)

func main() {
	c := lintutil.CheckerConfig{
		Checker:     errcheck.NewChecker(),
		ExitNonZero: true,
	}
	lintutil.ProcessArgs("errcheck-ng", []lintutil.CheckerConfig{c}, os.Args[1:])
}
