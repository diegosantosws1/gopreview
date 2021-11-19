package metadata_test

import (
	"testing"

	"github.com/DiegoSantosWS/gopreview/metadata"
	"github.com/google/go-cmp/cmp"
)

func compareResult(t *testing.T, tcName string, exp, got metadata.HTMLMeta) {
	if diff := cmp.Diff(exp, got); len(diff) > 0 {
		t.Errorf(" test [%s] mismatch (-want +got):\n%s", tcName, diff)
	}
}
