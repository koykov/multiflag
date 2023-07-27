package multiflag

import (
	"flag"
	"os"
	"time"
)

var CommandLine = NewMultiFlag(os.Args[0], flag.ExitOnError)

func BoolVars(p *bool, names []string, value bool, usage string) {
	CommandLine.BoolVars(p, names, value, usage)
}

func Bools(names []string, value bool, usage string) *bool {
	return CommandLine.Bools(names, value, usage)
}

func IntVars(p *int, names []string, value int, usage string) {
	CommandLine.IntVars(p, names, value, usage)
}

func Ints(names []string, value int, usage string) *int {
	return CommandLine.Ints(names, value, usage)
}

func Int64Vars(p *int64, names []string, value int64, usage string) {
	CommandLine.Int64Vars(p, names, value, usage)
}

func Ints64(names []string, value int64, usage string) *int64 {
	return CommandLine.Ints64(names, value, usage)
}

func UintVars(p *uint, names []string, value uint, usage string) {
	CommandLine.UintVars(p, names, value, usage)
}

func Uints(names []string, value uint, usage string) *uint {
	return CommandLine.Uints(names, value, usage)
}

func Uint64Vars(p *uint64, names []string, value uint64, usage string) {
	CommandLine.Uint64Vars(p, names, value, usage)
}

func Uints64(names []string, value uint64, usage string) *uint64 {
	return CommandLine.Uints64(names, value, usage)
}

func Float64Vars(p *float64, names []string, value float64, usage string) {
	CommandLine.Float64Vars(p, names, value, usage)
}

func Floats64(names []string, value float64, usage string) *float64 {
	return CommandLine.Floats64(names, value, usage)
}

func StringVars(p *string, names []string, value string, usage string) {
	CommandLine.StringVars(p, names, value, usage)
}

func Strings(names []string, value string, usage string) *string {
	return CommandLine.Strings(names, value, usage)
}

func DurationVars(p *time.Duration, names []string, value time.Duration, usage string) {
	CommandLine.DurationVars(p, names, value, usage)
}

func Durations(names []string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Durations(names, value, usage)
}

func Parse() {
	_ = CommandLine.Parse(os.Args[1:])
}

func Parsed() bool {
	return CommandLine.Parsed()
}

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = BoolVars, Bools, IntVars, Ints, Int64Vars, Ints64, UintVars,
	Uints, Uint64Vars, Uints64, Float64Vars, Floats64, StringVars, Strings, DurationVars, Durations, Parse, Parsed
