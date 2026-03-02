package main

import "fmt"

func main() {
	// Criamos a blockchain — o bloco gênesis já será minerado
	bc := NovaBlockchain()

	// Adicionamos blocos — cada um exige mineração (PoW)
	bc.AdicionarBloco("Maria enviou 5 moedas para João")
	bc.AdicionarBloco("João enviou 2 moedas para Ana")

	// Imprimimos a cadeia com validação de cada bloco
	for i, bloco := range bc.blocos {
		pow := NovoProofOfWork(bloco)

		fmt.Printf("--- Bloco %d ---\n", i)
		fmt.Printf("Dados:         %s\n", bloco.Data)
		fmt.Printf("Hash:          %x\n", bloco.Hash)
		fmt.Printf("Hash anterior: %x\n", bloco.HashAnterior)
		fmt.Printf("Nonce:         %d\n", bloco.Nonce)
		fmt.Printf("PoW válido?    %t\n", pow.Validar())
		fmt.Println()
	}
}