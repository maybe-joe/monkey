package repl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Repl(t *testing.T) {
	var (
		text = "let add = fn(x, y) { x + y; };"
		in   = strings.NewReader(text)
		out  strings.Builder
	)

	err := Run(in, &out)
	require.NoError(t, err)

	exptected := `LET
IDENT add
=
FUNCTION
(
IDENT x
,
IDENT y
)
{
IDENT x
+
IDENT y
;
}
;
EOF
`

	require.Equal(t, exptected, out.String())
}
