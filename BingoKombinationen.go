package main

import "fmt"

func isBingo(n int) bool {
	if isLine(n) || isRow(n) || isDiagonal(n) {return true}
	return false
}

func isDiagonal(n int) bool {
	mask1 := 0b1000010000100001
	mask2 := 0b0001001001001000
	if countOnes(n & mask1)==4 {return true}
	if countOnes(n & mask2)==4 {return true}

	return false
}

func isRow(n int) bool {
	mask := 0b1000100010001
	for n!=0 {
		if countOnes(n & mask)==4 {return true}
		n >>= 1
	}
	return false
}

func isLine(n int) bool {
	mask := 0b1111
	for n!=0 {
		if n & mask == 0b1111 {return true}
		n >>= 4
	}
	return false
}

func countOnes(n int) (i int){
	for n != 0{
		i += n&1
		n>>=1
	}
	return
}

func main() {
	for numOfBits := 0 ; numOfBits <=16 ; numOfBits++{
		sum,nsum := countBingos(numOfBits)
		fmt.Println("Number of Combinations with", numOfBits, "crosses to not win Bingo is", nsum, "to win", sum)
	}
}

func countBingos(numOfBits int) (bingos, nonbingos int) {
	for i := 0; i < 0b1<<16; i++ {
		if countOnes(i) == numOfBits{
			if isBingo(i) {
				bingos++
			} else {
				nonbingos++
			}
		}
	}
	return
}
