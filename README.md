# Usage 


### Install

```bash
$ go get github.com/diegosantosws1/gopreview
```

# Example

Your project you import how on example

```go
func main() {
    url := "https://www.infomoney.com.br/mercados/ambev-abev3-tem-preco-alvo-elevado-pelo-morgan-stanley-mas-analistas-do-banco-seguem-ceticos-com-acao/"
    res, err := gopreview.Parser(url)
    if err != nil {
        log.Printf("%v", err)
        return
    }
}
```

### Output

```json
{
	"title": "Ambev (ABEV3) tem preço-alvo elevado pelo Morgan Stanley, mas analistas do banco seguem céticos com ação",
	"description": "Apesar de alta do volume vendido no terceiro trimestre, concorrência e crise na América Latina são riscos para 2022, aponta a análise",
	"image": "https://www.infomoney.com.br/wp-content/uploads/2020/09/GettyImages-542639088.jpg?quality=70",
	"url": "https://www.infomoney.com.br/mercados/ambev-abev3-tem-preco-alvo-elevado-pelo-morgan-stanley-mas-analistas-do-banco-seguem-ceticos-com-acao/"
}
```