package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 7")

	lines, _ := utils.LoadData("7.txt")
	equations := createEquationMap(lines)

	// Part 1
	start := time.Now()
	partOneSol := partOne(equations)
	duration := time.Since(start)
	fmt.Printf("Part 1: %s (Execution time: %d ms)\n", partOneSol, duration.Milliseconds())

	// Part 2
	start = time.Now()
	partTwoSol := partTwo(equations)
	duration = time.Since(start)
	fmt.Printf("Part 2: %s (Execution time: %d ms)\n", partTwoSol, duration.Milliseconds())

}

func createEquationMap(lines []string) map[*big.Int][]*big.Int {
	equations := make(map[*big.Int][]*big.Int)
	for _, line := range lines {
		lhs, _ := strconv.Atoi(strings.Split(line, ":")[0])
		rhs := strings.Split(line, ": ")[1]
		lhsBig := big.NewInt(int64(lhs))
		for _, operandStr := range strings.Split(rhs, " ") {
			operand, _ := strconv.Atoi(operandStr)
			operandBig := big.NewInt(int64(operand))
			equations[lhsBig] = append(equations[lhsBig], operandBig)
		}
	}
	return equations
}

func partOne(equations map[*big.Int][]*big.Int) *big.Int {
	sum := big.NewInt(0)
	operators := []string{"*", "+"}
	for target, operands := range equations {
		reachesTarget := resolves(operands, target, operators)
		if reachesTarget {
			sum.Add(sum, target)
		}
	}
	return sum
}

func partTwo(equations map[*big.Int][]*big.Int) *big.Int {
	sum := big.NewInt(0)
	operators := []string{"*", "+", "||"}
	for target, operands := range equations {
		reachesTarget := resolves(operands, target, operators)
		if reachesTarget {
			sum.Add(sum, target)
		}
	}
	return sum
}

func resolves(operands []*big.Int, target *big.Int, operators []string) bool {
	if len(operands) < 2 {
		return false
	}

	totals := []*big.Int{}
	totals = append(totals, operands[0])

	for i := 1; i < len(operands); i++ {
		var newTotals []*big.Int
		for _, total := range totals {
			for _, op := range operators {
				result := new(big.Int)
				switch op {
				case "*":
					result.Mul(total, operands[i])
				case "+":
					result.Add(total, operands[i])
				case "||":
					concatStr := total.String() + operands[i].String()
					concatVal, ok := new(big.Int).SetString(concatStr, 10)
					if !ok {
						fmt.Println("Error converting concatenated string to big.Int")
						continue
					} else {
						result.Set(concatVal)
					}
				}
				if result.Cmp(target) <= 0 {
					newTotals = append(newTotals, new(big.Int).Set(result))
				}

				if result.Cmp(target) == 0 && i == len(operands)-1 {
					return true
				}
			}
		}
		if len(newTotals) == 0 {
			break
		}
		totals = newTotals
	}

	return false
}
