// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package api provides the common API for the tranpilation of the source
// to CSS.
package api

import (
	"io"
)

type Transpiler interface {
	Execute(dst io.Writer, src io.Reader) error
}
