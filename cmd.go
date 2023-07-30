package multiflag

import (
	"flag"
	"os"
	"time"
)

var CommandLine = NewMultiFlag(os.Args[0], flag.ExitOnError)

func BoolsVar(p *bool, names []string, value bool, usage string) {
	CommandLine.BoolsVar(p, names, value, usage)
}

func Bools(names []string, value bool, usage string) *bool {
	return CommandLine.Bools(names, value, usage)
}

func IntsVar(p *int, names []string, value int, usage string) {
	CommandLine.IntsVar(p, names, value, usage)
}

func Ints(names []string, value int, usage string) *int {
	return CommandLine.Ints(names, value, usage)
}

func Ints64Var(p *int64, names []string, value int64, usage string) {
	CommandLine.Ints64Var(p, names, value, usage)
}

func Ints64(names []string, value int64, usage string) *int64 {
	return CommandLine.Ints64(names, value, usage)
}

func UintsVar(p *uint, names []string, value uint, usage string) {
	CommandLine.UintsVar(p, names, value, usage)
}

func Uints(names []string, value uint, usage string) *uint {
	return CommandLine.Uints(names, value, usage)
}

func Uints64Var(p *uint64, names []string, value uint64, usage string) {
	CommandLine.Uints64Var(p, names, value, usage)
}

func Uints64(names []string, value uint64, usage string) *uint64 {
	return CommandLine.Uints64(names, value, usage)
}

func Floats64Var(p *float64, names []string, value float64, usage string) {
	CommandLine.Floats64Var(p, names, value, usage)
}

func Floats64(names []string, value float64, usage string) *float64 {
	return CommandLine.Floats64(names, value, usage)
}

func StringsVar(p *string, names []string, value string, usage string) {
	CommandLine.StringsVar(p, names, value, usage)
}

func Strings(names []string, value string, usage string) *string {
	return CommandLine.Strings(names, value, usage)
}

func DurationsVar(p *time.Duration, names []string, value time.Duration, usage string) {
	CommandLine.DurationsVar(p, names, value, usage)
}

func Durations(names []string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Durations(names, value, usage)
}

func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.BoolVar(p, name, value, usage)
}

func Bool(name string, value bool, usage string) *bool {
	return CommandLine.Bool(name, value, usage)
}

func IntVar(p *int, name string, value int, usage string) {
	CommandLine.IntVar(p, name, value, usage)
}

func Int(name string, value int, usage string) *int {
	return CommandLine.Int(name, value, usage)
}

func Int64Var(p *int64, name string, value int64, usage string) {
	CommandLine.Int64Var(p, name, value, usage)
}

func Int64(name string, value int64, usage string) *int64 {
	return CommandLine.Int64(name, value, usage)
}

func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.UintVar(p, name, value, usage)
}

func Uint(name string, value uint, usage string) *uint {
	return CommandLine.Uint(name, value, usage)
}

func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Uint64Var(p, name, value, usage)
}

func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}

func Float64Var(p *float64, name string, value float64, usage string) {
	CommandLine.Float64Var(p, name, value, usage)
}

func Float64(name string, value float64, usage string) *float64 {
	return CommandLine.Float64(name, value, usage)
}

func StringVar(p *string, name string, value string, usage string) {
	CommandLine.StringVar(p, name, value, usage)
}

func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}

func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	CommandLine.DurationVar(p, name, value, usage)
}

func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}

func Parse() {
	_ = CommandLine.Parse(os.Args[1:])
}

func Parsed() bool {
	return CommandLine.Parsed()
}

var _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = BoolsVar,
	Bools, IntsVar, Ints, Ints64Var, Ints64, UintsVar, Uints, Uints64Var, Uints64, Floats64Var, Floats64,
	StringsVar, Strings, DurationsVar, Durations, BoolVar, Bool, IntVar, Int, Int64Var, Int64, UintVar,
	Uint, Uint64Var, Uint64, Float64Var, Float64, StringVar, String, DurationVar, Duration, Parse, Parsed
