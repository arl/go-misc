package testutil

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var updateGolden = flag.Bool("update", false, "update golden files")

// DiffBytesWithGolden is a test helper that compares some bytes with a file
// content whose path is provided in the 'goldenFile' argument. If -update flag
// is given to go test, the golden file is updated with the new content,
// provided by 'buf'.
func DiffBytesWithGolden(tb testing.TB, buf []byte, goldenFile string) {
	tb.Helper()

	if *updateGolden {
		if err := os.WriteFile(goldenFile, buf, 0644); err != nil {
			tb.Fatalf("can't update golden file %s: %v", goldenFile, err)
		}
		return
	}

	goldenBuf, err := os.ReadFile(goldenFile)
	if err != nil {
		tb.Fatalf("can't read golden file %s: %v", goldenFile, err)
		return
	}

	DiffBytes(tb, buf, goldenBuf, "got", filepath.Base(goldenFile))
}

// DiffBytes is a test helper that compares 2 bytes slices. If they differ, the
// test is failed and DiffBytes shows an hexadecimal diff (à la hexdump).
//
// labels[0] and labels[1] can be used to personalize the title diff, instead of
// 'got' and 'want', respectively.
func DiffBytes(tb testing.TB, got, want []byte, gotlabel, wantLabel string) {
	tb.Helper()

	if diff := cmp.Diff(got, want); diff != "" {
		tb.Errorf("actual string differs with golden file\n\n--- %s\n+++ %s\n\n%s", gotlabel, wantLabel, diff)
	}
}
