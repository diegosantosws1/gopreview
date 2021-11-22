package hashtags_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/diegosantosws1/gopreview/hashtags"
)

var (
	caseTest = flag.String("case", "", "specifies which test case to request")
	update   = flag.Bool("update", false, "update result file")
)

func TestGetHashtags(t *testing.T) {
	casesTests := []struct {
		Name    string
		Msg     string
		OutFile string
		LenExp  int
	}{
		{
			Name:   "1 hastag",
			Msg:    "Olá #voce, vamos nos falar mais",
			LenExp: 1,
		},
		{
			Name:   "2 hastag",
			Msg:    "#EuAmoSP uhuuuuu!!! #VamosColocarNessaVida",
			LenExp: 2,
		},
		{
			Name:   "3 hastag",
			Msg:    "Eu quero ver você #aqui comigo #ILove, #fechadocomvc",
			LenExp: 3,
		},
	}

	for _, tc := range casesTests {
		t.Run(tc.Name, func(t *testing.T) {
			testGetHashtags(t, tc.Name, tc.Msg, tc.OutFile, tc.LenExp)
		})
	}
}

func testGetHashtags(t *testing.T, name, msg, outFile string, lenExp int) {
	if len(*caseTest) > 0 {
		if !strings.Contains(outFile, *caseTest) {
			t.Skipf("skiping test [%s] file [%s]. Does not match [%s]", name, outFile, *caseTest)
			return
		}
	}

	tags, _ := hashtags.GetHashtags(msg)
	if len(tags) != lenExp {
		t.Errorf("[] expected [%v] got [%v] ", len(tags), lenExp)
	}
}

func TestGetSentenceWithoutHashtag(t *testing.T) {
	casesTests := []struct {
		Name    string
		Msg     string
		OutFile string
		LenExp  int
	}{
		{
			Name:   "1 hastag",
			Msg:    "Olá #voce, vamos nos falar mais",
			LenExp: 1,
		},
		{
			Name:   "2 hastag",
			Msg:    "#EuAmoSP uhuuuuu!!! #VamosColocarNessaVida",
			LenExp: 2,
		},
		{
			Name:   "3 hastag",
			Msg:    "Eu quero ver você #aqui comigo #ILove, #fechadocomvc",
			LenExp: 3,
		},
	}

	for _, tc := range casesTests {
		t.Run(tc.Name, func(t *testing.T) {
			testGetSentenceWithoutHashtag(t, tc.Name, tc.Msg, tc.OutFile, tc.LenExp)
		})
	}
}

func testGetSentenceWithoutHashtag(t *testing.T, name, msg, outFile string, lenExp int) {
	if len(*caseTest) > 0 {
		if !strings.Contains(outFile, *caseTest) {
			t.Skipf("skiping test [%s] file [%s]. Does not match [%s]", name, outFile, *caseTest)
			return
		}
	}

	tags, _ := hashtags.GetSentenceWithoutHashtag(msg)
	if len(tags) != lenExp {
		t.Errorf("[] expected [%v] got [%v] ", len(tags), lenExp)
	}
}
