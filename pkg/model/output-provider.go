package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// OutputProps contains properties of an output
type OutputProps map[string]interface{}

// Get property string value
func (p OutputProps) Get(key string) string {
	if val, ok := p[key]; ok {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

// OutputProvider is the output provider interface
type OutputProvider interface {
	Send(article *Article) error
	GetDef() OutputDef
}

// OutputDefCollection is an array of output definition
type OutputDefCollection []*OutputDef

// OutputDef contains output definition
type OutputDef struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
	Spec
	Tags    []string    `json:"tags,omitempty"`
	Props   OutputProps `json:"props:omitempty"`
	Enabled bool        `json:"enabled"`
}

// Hash computes spec hash
func (spec OutputDef) Hash() string {
	key := fmt.Sprintf("%s-%d", spec.Name, spec.ID)
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
