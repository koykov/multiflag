package multiflag

import (
	"encoding"
	"flag"
	"io"
	"time"
)

func (f *MultiFlag) Output() io.Writer                  { return f.fs().Output() }
func (f *MultiFlag) Name() string                       { return f.fs().Name() }
func (f *MultiFlag) ErrorHandling() flag.ErrorHandling  { return f.fs().ErrorHandling() }
func (f *MultiFlag) SetOutput(output io.Writer)         { f.fs().SetOutput(output) }
func (f *MultiFlag) VisitAll(fn func(flag2 *flag.Flag)) { f.fs().VisitAll(fn) }
func (f *MultiFlag) Visit(fn func(*flag.Flag))          { f.fs().Visit(fn) }
func (f *MultiFlag) Lookup(name string) *flag.Flag      { return f.fs().Lookup(name) }
func (f *MultiFlag) Set(name, value string) error       { return f.fs().Set(name, value) }
func (f *MultiFlag) NFlag() int                         { return f.fs().NFlag() }
func (f *MultiFlag) Arg(i int) string                   { return f.fs().Arg(i) }
func (f *MultiFlag) NArg() int                          { return f.fs().NArg() }
func (f *MultiFlag) Args() []string                     { return f.fs().Args() }

func (f *MultiFlag) BoolVar(p *bool, name string, value bool, usage string) {
	f.fs().BoolVar(p, name, value, usage)
}

func (f *MultiFlag) Bool(name string, value bool, usage string) *bool {
	return f.fs().Bool(name, value, usage)
}

func (f *MultiFlag) IntVar(p *int, name string, value int, usage string) {
	f.fs().IntVar(p, name, value, usage)
}

func (f *MultiFlag) Int(name string, value int, usage string) *int {
	return f.fs().Int(name, value, usage)
}

func (f *MultiFlag) Int64Var(p *int64, name string, value int64, usage string) {
	f.fs().Int64Var(p, name, value, usage)
}

func (f *MultiFlag) Int64(name string, value int64, usage string) *int64 {
	return f.fs().Int64(name, value, usage)
}

func (f *MultiFlag) UintVar(p *uint, name string, value uint, usage string) {
	f.fs().UintVar(p, name, value, usage)
}

func (f *MultiFlag) Uint(name string, value uint, usage string) *uint {
	return f.fs().Uint(name, value, usage)
}

func (f *MultiFlag) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.fs().Uint64Var(p, name, value, usage)
}

func (f *MultiFlag) Uint64(name string, value uint64, usage string) *uint64 {
	return f.fs().Uint64(name, value, usage)
}

func (f *MultiFlag) StringVar(p *string, name string, value string, usage string) {
	f.fs().StringVar(p, name, value, usage)
}

func (f *MultiFlag) String(name string, value string, usage string) *string {
	return f.fs().String(name, value, usage)
}

func (f *MultiFlag) Float64Var(p *float64, name string, value float64, usage string) {
	f.fs().Float64Var(p, name, value, usage)
}

func (f *MultiFlag) Float64(name string, value float64, usage string) *float64 {
	return f.fs().Float64(name, value, usage)
}

func (f *MultiFlag) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.fs().DurationVar(p, name, value, usage)
}

func (f *MultiFlag) Duration(name string, value time.Duration, usage string) *time.Duration {
	return f.fs().Duration(name, value, usage)
}

func (f *MultiFlag) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string) {
	f.fs().TextVar(p, name, value, usage)
}

func (f *MultiFlag) Func(name, usage string, fn func(string) error)  { f.fs().Func(name, usage, fn) }
func (f *MultiFlag) Var(value flag.Value, name string, usage string) { f.fs().Var(value, name, usage) }
func (f *MultiFlag) Parse(arguments []string) error                  { return f.fs().Parse(arguments) }
func (f *MultiFlag) Parsed() bool                                    { return f.fs().Parsed() }

func (f *MultiFlag) Init(name string, errorHandling flag.ErrorHandling) {
	f.fs().Init(name, errorHandling)
}

func (f *MultiFlag) Usage() { f.s.Usage() }
