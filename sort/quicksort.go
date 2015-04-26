package sort

func Quicksort(lista []int) []int {
	if len(lista) <= 1 {
		return lista
	}

	return sort(lista)
}

func sort(lista []int) []int {

	numeros := make([]int, len(lista))
	copy(numeros, lista)
	
	indexPivo := len(numeros) / 2
	
	pivo := numeros[indexPivo]
	
	numeros = append(numeros[:indexPivo], numeros[indexPivo+1:]...)
	
	menores, maiores := particionar(numeros, pivo)
	
	return append(
				append(Quicksort(menores), pivo),
				Quicksort(maiores)...)
}

func particionar(numeros[] int, pivo int) (menores [] int, maiores[] int) {

	for _,n := range numeros {

		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}