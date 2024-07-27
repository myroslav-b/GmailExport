package main

import (
	"context"
	"gmailexport/app/getclient"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var user = "me"

// tFilter represents the filter options for searching Gmail messages
type tFilter struct {
	MessageId string `short:"m" long:"message" description:"message id"`
	Label     string `short:"l" long:"label" description:"label"`
	From      string `short:"f" long:"from" description:"sender's email address"`
	To        string `short:"t" long:"to" description:"recipient's email address"`
	Subject   string `short:"s" long:"subject" description:"email subject"`
}

// query constructs a Gmail search query string from the filter options
func (filter tFilter) query() string {
	ss := []string{filter.messageId(), filter.label(), filter.from(), filter.to(), filter.subject()}
	q := ""
	for _, s := range ss {
		if s != "" {
			if q == "" {
				q = s
			} else {
				q = q + " AND " + s
			}
		}
	}
	return q
}

// Helper methods to construct individual query parts
func (filter tFilter) messageId() string {
	s := ""
	if filter.MessageId != "" {
		s = "rfc822msgid:" + filter.MessageId
	}
	return s
}

func (filter tFilter) label() string {
	s := ""
	if filter.Label != "" {
		s = "label:" + filter.Label
	}
	return s
}

func (filter tFilter) from() string {
	s := ""
	if filter.From != "" {
		s = "from:" + filter.From
	}
	return s
}

func (filter tFilter) to() string {
	s := ""
	if filter.To != "" {
		s = "to:" + filter.To
	}
	return s
}

func (filter tFilter) subject() string {
	s := ""
	if filter.Subject != "" {
		s = "subject:" + filter.Subject
	}
	return s
}

// tStatement represents the output options for the exported messages
type tStatement struct {
	Output string `short:"O" long:"output" default:"stdout" optional:"non-empty" optional-value:"gmail" description:"output path: stdout - if missing, else output to file; value_of_param - template for the name (the equal sign (=) is required), or gmail - if option occurs without an argument"`
	Split  bool   `short:"S" long:"split" description:"split output into multiple files"`
	Format string `short:"F" long:"format" choice:"json" choice:"txt" default:"json" description:"output format"`
	Area   string `short:"A" long:"area" choice:"raw" choice:"all" choice:"small" choice:"easy" default:"all" description:"fullness of the output"`
}

// tOpts combines the filter and statement options
type tOpts struct {
	Statement tStatement `group:"Presentation of results"`
	Filter    tFilter    `group:"Selection conditions"`
}

func (opts tOpts) filter() tFilter {
	return opts.Filter
}

func main() {
	var opts tOpts
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getclient.GetClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	err = export(srv, user, opts)
	if err != nil {
		log.Fatalf("Func export: %v", err)
	}

}
