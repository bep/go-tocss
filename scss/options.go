// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package sass provides options for SASS transpilers. Note that there are no
// current pure Go SASS implementation, so the only option is CGO and LibSASS.
// But hopefully, fingers crossed, this will happen.
package scss

type (
	OutputStyle int
)

const (
	NestedStyle OutputStyle = iota
	ExpandedStyle
	CompactStyle
	CompressedStyle
)

type Options struct {
	OutputStyle OutputStyle

	// WithImportResolver can be used to supply a custom import resolver, both to redirect
	// to another URL or to return the body.
	ImportResolver func(url string, prev string) (newURL string, body string, resolved bool)

	EnableEmbeddedSourceMap bool // TODO(bep) test this

}
