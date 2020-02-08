//package main Informa que será este será o pacote principal e a execução do código começara aqui
package main

import "fmt"

//Todo programa principal tem  função principal, neste caso func main()
func main() {
	fmt.Println("Olá mundo! Fabio!") //Print com P maiúscula. Isso demonstra que não veio do seu pacote, veio de um pacote externo
	//para compilar
	// go build hello.go -> isso só compila. Para rodar teria que fazer ./hello.exe, chamar a sua execução
	//go run hello.go -> este comando compilar e depis executa a aplicação.
	// Ponto e vírgula é opcional em Go, tanto é verdade que o próprio VSCod remove ele. É uma convenção do Go que o VSC segue.
	//As chaves em Go também seguem um padrão, como está neste código.

}
