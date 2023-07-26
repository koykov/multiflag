package multiflag

import (
	"bytes"
	"flag"
	"testing"
)

func TestMultiFlag(t *testing.T) {
	t.Run("bool var", func(t *testing.T) {
		var b bool
		var buf bytes.Buffer
		ms := NewMultiFlag("", flag.ExitOnError)
		ms.SetOutput(&buf)
		ms.BoolVars(&b, []string{"nc", "no-cache", "nc"}, false, "Disable cache flag.")
		ms.s.Usage()
		println(buf.String())
	})
}
