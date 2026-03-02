package main

// main é o ponto de entrada do programa.
// Cria a blockchain e delega o controle para o CLI.
func main() {
	bc := NovaBlockchain()

	cli := CLI{bc: bc}
	cli.Executar()
}