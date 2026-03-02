package main

import (
	"time"
)

// Bloco representa uma unidade de dados na blockchain.
// Cada bloco contém seus próprios dados e uma referência ao bloco anterior,
// formando assim a "corrente" (chain) de blocos.
type Block struct {
	// Timestamp é o momento exato em que o bloco foi criado (Unix timestamp)
	Timestamp int64

	// Data são os dados armazenados neste bloco (ex: transações)
	Data []byte

	// HashAnterior é o hash do bloco que vem antes na cadeia.
	// É este campo que "liga" os blocos entre si.
	HashAnterior []byte

	// Hash é a impressão digital única deste bloco,
	// calculada a partir dos seus próprios dados.
	Hash []byte

	// Nonce é o número que o minerador varia até encontrar um hash válido.
	// É a "prova" de que trabalho computacional foi realizado.
	Nonce int
}

// NovoBloco cria e retorna um novo bloco com os dados fornecidos.
// Agora executa o Proof of Work para encontrar um hash válido antes de retornar.
func NovoBloco(dados string, hashAnterior []byte) *Block {
	bloco := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         []byte(dados),
		HashAnterior: hashAnterior,
		Hash:         []byte{},
		Nonce:        0,
	}

	// Executamos o PoW — isso pode levar alguns milissegundos
	pow := NovoProofOfWork(bloco)
	nonce, hash := pow.Executar()

	bloco.Nonce = nonce
	bloco.Hash = hash

	return bloco
}