package main

import (
	"fmt"
)

func main() {
	const dimen int = 4

	matrizA := [dimen][dimen]int{{1, 1}, {1, 1}}
	matrizB := [dimen][dimen]int{{1, 1}, {1, 1}}

	fmt.Println("Resultado:")
	for i := 0; i < dimen; i++ {
		for j := 0; j < dimen; j++ {
			fmt.Print(matrizA[i][j]+matrizB[i][j], "\t")
		}
		fmt.Println()
	}
}
