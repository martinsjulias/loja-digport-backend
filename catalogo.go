package main

import (
	"github.com/martinsjulias/loja-digport-backend/model" // / nome do pacote que queremos importar
)

func catalogo() []model.Produto {
	produtos := []model.Produto{
		{
			Nome:       " Blusa",
			Descricao:  " Blusa Rosa",
			Categoria:  " Blusa",
			ID:         " 1",
			Valor:      20.00,
			Quantidade: 2,
		},

		{
			Nome:       " Regata",
			Descricao:  " Regata Roxa",
			Categoria:  " Regata",
			ID:         " 2",
			Valor:      15.00,
			Quantidade: 3,
		},
	}
	return produtos
}
