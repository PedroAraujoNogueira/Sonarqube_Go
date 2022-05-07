package main

import "testing"

// Para rodar um teste em Go basta irmos no diretório do nosso pacote e digitar no terminal
// o comando "go mod init firstappgo" e em seguida "go test". Com isso, Go encontrará todos 
// os arquivos com o sufixo "_test.go" e todas as funções iniciadas com Test e as executará.

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}

}

func TestSub(t *testing.T) {

	result := sub(3, 3)

	if result != 0 {
		t.Error("The result must be 0")
	}

}

func TestTimes(t *testing.T) {

	result := times(3, 3)

	if result != 9 {
		t.Error("The result must be 9")
	}

}

func TestSumX(t *testing.T) {

	result := sumX(2, 3)

	if result != 7 {
		t.Error("The result must be 7")
	}

}