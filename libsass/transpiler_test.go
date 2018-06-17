// Copyright © 2018 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package libsass

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicTranspile(t *testing.T) {
	assert := require.New(t)
	src := bytes.NewBufferString(`div { p { color: red; } }`)
	var dst bytes.Buffer

	transpiler, err := New()
	assert.NoError(err)

	assert.NoError(transpiler.Execute(&dst, src))
	assert.Equal("div p {\n  color: red; }\n", dst.String())
}
