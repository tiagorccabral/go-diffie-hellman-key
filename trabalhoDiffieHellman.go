// Autor: Tiago Rodrigues da Cunha Cabral
// Matrícula: 15/0150296
// Disciplina: Segurança Computacional - 117927 - Turma A
// Algoritmo: Troca de Chaves Diffie-Hellman

package main

import (
	"fmt"
	"math"
)

var (
	commonPrime      int64   = 23
	commonSquareRoot float64 = 5
)

func modularExponent(x int64, y int64, modulos int64) int64 {
	if y == 0 {
		return 1
	}
	if y%2 == 1 {
		return (x * modularExponent(x, y-1, modulos)) % modulos
	}
	t := modularExponent(x, y/2, modulos)
	return (t * t) % modulos
}

func mixKeys(privateKey float64) int64 {
	mixedKey := math.Pow(commonSquareRoot, privateKey)
	return int64(int64(mixedKey) % commonPrime)
}

func mixSecretKeys(receivedMixKey int64, ownSecretKey float64) int64 {
	mixedKey := math.Pow(float64(receivedMixKey), ownSecretKey)
	return int64(int64(mixedKey) % commonPrime)
}

func main() {

	// publicKey := "blue"
	var aliceSecretKey, bobSecretKey float64 = 6, 15

	aliceMixedKey := mixKeys(aliceSecretKey)
	bobMixedKey := mixKeys(bobSecretKey)

	fmt.Println("Mixed Alice = ")
	fmt.Println(aliceMixedKey)
	fmt.Println("Mixed Bob = ")
	fmt.Println(bobMixedKey)

	fmt.Println("\n" + "-------------------------------" + "\n")

	fmt.Print("Alice envia ")
	fmt.Print(aliceMixedKey)
	fmt.Println()

	fmt.Print("Bob envia ")
	fmt.Print(bobMixedKey)
	fmt.Println()

	aliceMixedSecret := mixSecretKeys(bobMixedKey, aliceSecretKey)
	bobMixedSecret := mixSecretKeys(aliceMixedKey, bobSecretKey)

	fmt.Println(aliceMixedSecret)
	fmt.Println(bobMixedSecret)

}
