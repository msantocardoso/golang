package util
import (
	"time"
	"fmt"
)
func Cronometrar(operacao string, funcao func()) {
	inicial := time.Now()

	funcao()

	fmt.Printf("\nTempo de execução [%s]: %s\n", operacao, time.Since(inicial))
}