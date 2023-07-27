package multiflag

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

type MultiFlag struct {
	s     *flag.FlagSet
	alias map[string][]string
}

func NewMultiFlag(name string, errorHandling flag.ErrorHandling) *MultiFlag {
	f := MultiFlag{
		s:     flag.NewFlagSet(name, errorHandling),
		alias: make(map[string][]string),
	}
	f.s.Usage = f.defaultUsage
	return &f
}

func (f *MultiFlag) BoolVars(p *bool, names []string, value bool, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) StringVars(p *string, names []string, value string, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) PrintDefaults() {
	var isZeroValueErrs []error
	var skip = map[string]struct{}{}
	f.VisitAll(func(flag2 *flag.Flag) {
		if _, ok := skip[flag2.Name]; ok {
			return
		}
		var b strings.Builder
		_, _ = fmt.Fprintf(&b, "  -%s", flag2.Name)
		if aliases, ok := f.alias[flag2.Name]; ok {
			for i := 0; i < len(aliases); i++ {
				a := aliases[i]
				if _, ok = skip[a]; ok || a == flag2.Name {
					continue
				}
				_, _ = fmt.Fprintf(&b, ", -%s", a)
				skip[a] = struct{}{}
			}
		}
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

func (f *MultiFlag) var_(p any, names []string, value any, usage string) {
	if names = uniqStr(names); len(names) > 0 {
		key := names[0]
		for i := 0; i < len(names); i++ {
			switch value.(type) {
			case bool:
				f.fs().BoolVar(p.(*bool), names[i], value.(bool), usage)
			case string:
				f.fs().StringVar(p.(*string), names[i], value.(string), usage)
			}
			f.alias[key] = append(f.alias[key], names[i])
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
	if f.alias == nil {
		f.alias = make(map[string][]string)
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
