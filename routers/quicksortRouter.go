package routers

import(
	"fmt"
	"os"
	"net/http"
	"time"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
	"github.com/msantocardoso/golang/sort"
	"github.com/msantocardoso/golang/util"
)

func QuicksortRouter(w http.ResponseWriter, r *http.Request){
	
	params := mux.Vars(r)

	tamanho,err := strconv.Atoi(params["tamanhoLista"])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	var lista []int = make([]int, tamanho)

	util.Cronometrar("gerarLista",func() {
		lista = gerarLista(tamanho)
	})

	var ordenada []int = make([]int, len(lista))
	
	util.Cronometrar("quicksort",func() {
		ordenada=sort.Quicksort(lista)
	})

	fmt.Fprintln(w,"{lista:",ordenada,"}")
}

func gerarLista(tamanho int) []int {
	rand.Seed(time.Now().UnixNano())

	itens := make([]int, tamanho)

	for i := 1; i<tamanho; i++ {
		itens[i] = rand.Intn(tamanho)
	}

	return itens	
}