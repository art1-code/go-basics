package main

import "fmt"

func main()  {
	
	var nomes = [5]string{"arthur","carol","tito","magna","givaldo"}

	for i := 0; i < len(nomes); i++ {
		fmt.Printf("O nome atual é: %s\n", nomes[i])
	}
	
	for i, nome := range nomes{
		fmt.Printf("Indice Atual : %d. O nome atual é: %s\n", i, nome)
	}

	for _ , nome := range nomes{
		fmt.Printf("O nome atual é: %s\n", nome)
	}


	// usar for como loop eterno for {...}


	/* for {
		fmt.Println("Executando...")
	} */
}