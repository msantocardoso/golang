package routers

import(
	"fmt"
	"os"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/msantocardoso/golang/sort"
)

func QuicksortRouter(w http.ResponseWriter, r *http.Request){
	
	params := mux.Vars(r)

	tamanho,err := strconv.Atoi(params["tamanhoLista"])
 	if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	lista := sort.Quicksort(tamanho)

	fmt.Fprintln(w,"{lista:",lista,"}")
}