package multiflag

import (
	"bytes"
	"flag"
	"testing"
	"time"
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
		{
			"string",
			new(string),
			[]string{"s", "src", "source"},
			"",
			"Source files `directory`.",
			`Usage:
  -s, -source, -src directory
    	Source files directory.
`,
		},
		{
			"int",
			new(int),
			[]string{"l", "limit", "lim"},
			0,
			"Queue workers `limit`.",
			`Usage:
  -l, -lim, -limit limit
    	Queue workers limit.
`,
		},
		{
			"int64",
			new(int),
			[]string{"d", "dep", "depth"},
			64,
			"Recursion max `depth`.",
			`Usage:
  -d, -dep, -depth depth
    	Recursion max depth.
`,
		},
		{
			"uint",
			new(uint),
			[]string{"count-links", "l"},
			uint(0),
			"count sizes many times if hard linked",
			`Usage:
  -count-links, -l uint
    	count sizes many times if hard linked
`,
		},
		{
			"uint64",
			new(uint64),
			[]string{"d", "max-depth"},
			uint64(0),
			"print the total for a directory (or file, with --all)",
			`Usage:
  -d, -max-depth uint
    	print the total for a directory (or file, with --all)
`,
		},
		{
			"float64",
			new(float64),
			[]string{"p", "precision", "prec"},
			.1e6,
			"Output `precision`.",
			`Usage:
  -p, -prec, -precision precision
    	Output precision.
`,
		},
		{
			"duration",
			new(time.Duration),
			[]string{"to", "t", "timeout"},
			time.Second,
			"Max `timeout` value.",
			`Usage:
  -t, -timeout, -to timeout
    	Max timeout value.
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
