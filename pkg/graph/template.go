//go:build go1.16
// +build go1.16

package graph

import _ "embed"

//go:embed template.html
var GCVIS_TMPL string
