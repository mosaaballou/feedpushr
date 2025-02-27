package output

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/pkg/model"
)

var spec = model.Spec{
	Name: "readflow",
	Desc: "Send new articles a readflow instance.",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.Text,
		},
		{
			Name: "apiKey",
			Desc: "API KEY",
			Type: model.Password,
		},
	},
}

// ReadflowOutputPlugin is the Readflow output plugin
type ReadflowOutputPlugin struct{}

// Spec returns plugin spec
func (p *ReadflowOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Readflow output provider instance
func (p *ReadflowOutputPlugin) Build(output *model.OutputDef) (model.OutputProvider, error) {
	u := output.Props.Get("url")
	if u == "" {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	apiKey := output.Props.Get("apiKey")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API key property")
	}
	return &ReadflowOutputProvider{
		id:        output.ID,
		alias:     output.Alias,
		spec:      spec,
		tags:      output.Tags,
		targetURL: _url.String(),
		apiKey:    apiKey,
		enabled:   output.Enabled,
	}, nil
}

var readflowOutputPlugin = &ReadflowOutputPlugin{}

// ReadflowOutputProvider output provider to send articles to Readflow
type ReadflowOutputProvider struct {
	id        int
	alias     string
	spec      model.Spec
	tags      []string
	enabled   bool
	nbError   uint64
	nbSuccess uint64
	targetURL string
	apiKey    string
}

// Send article to a Readflow instance.
func (op *ReadflowOutputProvider) Send(article *model.Article) error {
	nb, err := sendToReadflow(op.targetURL, op.apiKey, article)
	if err != nil {
		atomic.AddUint64(&op.nbError, 1)
		return err
	}
	atomic.AddUint64(&op.nbSuccess, uint64(nb))
	return nil
}

// GetDef return output definition
func (op *ReadflowOutputProvider) GetDef() model.OutputDef {
	result := model.OutputDef{
		ID:      op.id,
		Alias:   op.alias,
		Spec:    op.spec,
		Tags:    op.tags,
		Enabled: op.enabled,
	}
	result.Props = map[string]interface{}{
		"url":       op.targetURL,
		"apiKey":   op.apiKey,
		"nbError":   op.nbError,
		"nbSuccess": op.nbSuccess,
	}
	return result
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OUTPUT_PLUGIN,
	}
}

// GetOutputPlugin return output plugin instance
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &ReadflowOutputPlugin{}, nil
}
