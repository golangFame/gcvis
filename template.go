//go:build go1.16
// +build go1.16

package main

import _ "embed"

//go:embed template.html
var GCVIS_TMPL string
