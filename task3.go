package main

import (
	"fmt"
	"math"
	"math/big"
)

// Bilangan Prima
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// Fast Exponent logarithmic
func powBig(a, n int) *big.Int {
	tmp := big.NewInt(int64(a))
	res := big.NewInt(1)
	for n > 0 {
		temp := new(big.Int)
		if n%2 == 1 {
			temp.Mul(res, tmp)
			res = temp
		}
		temp = new(big.Int)
		temp.Mul(tmp, tmp)
		tmp = temp
		n /= 2
	}
	return res
}

// Array Merge
func ArrayMerge(a1, a2 []string) []string {

	var newA []string
	noSame := true
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			if a1[i] == a2[j] {
				noSame = false
			}
		}
		if noSame == true {
			newA = append(newA, a1[i])
		}
	}

	merge := append(newA, a2[:]...)
	return merge
}

// Muncul sekali

// Pair Target With sum
func PairTargetWithSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}

func main() {
	var prime int64
	fmt.Printf("Masukkan Angka yang ingin dicheck: ")
	fmt.Scan(&prime)
	IsPrime(int(prime))

	fmt.Println(ArrayMerge([]string{“sergei”, “jin”}, []string{“jin”, “steve”, “bryan”}))
	fmt.Println(PairTargetWithSum([]int{1, 2, 3, 4, 6}, 6))
}
