//
// Show our version.
//

package vpn

import (
	"context"
	"flag"
	"fmt"
	"runtime"

	"github.com/google/subcommands"
)

//
// This is set via the travis-builder (or will be when I set that up).
//
var (
	version = "unreleased"
)

type versionCmd struct {
	verbose bool
}

//
// Glue
//
func (*versionCmd) Name() string     { return "version" }
func (*versionCmd) Synopsis() string { return "Show our version." }
func (*versionCmd) Usage() string {
	return `version :
  Report upon our version, and exit.
`
}

//
// Flag setup
//
func (p *versionCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.verbose, "verbose", false, "Show go version the binary was generated with.")
}

//
// Show our version.
//
func showVersion(verbose bool) {
	fmt.Printf("%s\n", version)
	if verbose {
		fmt.Printf("Built with %s\n", runtime.Version())
	}
}

//
// Entry-point.
//
func (p *versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	showVersion(p.verbose)
	return subcommands.ExitSuccess
}
