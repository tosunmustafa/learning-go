// A Dynamic Programming based solution for 0-1 Knapsack problem
// See: www.cdn.geeksforgeeks.org/knapsack-problem

package main

import(
	"fmt"
)

func max(a int, b int) int {
	switch {
		case a-b >= 0: return a
		case a-b < 0 : return b
		default      : return a
	}
}


func knapsack(pK *[][]int, W int, pWeightArray *[]int, pValueArray *[]int, N int) int{

	var K, weightArray, valueArray = *pK, *pWeightArray, *pValueArray //dereference pointers to arrays

	for i:= 0; i <= N; i++ {
		for j:=0; j <= W ; j++ {
			if i==0 || j ==0 {
				K[i][j] = 0
			}else if weightArray[i-1] <= j {
				K[i][j] = max(valueArray[i-1] + K[i-1][j-weightArray[i-1]],  K[i-1][j])
			}else {
				K[i][j] = K[i-1][j]
			}
		}
	}
	return K[N][W]
}


func main(){
    valueArray  := []int {60, 100, 120}
    weightArray := []int {10, 20, 30  }
    const W,N int = 50, 3

    //initialize 2 dimensional array
    K := make([][]int, N+1)
    for i:= range K {
        K[i] = make([]int, W+1)
    }

    // Arrays maybe too big, pass their adresses as arguments
    // Go is call-by-value!
    fmt.Println(knapsack(&K, W, &weightArray, &valueArray, N)) //Should print 220 :)
}