package impl

import (
	. "log"

	def "github.com/emuggie/dipping/package/log"
)

type delegator struct {
}

func (delegator) Fatal(v ...interface{}) {
	Fatal(v...)
}
func (delegator) Fatalf(format string, v ...interface{}) {
	Fatalf(format, v...)
}
func (delegator) Fatalln(v ...interface{}) {
	Fatalln(v...)
}
func (delegator) Flags() int {
	return Flags()
}
func (delegator) Output(calldepth int, s string) error {
	return Output(calldepth, s)
}
func (delegator) Panic(v ...interface{}) {
	Panic(v...)
}
func (delegator) Panicf(format string, v ...interface{}) {
	Panicf(format, v...)
}
func (delegator) Panicln(v ...interface{}) {
	Panicln(v...)
}
func (delegator) Prefix() string {
	return Prefix()
}
func (delegator) Print(v ...interface{}) {
	Print(v...)
}
func (delegator) Printf(format string, v ...interface{}) {
	Printf(format, v...)
}
func (delegator) Println(v ...interface{}) {
	Println(v...)
}
func (delegator) SetFlags(flag int) {
	SetFlags(flag)
}

//func (delegator) SetOutput(w io.Writer){}
func (delegator) SetPrefix(prefix string) {
	SetPrefix(prefix)
}

//func (delegator) Writer() io.Writer{}

var Delegator def.Log

func init() {
	Delegator = new(delegator)
}

func Delegate() {
	def.Import = func() def.Log {
		return Delegator
	}
}
