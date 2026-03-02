package main

// Blockchain representa a cadeia completa de blocos.
// Internamente é apenas uma fatia (slice) de ponteiros para Block.
type Blockchain struct {
	// blocos armazena todos os blocos na ordem em que foram adicionados
	blocos []*Block
}

// NovaBlockchain cria uma blockchain já inicializada com o bloco gênesis.
// O bloco gênesis é o primeiro bloco da cadeia — não tem pai, por isso
// seu HashAnterior é um slice vazio.
func NovaBlockchain() *Blockchain {
	return &Blockchain{
		blocos: []*Block{criarBlocoGenesis()},
	}
}

// criarBlocoGenesis cria o bloco inicial da cadeia.
// Por convenção, seu conteúdo é fixo e seu HashAnterior é vazio.
func criarBlocoGenesis() *Block {
	return NovoBloco("Bloco Gênesis", []byte{})
}

// AdicionarBloco cria um novo bloco com os dados fornecidos
// e o anexa ao final da cadeia.
// O hash do último bloco atual é passado como HashAnterior do novo bloco,
// garantindo a ligação entre eles.
func (bc *Blockchain) AdicionarBloco(dados string) {
	// Pegamos o último bloco da cadeia para obter seu hash
	ultimoBloco := bc.blocos[len(bc.blocos)-1]

	// Criamos o novo bloco referenciando o hash do bloco anterior
	novoBloco := NovoBloco(dados, ultimoBloco.Hash)

	// Adicionamos o novo bloco ao final da cadeia
	bc.blocos = append(bc.blocos, novoBloco)
}