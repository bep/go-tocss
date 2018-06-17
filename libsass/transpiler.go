// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package libsass provides a SASS ans SCSS transpiler to CSS
// using github.com/wellington/go-libsass/libs.
package libsass

import (
	"io"
	"io/ioutil"

	"github.com/wellington/go-libsass/libs"
)

type Transpiler struct {
}

func New() (*Transpiler, error) {
	return &Transpiler{}, nil
}

func (t *Transpiler) Execute(dst io.Writer, src io.Reader) error {

	// TODO(bep) SASS vs SCSS syntax
	srcb, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	dataCtx := libs.SassMakeDataContext(string(srcb))
	opts := libs.SassDataContextGetOptions(dataCtx)

	// TODO(bep)
	//libs.SassOptionSetSourceComments(opts, true)
	libs.SassDataContextSetOptions(dataCtx, opts)

	ctx := libs.SassDataContextGetContext(dataCtx)
	compiler := libs.SassMakeDataCompiler(dataCtx)

	libs.SassCompilerParse(compiler)
	libs.SassCompilerExecute(compiler)

	//TODO(bep) libs.SassOptionSetSourceMapEmbed(goopts, true)
	defer libs.SassDeleteCompiler(compiler)

	result := libs.SassContextGetOutputString(ctx)

	io.WriteString(dst, result)

	// Error handling.
	//libs.SassContextGetErrorStatus(goctx)
	//libs.SassContextGetErrorJSON(goctx)

	return nil
}
