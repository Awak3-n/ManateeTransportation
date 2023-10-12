package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
(greedy approach)
generates every possible combination of manatees being held in the tubs by use of a binary number with range of 0 to 2^n (n == number of manatees total).
a seprate function uses the number as a set of instructions. In the binary number,
0 represents a manatee sent into the left tub, and 1 repesents a manatee sent into the right tub.
it then compares each of the maximums from all of the given binary combinations and returns the highest maximum
*/
func getMaxManatees(manateeList *list.List, tubLength int) int {
	len := manateeList.Len()                    // amount of manatees stored in list
	possibleCombos := math.Pow(2, float64(len)) // maximum amount of possible combinations with inputting manatees into the left or right tub (2^n)
	var localMax int
	var globalMax int

	for i := 0; i < int(possibleCombos); i++ {
		currentCombo := PadLeft(strconv.FormatInt(int64(i), 2), len) // currentCombo is a padded binary number of type string, aka the instructions that getLocalMaxManatees will follow
		localMax = getLocalMaxManatees(manateeList, tubLength, currentCombo)

		if localMax > globalMax {
			globalMax = localMax
		}
	}

	return globalMax
}

// pads string with 0 from the left side
// useful for adding in the 0's that FormatInt throws away when producing a binary number
func PadLeft(binaryStr string, length int) string {

	for len(binaryStr) < length {
		binaryStr = "0" + binaryStr
	}

	return binaryStr
}

/*
getLocalMaxManatees calculates the amount of manatees that can fit based of a binary number that getMaxManatees supplies it with
it then returns the maximum amount of manatees it can fit into the tube based on the provided instructions
*/
func getLocalMaxManatees(manateeQueue *list.List, tublength int, strInstructions string) int {
	var localMax, leftTub, rightTub int

	// base case #1: no manatees :(
	if manateeQueue.Front() == nil {
		return localMax
	}

	// i know im only pushing ints into the list, so I be casting
	currentManatee := manateeQueue.Front()
	currentIntManatee := currentManatee.Value.(int)

	for i := 0; i < len(strInstructions); i++ {

		// puts currentIntManatee into lefttub when the current instruction is a 0
		if string(strInstructions[i]) == "0" && currentIntManatee+leftTub <= tublength {
			leftTub += currentIntManatee
			localMax++
		}

		// puts currentIntManatee into righttub when the current instruction is a 1
		if string(strInstructions[i]) == "1" && currentIntManatee+rightTub <= tublength {
			rightTub += currentIntManatee
			localMax++
		}

		// base case #2: no more manatees to add to the tubs
		if currentManatee.Next() != nil {
			currentManatee = currentManatee.Next()
			currentIntManatee = currentManatee.Value.(int)
		} else {
			return localMax
		}

		//base case #3: tubs are full
		if currentIntManatee+leftTub > tublength && currentIntManatee+rightTub > tublength {
			return localMax
		}
	}

	// returns error if the base cases do not return
	return -1
}

//func addPort() {

//}

//func addStarboard {

//}

func main() {
	var tubLength int          // how big both of the tubs are
	manateeQueue := list.New() // list that holds all of the given manatees lengths
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	tubLength, err := strconv.Atoi(scanner.Text())

	if err == nil {
		if tubLength >= 1 && tubLength <= 100 {
			tubLength *= 100
			fmt.Printf("tublength: %d\n", tubLength)
		}
	}

	for true {
		scanner.Scan()
		manatee, err := strconv.Atoi(scanner.Text())

		if err == nil {
			if manatee != 0 {
				if manatee >= 100 && manatee <= 3000 {
					manateeQueue.PushBack(manatee)
					fmt.Printf("pushing manatee: %d\n", manatee)
				}
			} else {
				break
			}
		}
	}

	max := getMaxManatees(manateeQueue, tubLength) // maximum amount of possible combinations of manatees

	fmt.Printf("max amount of manatees: %d\n", max)
	fmt.Printf("tublength: %d\n", tubLength)

	for i := manateeQueue.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
