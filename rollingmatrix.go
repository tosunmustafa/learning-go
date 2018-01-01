//How to compile
//  SET GOOS=windows
//	SET GOARCH=amd64
//	go build -a -o rollingmatrix_amd64.exe rollingmatrix.go
//
package main

import(
	"fmt"
	"os"
	"strconv"
)

// const N = 1  --> Had to compile every time! pfff

func printMatrix(N int) {
	for i := 1 ; i <= N ; i += 1 { //row loop

		for j, k := i, 1 ; k <= N; j,k = j+1, k+1 { //column loop
			if j > N {
				j = 1
			}
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func main(){

	strN := os.Args[1:][0] //take commandline argument, string form of N (size of matrix)

	if N, err := strconv.Atoi(strN); err == nil { //strconv is multiple returning function, use it like this
		printMatrix(N)
	}
}
