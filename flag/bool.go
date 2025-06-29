package main

import (
	"flag"
	"fmt"
	"strconv"
)

type BoolFlag struct{ val, set bool }

func (f BoolFlag) String() string {
	if !f.set {
		return "<unset>"
	}
	return strconv.FormatBool(f.val)
}

func (f *BoolFlag) Set(value string) error {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("invalid boolean value: %s", value)
	}
	f.val = v
	f.set = true
	return nil
}

// IsSet reports whether the flag has been set.
func (f *BoolFlag) IsSet() bool { return f.set }

// Value reports the flag value if it has been set, else always false.
func (f *BoolFlag) Value() bool { return f.val }

var battery BoolFlag

func main() {
	flag.Var(&battery, "b", "Filter by presence of battery-packed RAM")
	flag.Parse()

	if battery.IsSet() {
		fmt.Println("battery value is", battery.Value())
	} else {
		fmt.Println("battery value is", battery.String())
	}
}
