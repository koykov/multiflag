package multiflag

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

type MultiFlag struct {
	s *flag.FlagSet
}

func NewMultiFlag(name string, errorHandling flag.ErrorHandling) *MultiFlag {
	f := MultiFlag{s: flag.NewFlagSet(name, errorHandling)}
	f.s.Usage = f.defaultUsage
	return &f
}

func (f *MultiFlag) BoolVars(p *bool, names []string, value bool, usage string) {
	names = uniqStr(names)
	for i := 0; i < len(names); i++ {
		f.fs().BoolVar(p, names[i], value, usage)
	}
}

func (f *MultiFlag) StringVars(p *string, names []string, value string, usage string) {}

func (f *MultiFlag) PrintDefaults() {
	var isZeroValueErrs []error
	f.VisitAll(func(flag2 *flag.Flag) {
		var b strings.Builder
		_, _ = fmt.Fprintf(&b, "  -%s", flag2.Name) // Two spaces before -; see next two comments.
		name, usage := flag.UnquoteUsage(flag2)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}
		if b.Len() <= 4 {
			b.WriteString("\t")
		} else {
			b.WriteString("\n    \t")
		}
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

		if isZero, err := isZeroValue(flag2, flag2.DefValue); err != nil {
			isZeroValueErrs = append(isZeroValueErrs, err)
		} else if !isZero {
			// if _, ok := flag2.Value.(*stringValue); ok {
			// 	_, _ = fmt.Fprintf(&b, " (default %q)", flag2.DefValue)
			// } else {
			_, _ = fmt.Fprintf(&b, " (default %v)", flag2.DefValue)
			// }
		}
		_, _ = fmt.Fprint(f.Output(), b.String(), "\n")
	})
	if errs := isZeroValueErrs; len(errs) > 0 {
		_, _ = fmt.Fprintln(f.Output())
		for _, err := range errs {
			_, _ = fmt.Fprintln(f.Output(), err)
		}
	}
}

func isZeroValue(flag *flag.Flag, value string) (ok bool, err error) {
	return flag.Value.String() == value, nil
}

func (f *MultiFlag) fs() *flag.FlagSet {
	if f.s == nil {
		f.s = flag.NewFlagSet("", flag.ExitOnError)
	}
	return f.s
}

func (f *MultiFlag) defaultUsage() {
	if len(f.s.Name()) == 0 {
		_, _ = fmt.Fprintf(f.Output(), "Usage:\n")
	} else {
		_, _ = fmt.Fprintf(f.Output(), "Usage of %s:\n", f.s.Name())
	}
	f.PrintDefaults()
}

func uniqStr(a []string) []string {
	sort.Strings(a)
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			copy(a[i-1:], a[i:])
			n--
		}
	}
	return a[:n]
}
