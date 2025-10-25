package main

import "fmt"


func main()  {
	var age int;

	for {
		fmt.Println("Digite uma idade valida (+18):")
		_, err := fmt.Scanf("%d\n",&age)

		if err != nil {
			fmt.Println("Entrada inválida, digite novamente")
			var discard string
      fmt.Scanln(&discard) 
			continue
		}
		if age >= 18 {
			fmt.Println("Idade Valida, parabens vc foi preso!")
		} else {
			fmt.Println("Idade Valida, parabens vc foi solto, até a proxima!")
		}
		break
	}
}