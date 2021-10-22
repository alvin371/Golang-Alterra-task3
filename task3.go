package main

import (
	"fmt"
	"math"
	"math/big"
)

// Bilangan Prima
func IsPrime(value int) bool {
	if value < 2 {
		return false
	} else {
		pangkat := int(math.Sqrt(float64(value)))
		for i := 2; i <= pangkat; i++ {
			if int(value)%i == 0 {
				return false
			}
		}
	}
	return true
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

func munculSekali(a string) []int {
	charAngka := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	notSame := true
	var angka []int
	for i := 0; i < len(a); i++ {
		notSame = true
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] && i != j {
				notSame = false
				break
			}
		}

		if notSame == true {
			angka = append(angka, charAngka[string(a[i])])
		}
	}

	return angka
}

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
	fmt.Println(IsPrime(1000000007)) // true

	fmt.Println(IsPrime(1500450271)) // true

	fmt.Println(IsPrime(1000000000)) // false

	fmt.Println(IsPrime(10000000019)) // true

	fmt.Println(IsPrime(10000000033)) // true

	fmt.Println(powBig(2, 3))
	fmt.Println(powBig(7, 2))
	fmt.Println(powBig(10, 5))
	fmt.Println(powBig(17, 6))
	fmt.Println(powBig(5, 3))

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]

	fmt.Println(PairTargetWithSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairTargetWithSum([]int{2, 5, 9, 11}, 11)) // [0, 2]
	fmt.Println(PairTargetWithSum([]int{1, 3, 5, 7}, 12))  // [2, 3]
	fmt.Println(PairTargetWithSum([]int{1, 4, 6, 8}, 10))  // [1, 2]
	fmt.Println(PairTargetWithSum([]int{1, 5, 6, 7}, 6))   // [0, 1]
}
