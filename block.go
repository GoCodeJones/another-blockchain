package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
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
}

// calcularHash gera o hash SHA-256 do bloco.
// Concatenamos todos os campos relevantes em bytes e aplicamos SHA-256.
// Se qualquer campo mudar (inclusive o timestamp), o hash muda completamente.
func (b *Block) calcularHash() {
	// Convertemos o timestamp para string e depois para bytes
	timestamp := []byte(fmt.Sprintf("%d", b.Timestamp))

	// Juntamos todos os campos do bloco em um único slice de bytes
	cabecalho := bytes.Join(
		[][]byte{b.HashAnterior, b.Data, timestamp},
		[]byte{}, // separador vazio — apenas concatena
	)

	// Aplicamos SHA-256 e armazenamos o resultado em b.Hash
	hash := sha256.Sum256(cabecalho)
	b.Hash = hash[:] // hash[:] converte [32]byte para []byte
}

// NovoBloco cria e retorna um novo bloco com os dados fornecidos.
// Recebe os dados do bloco e o hash do bloco anterior na cadeia.
func NovoBloco(dados string, hashAnterior []byte) *Block {
	bloco := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         []byte(dados),
		HashAnterior: hashAnterior,
		Hash:         []byte{},
	}

	// Calculamos o hash logo após criar o bloco
	bloco.calcularHash()

	return bloco
}