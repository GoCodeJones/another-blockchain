package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// dificuldade define quantos zeros iniciais o hash precisa ter.
// Quanto maior, mais difícil (e lento) é minerar um bloco.
// Valor 16 é razoável para demonstração — Bitcoin usa ~76 zeros hoje.
const dificuldade = 16

// ProofOfWork agrupa o bloco e o alvo (target) do desafio.
type ProofOfWork struct {
	bloco *Block

	// alvo é o valor numérico que o hash precisa ser MENOR para ser válido.
	// Ex: dificuldade=16 → alvo = 0x0000FFFFFFFFFFFF...
	alvo *big.Int
}

// NovoProofOfWork cria um PoW para o bloco fornecido.
// O alvo é calculado deslocando 1 bit para a esquerda (256 - dificuldade) posições.
func NovoProofOfWork(b *Block) *ProofOfWork {
	// Começamos com o valor 1 e deslocamos para criar o alvo
	alvo := big.NewInt(1)
	alvo.Lsh(alvo, uint(256-dificuldade))
	// Resultado: um número com (256-dificuldade) bits — hashes abaixo disso são válidos

	return &ProofOfWork{bloco: b, alvo: alvo}
}

// prepararDados monta o slice de bytes que será hasheado,
// incluindo o nonce atual na tentativa.
func (pow *ProofOfWork) prepararDados(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.bloco.HashAnterior,
			pow.bloco.Data,
			[]byte(fmt.Sprintf("%d", pow.bloco.Timestamp)),
			[]byte(fmt.Sprintf("%d", dificuldade)),
			[]byte(fmt.Sprintf("%d", nonce)), // varia a cada tentativa
		},
		[]byte{},
	)
}

// Executar faz a mineração: incrementa o nonce até encontrar um hash válido.
// Retorna o nonce vencedor e o hash resultante.
func (pow *ProofOfWork) Executar() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Minerando bloco com dados: \"%s\"\n", pow.bloco.Data)

	for {
		// Montamos os dados com o nonce atual e calculamos o hash
		dados := pow.prepararDados(nonce)
		hash = sha256.Sum256(dados)

		// Convertemos o hash para big.Int para comparar com o alvo
		hashInt.SetBytes(hash[:])

		// Se o hash for menor que o alvo, encontramos um hash válido!
		if hashInt.Cmp(pow.alvo) == -1 {
			break
		}

		// Caso contrário, tentamos com o próximo nonce
		nonce++
	}

	fmt.Printf("Encontrado! Nonce: %d | Hash: %x\n\n", nonce, hash)
	return nonce, hash[:]
}

// Validar verifica se o hash armazenado no bloco é realmente válido.
// Usado para checar a integridade da cadeia.
func (pow *ProofOfWork) Validar() bool {
	var hashInt big.Int

	dados := pow.prepararDados(pow.bloco.Nonce)
	hash := sha256.Sum256(dados)
	hashInt.SetBytes(hash[:])

	// O hash é válido se for menor que o alvo
	return hashInt.Cmp(pow.alvo) == -1
}