package main // import "github.com/jyu617/go-tools/cmd/stylecheck"
import (
	"os"

	"github.com/jyu617/go-tools/lint/lintutil"
	"github.com/jyu617/go-tools/stylecheck"
)

func main() {
	fs := lintutil.FlagSet("stylecheck")
	gen := fs.Bool("generated", false, "Check generated code")
	fs.Parse(os.Args[1:])
	c := stylecheck.NewChecker()
	c.CheckGenerated = *gen
	cfg := lintutil.CheckerConfig{
		Checker:     c,
		ExitNonZero: true,
	}
	lintutil.ProcessFlagSet([]lintutil.CheckerConfig{cfg}, fs)
}
