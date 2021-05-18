package main

import (
	"fmt"
	"time"
)

func isBingo(n int, gridsize int) bool {
	if isLine(n, gridsize) || isRow(n, gridsize) || isDiagonal(n, gridsize) {
		return true
	}
	return false
}

func isDiagonal(n, gridsize int) bool {
	mask1 := 0b1
	for i := 1; i < gridsize; i++ {
		mask1 <<= gridsize + 1
		mask1 += 0b1
	}
	if countOnes(n&mask1) == gridsize {
		return true
	}

	mask2 := 0b1
	for i := 1; i < gridsize; i++ {
		mask2 <<= gridsize - 1
		mask2 += 0b1
	}
	mask2 <<= gridsize - 1
	if countOnes(n&mask2) == gridsize {
		return true
	}

	return false
}

func isRow(n, gridsize int) bool {
	mask := 0b1
	for i := 1; i < gridsize; i++ {
		mask <<= gridsize
		mask += 0b1
	}
	for i := 0; i < gridsize; i++ {
		if countOnes(n&mask) == gridsize {
			return true
		}
		n >>= 1
	}
	return false
}

func isLine(n int, gridsize int) bool {
	mask := 0b0
	for i := 0; i < gridsize; i++ {
		mask += 0b1 << i
	}
	for n != 0 {
		if n&mask == mask {
			return true
		}
		n >>= gridsize
	}
	return false
}

func countOnes(n int) (i int) {
	for n != 0 {
		i += n & 1
		n >>= 1
	}
	return
}

func countBingos(numOfBits, gridsize int) (bingos, nonbingos int) {
	for i := 0; i < 0b1<<(gridsize*gridsize); i++ {
		if countOnes(i) == numOfBits {
			if isBingo(i, gridsize) {
				bingos++
			} else {
				nonbingos++
			}
		}
	}
	return
}

func main() {
	t := time.Now()
	fmt.Println("Number of Combinations to win Bingo on different Gridsizes")
	fmt.Println("")
	gridsize := 4
	for numOfBits := 0; numOfBits <= (gridsize * gridsize); numOfBits++ {
		sum, nsum := countBingos(numOfBits, gridsize)
		fmt.Printf("%d | %d | %d | %d \n", gridsize, numOfBits, sum, nsum)
	}
	fmt.Println(time.Since(t))
}
