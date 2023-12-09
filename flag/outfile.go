package main

import (
	"flag"
	"io"
	"os"
)

// Outfile is a flag-compatible type that can be used to specify output files.
// It implements flag.Value and io.WriteCloser.
//
// Supports setting "stdout" and "stderr" as file names, but in this case Close
// will be a no-op, to avoid closing the standard streams and missing
// panics/crashes the runtime might want to report.
type Outfile struct {
	io.WriteCloser
	name string
}

func (f *Outfile) IsSet() bool    { return f.name != "" }
func (f *Outfile) String() string { return f.name }

func (f *Outfile) Set(s string) error {
	f.name = s
	switch s {
	case "stdout":
		f.WriteCloser = nopCloser{os.Stdout}
	case "stderr":
		f.WriteCloser = nopCloser{os.Stderr}
	default:
		fd, err := os.Create(s)
		if err != nil {
			return err
		}
		f.WriteCloser = fd
	}
	return nil
}

type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

func main() {
	var disasmLog Outfile
	flag.Var(&disasmLog, "dbglog", "write execution log to [file|stdout|stderr] (for testing/debugging")
	flag.Parse()

	if disasmLog.IsSet() {
		defer disasmLog.Close()
		disasmLog.Write([]byte("hello world\n"))
	}
}
