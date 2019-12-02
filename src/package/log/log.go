package log

import "errors"

type Log interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Prefix() string
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	SetFlags(flag int)
	// SetOutput(w io.Writer)
	SetPrefix(prefix string)
	// Writer() io.Writer
}

var Import func() Log = func() Log {
	panic(errors.New("Package 'log' not delegated."))
}

func init() {

}
