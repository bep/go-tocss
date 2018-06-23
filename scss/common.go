// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package scss provides options for SCSS transpilers. Note that there are no
// current pure Go SASS implementation, so the only option is CGO and LibSASS.
// But hopefully, fingers crossed, this will happen.
package scss

import (
	"encoding/json"
	"fmt"
)

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
	//TODO(bep) icon font path
	// Default is nested.
	OutputStyle OutputStyle

	// Precision of floating point math.
	Precision int

	// File paths to use to resolve imports.
	IncludePaths []string

	// ImportResolver can be used to supply a custom import resolver, both to redirect
	// to another URL or to return the body.
	ImportResolver func(url string, prev string) (newURL string, body string, resolved bool)

	EnableEmbeddedSourceMap bool // TODO(bep) test this

}

func JSONToError(jsonstr string) (e Error) {
	if err := json.Unmarshal([]byte(jsonstr), &e); err != nil {
		e.Message = "unknown error"
	}
	return
}

type Error struct {
	Status  int    `json:"status"`
	Column  int    `json:"column"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("file %q, line %d, col %d: %s ", e.File, e.Line, e.Column, e.Message)
}
