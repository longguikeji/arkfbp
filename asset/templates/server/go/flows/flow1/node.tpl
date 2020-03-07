package {{ .FlowName }}

import (
	"github.com/longguikeji/arkfbp-go/node"
)

// {{ .ClassName }} ...
type {{ .ClassName }} struct {
	node.{{ .BaseClassName }}
}

// ID ...
func (n {{ .ClassName }}) ID() string {
	return "{{ .ID }}"
}
