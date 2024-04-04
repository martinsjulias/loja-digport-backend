package model

//import "fmt"

type Carrinho struct {
	ID           string
	UsuarioId    string
	Quantidade   int
	ValorTotal   float64
	InfosProduto []InfosProduto
}

type InfosProduto struct {
	ProdutoId           string
	QuantidadeDeProduto int
}
