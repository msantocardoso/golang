package sort

import "math/rand"

func Quicksort(tamanhoLista int)[]float64 {

	itens := make([]float64, tamanhoLista)

	i := 1
	for ; i<tamanhoLista; i++ {
		itens[i] = (rand.Float64())
	}
	return itens
}