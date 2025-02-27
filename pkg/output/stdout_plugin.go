package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

const jsonFormatDesc = `
JSON Format:
~~~~
{
	title: "Article title",
	description: "Article description",
	content: "Article HTML content",
	link: "Article URL",
	updated: "Article update date (String format)",
	updatedParsed: "Article update date (Date format)",
	published: "Article publication date (String format)",
	publishedParsed: "Article publication date (Date format)",
	guid: "Article feed GUID",
	meta: {
		"key": "Metadata keys and values"
	},
	tags: ["list", "of", "tags"]
}
~~~~
`

var stdoutSpec = model.Spec{
	Name:      "stdout",
	Desc:      "New articles are sent as JSON document to the standard output of the process.\n\n" + jsonFormatDesc,
	PropsSpec: []model.PropSpec{},
}

// StdoutOutputPlugin is the STDOUT output plugin
type StdoutOutputPlugin struct{}

// Spec returns plugin spec
func (p *StdoutOutputPlugin) Spec() model.Spec {
	return stdoutSpec
}

// Build creates output provider instance
func (p *StdoutOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	return &StdOutputProvider{
		id:      output.ID,
		alias:   output.Alias,
		spec:    stdoutSpec,
		tags:    output.Tags,
		enabled: output.Enabled,
	}, nil
}

var stdoutOutputPlugin = &StdoutOutputPlugin{}

// StdOutputProvider STDOUT output provider
type StdOutputProvider struct {
	id        int
	alias     string
	spec      model.Spec
	tags      []string
	nbSuccess uint64
	enabled   bool
}

// Send article to STDOUT.
func (op *StdOutputProvider) Send(article *model.Article) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(article)
	fmt.Println(b.String())
	atomic.AddUint64(&op.nbSuccess, 1)
	return nil
}

// GetDef return output provider definition
func (op *StdOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		ID:      op.id,
		Alias:   op.alias,
		Spec:    op.spec,
		Tags:    op.tags,
		Enabled: op.enabled,
	}
	result.Props = map[string]interface{}{
		"nbSuccess": op.nbSuccess,
	}
	return result
}
