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

	"github.com/bep/go-tocss/api"
	"github.com/bep/go-tocss/scss"
	"github.com/wellington/go-libsass/libs"
)

var _ api.Transpiler = (*libsassTranspiler)(nil)

type libsassTranspiler struct {
	options scss.Options
}

func New(options scss.Options) (*libsassTranspiler, error) {
	return &libsassTranspiler{options: options}, nil
}

func (t *libsassTranspiler) Execute(dst io.Writer, src io.Reader) error {
	// TODO(bep) basepath

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	sourceStr := string(b)

	dataCtx := libs.SassMakeDataContext(sourceStr)
	opts := libs.SassDataContextGetOptions(dataCtx)

	if t.options.ImportResolver != nil {
		idx := libs.BindImporter(opts, t.options.ImportResolver)
		defer libs.RemoveImporter(idx)
	}

	libs.SassOptionSetSourceMapEmbed(opts, t.options.EnableEmbeddedSourceMap)
	//libs.SassOptionSetIncludePath(opts, incs)
	//libs.SassOptionSetPrecision(opts, TODO)
	libs.SassOptionSetOutputStyle(opts, int(t.options.OutputStyle))
	libs.SassOptionSetSourceComments(opts, false)
	libs.SassDataContextSetOptions(dataCtx, opts)

	ctx := libs.SassDataContextGetContext(dataCtx)
	compiler := libs.SassMakeDataCompiler(dataCtx)

	libs.SassCompilerParse(compiler)
	libs.SassCompilerExecute(compiler)

	libs.SassOptionSetSourceMapEmbed(opts, t.options.EnableEmbeddedSourceMap)
	defer libs.SassDeleteCompiler(compiler)

	result := libs.SassContextGetOutputString(ctx)

	io.WriteString(dst, result)

	// Error handling.
	//libs.SassContextGetErrorStatus(goctx)
	//libs.SassContextGetErrorJSON(goctx)

	return nil
}