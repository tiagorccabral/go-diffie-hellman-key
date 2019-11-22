package main

import "fmt"

// PrintMixedKeys prints the result of the secret key from both participants in the communication
func PrintMixedKeys(aliceMixedKey int64, bobMixedKey int64) {
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
	fmt.Println("\n\t" + "Obtendo chaves misturadas" + "\n")
	fmt.Print("\tMixed key da Alice = ")
	fmt.Println(aliceMixedKey)
	fmt.Print("\tMixed key do Bob = ")
	fmt.Println(bobMixedKey)
}

// PrintKeyExchange shows on terminal window the exchange of the mixedKey from both participants in the communication
func PrintKeyExchange(aliceMixedKey int64, bobMixedKey int64) {
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
	fmt.Println("\n\t" + "Troca de chaves" + "\n")
	fmt.Print("\tAlice envia ")
	fmt.Print(aliceMixedKey)
	fmt.Println(" para Bob")

	fmt.Print("\tBob envia ")
	fmt.Print(bobMixedKey)
	fmt.Println(" para Alice")
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
}

// PrintCommonSecretKey shows on terminal window the resultant secret common keys from both participants in the communication
func PrintCommonSecretKey(aliceMixedSecret int64, bobMixedSecret int64) {
	fmt.Println("\n\t" + "Chaves compartilhadas obtidas" + "\n")
	fmt.Print("\tChave de Alice obtida: ")
	fmt.Println(aliceMixedSecret)
	fmt.Print("\tChave de Bob obtida: ")
	fmt.Println(bobMixedSecret)
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
}
