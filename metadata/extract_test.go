package metadata_test

import (
	"strings"
	"testing"

	"github.com/diegosantosws1/gopreview/metadata"
)

func TestExtractMetadata(t *testing.T) {
	testCases := []struct {
		Name    string
		URL     string
		InFile  string
		OutFile string
	}{
		{
			Name:    "Wsite brasil-blog",
			URL:     "https://www.wsitebrasil.com.br/blog/post/marketing-digital/dicas-de-como-aumentar-o-trafego-organico-no-seu-blog/132",
			InFile:  "./testdata/wsite/01/data.html",
			OutFile: "./testdata/wsite/01/result.json",
		},
		{
			Name:    "Instagram",
			URL:     "https://www.instagram.com/diegosantos132/",
			InFile:  "./testdata/insta/01/data.html",
			OutFile: "./testdata/insta/01/result.json",
		},
		{
			Name:    "Instagram - foto",
			URL:     "https://www.instagram.com/p/CUSOVlqLIC9a-2IkDasVjaj5bPxB015fFYj8RM0/",
			InFile:  "./testdata/insta/02/data.html",
			OutFile: "./testdata/insta/02/result.json",
		},
		{
			Name:    "Site - foto",
			URL:     "https://diegosantosws.com.br/posts/2021/06/convertendo-datas-em-string-usando-golang/",
			InFile:  "./testdata/site/01/data.html",
			OutFile: "./testdata/site/01/result.json",
		},
		{
			Name:    "Site - foto",
			URL:     "https://diegosantosws.com.br/posts/2021/06/artigos/",
			InFile:  "./testdata/site/01/data.html",
			OutFile: "./testdata/site/01/result.json",
		},
		{
			Name:    "Infomoney - foto",
			URL:     "https://www.infomoney.com.br/mercados/ambev-abev3-tem-preco-alvo-elevado-pelo-morgan-stanley-mas-analistas-do-banco-seguem-ceticos-com-acao/",
			InFile:  "./testdata/site/02/data.html",
			OutFile: "./testdata/site/02/result.json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testExtractMetadata(t, tc.Name, tc.URL, tc.InFile, tc.OutFile)
		})
	}
}

func testExtractMetadata(t *testing.T, name, url, inFile, outFile string) {
	if len(*caseTest) > 0 {
		if !strings.Contains(outFile, *caseTest) {
			t.Skipf("skiping test [%s] file [%s]. Does not match [%s]", name, outFile, *caseTest)
			return
		}
	}

	if *htmlcreate {
		read, err := metadata.NewRequest(url)
		if err != nil {
			t.Fatal(err)
		}
		createFileFromReader(t, inFile, read)
		return
	}

	r, _ := readFile(t, inFile)
	metas, _ := metadata.ExtractMatadata(r, "meta")
	if *update {
		createJSONFile(t, outFile, metas, true)
		return
	}

}
