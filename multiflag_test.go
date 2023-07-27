package multiflag

import (
	"bytes"
	"flag"
	"testing"
)

func TestMultiFlag(t *testing.T) {
	type stage struct {
		key    string
		ptr    any
		names  []string
		value  any
		usage  string
		expect string
	}
	var stages = []stage{
		{
			"bool",
			new(bool),
			[]string{"nc", "no-cache", "nc"},
			false,
			"Disable cache flag.",
			`Usage:
  -nc, -no-cache
    	Disable cache flag.
`,
		},
	}
	var buf bytes.Buffer
	for _, stg := range stages {
		t.Run(stg.key, func(t *testing.T) {
			buf.Reset()
			ms := NewMultiFlag("", flag.ExitOnError)
			ms.SetOutput(&buf)
			ms.var_(stg.ptr, stg.names, stg.value, stg.usage)
			ms.Usage()
			if buf.String() != stg.expect {
				t.FailNow()
			}
		})
	}
}
