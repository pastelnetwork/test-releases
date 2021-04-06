// Code generated by goa v3.3.1, DO NOT EDIT.
//
// walletnode HTTP client CLI support package
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/rest/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	ticketsc "github.com/pastelnetwork/walletnode/rest/gen/http/tickets/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tickets (add|list)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tickets add --body '{
      "address": "u40",
      "pastel_id": "123456789",
      "signature": "wdk231r23kkl23rn23jk423nkl."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		ticketsFlags = flag.NewFlagSet("tickets", flag.ContinueOnError)

		ticketsAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		ticketsAddBodyFlag = ticketsAddFlags.String("body", "REQUIRED", "")

		ticketsListFlags = flag.NewFlagSet("list", flag.ExitOnError)
	)
	ticketsFlags.Usage = ticketsUsage
	ticketsAddFlags.Usage = ticketsAddUsage
	ticketsListFlags.Usage = ticketsListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tickets":
			svcf = ticketsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tickets":
			switch epn {
			case "add":
				epf = ticketsAddFlags

			case "list":
				epf = ticketsListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tickets":
			c := ticketsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = ticketsc.BuildAddPayload(*ticketsAddBodyFlag)
			case "list":
				endpoint = c.List()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// ticketsUsage displays the usage of the tickets command and its subcommands.
func ticketsUsage() {
	fmt.Fprintf(os.Stderr, `The tickets serves tickets data.
Usage:
    %s [globalflags] tickets COMMAND [flags]

COMMAND:
    add: Add a new ticket and return its ID.
    list: List all stored tickets.

Additional help:
    %s tickets COMMAND --help
`, os.Args[0], os.Args[0])
}
func ticketsAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tickets add -body JSON

Add a new ticket and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` tickets add --body '{
      "address": "u40",
      "pastel_id": "123456789",
      "signature": "wdk231r23kkl23rn23jk423nkl."
   }'
`, os.Args[0])
}

func ticketsListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tickets list

List all stored tickets.

Example:
    `+os.Args[0]+` tickets list
`, os.Args[0])
}
