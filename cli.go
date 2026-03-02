package main

import (
	"flag"
	"fmt"
	"os"
)

// CLI gerencia os comandos de linha de comando da blockchain.
type CLI struct {
	bc *Blockchain
}

// uso imprime as instruções de uso e encerra o programa.
func (cli *CLI) uso() {
	fmt.Println("Uso:")
	fmt.Println("  addblock -data DADOS   → adiciona um novo bloco à cadeia")
	fmt.Println("  printchain             → imprime todos os blocos da cadeia")
	os.Exit(1)
}

// validarArgumentos garante que pelo menos um comando foi fornecido.
func (cli *CLI) validarArgumentos() {
	if len(os.Args) < 2 {
		cli.uso()
	}
}

// adicionarBloco lida com o comando addblock.
func (cli *CLI) adicionarBloco(dados string) {
	cli.bc.AdicionarBloco(dados)
	fmt.Println("Bloco adicionado com sucesso!")
}

// imprimirCadeia lida com o comando printchain.
func (cli *CLI) imprimirCadeia() {
	for i, bloco := range cli.bc.blocos {
		pow := NovoProofOfWork(bloco)

		fmt.Printf("============ Bloco %d ============\n", i)
		fmt.Printf("Dados:         %s\n", bloco.Data)
		fmt.Printf("Hash:          %x\n", bloco.Hash)
		fmt.Printf("Hash anterior: %x\n", bloco.HashAnterior)
		fmt.Printf("Nonce:         %d\n", bloco.Nonce)
		fmt.Printf("PoW válido?    %t\n", pow.Validar())
		fmt.Println()
	}
}

// Executar é o ponto de entrada do CLI.
// Faz o parsing dos argumentos e chama o comando correto.
func (cli *CLI) Executar() {
	cli.validarArgumentos()

	// Definimos os subcomandos com seus flags
	cmdAddBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	cmdPrintChain := flag.NewFlagSet("printchain", flag.ExitOnError)

	// Flag -data para o comando addblock
	addBlockData := cmdAddBlock.String("data", "", "Dados do bloco")

	// O primeiro argumento (os.Args[1]) determina o comando
	switch os.Args[1] {

	case "addblock":
		// Faz o parse dos args restantes para capturar o -data
		cmdAddBlock.Parse(os.Args[2:])

	case "printchain":
		cmdPrintChain.Parse(os.Args[2:])

	default:
		cli.uso()
	}

	// Executa o comando correspondente
	if cmdAddBlock.Parsed() {
		if *addBlockData == "" {
			cmdAddBlock.Usage()
			os.Exit(1)
		}
		cli.adicionarBloco(*addBlockData)
	}

	if cmdPrintChain.Parsed() {
		cli.imprimirCadeia()
	}
}