// cmd_client.go contains the core of the VPN-client

package vpn

import (
	"context"
	"flag"
	"fmt"
	vpn2 "github.com/doocaat/simple-vpn/client"
	"github.com/doocaat/simple-vpn/config"
	"github.com/google/subcommands"
)

// clientCmd is the structure for this sub-command.
//
type clientCmd struct {
	// The configuration file
	config *config.Reader
}

//
// Glue for our sub-command-library.
//
func (*clientCmd) Name() string     { return "client" }
func (*clientCmd) Synopsis() string { return "Start the VPN-client." }
func (*clientCmd) Usage() string {
	return `client :
  Launch the VPN-client.
`
}

//
// Flag setup
//
func (p *clientCmd) SetFlags(f *flag.FlagSet) {
}

//
// Entry-point.
//
func (p *clientCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	//
	// Ensure we have a configuration file.
	//
	if len(f.Args()) < 1 {
		fmt.Printf("We expect a configuration-file to be specified.\n")
		return subcommands.ExitFailure
	}

	//
	// Parse the configuration file.
	//
	var err error
	p.config, err = config.New(f.Args()[0])
	if err != nil {
		fmt.Printf("Failed to read the configuration file %s - %s\n", f.Args()[0], err.Error())
		return subcommands.ExitFailure
	}

	vpn := vpn2.NewVpnClient(
		p.config.Get("vpn"),
		p.config.Get("key"),
		p.config.Get("name"),
		p.config.Get("peers"),
		p.config.Get("up"),
	)
	err = vpn.Start()

	if err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
