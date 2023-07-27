package multiflag

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"
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

func (f *MultiFlag) Bools(names []string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVars(p, names, value, usage)
	return p
}

func (f *MultiFlag) IntVars(p *int, names []string, value int, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Ints(names []string, value int, usage string) *int {
	p := new(int)
	f.IntVars(p, names, value, usage)
	return p
}

func (f *MultiFlag) Int64Vars(p *int64, names []string, value int64, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Ints64(names []string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Vars(p, names, value, usage)
	return p
}

func (f *MultiFlag) UintVars(p *uint, names []string, value uint, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Uints(names []string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVars(p, names, value, usage)
	return p
}

func (f *MultiFlag) Uint64Vars(p *uint64, names []string, value uint64, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Uints64(names []string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Vars(p, names, value, usage)
	return p
}

func (f *MultiFlag) Float64Vars(p *float64, names []string, value float64, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Floats64(names []string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Vars(p, names, value, usage)
	return p
}

func (f *MultiFlag) StringVars(p *string, names []string, value string, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Strings(names []string, value string, usage string) *string {
	p := new(string)
	f.StringVars(p, names, value, usage)
	return p
}

func (f *MultiFlag) DurationVars(p *time.Duration, names []string, value time.Duration, usage string) {
	f.var_(p, names, value, usage)
}

func (f *MultiFlag) Durations(names []string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVars(p, names, value, usage)
	return p
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
			_, _ = fmt.Fprintf(&b, " (default %v)", flag2.DefValue)
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
			case int:
				f.fs().IntVar(p.(*int), names[i], value.(int), usage)
			case int64:
				f.fs().Int64Var(p.(*int64), names[i], value.(int64), usage)
			case uint:
				f.fs().UintVar(p.(*uint), names[i], value.(uint), usage)
			case uint64:
				f.fs().Uint64Var(p.(*uint64), names[i], value.(uint64), usage)
			case float64:
				f.fs().Float64Var(p.(*float64), names[i], value.(float64), usage)
			case string:
				f.fs().StringVar(p.(*string), names[i], value.(string), usage)
			case time.Duration:
				f.fs().DurationVar(p.(*time.Duration), names[i], value.(time.Duration), usage)
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
