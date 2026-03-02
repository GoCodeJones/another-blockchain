package main

import "fmt"

func main() {
	// Criamos um bloco de teste para ver o hash em ação.
	// O hash do bloco gênesis usa hashAnterior vazio ([]byte{})
	bloco := NovoBloco("Bloco de teste", []byte{})

	fmt.Printf("Dados:         %s\n", bloco.Data)
	fmt.Printf("Hash:          %x\n", bloco.Hash)
	fmt.Printf("Hash anterior: %x\n", bloco.HashAnterior)
	fmt.Printf("Timestamp:     %d\n", bloco.Timestamp)
}