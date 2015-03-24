package main 

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
    "math/big"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Converts the string 's' to an int.
func toInt(s string) int {
    x, err := strconv.Atoi(s)
    check(err)
    return x
}

// Cheaty way of checking if something is a prime.
// From the golang docs:
// ProbablyPrime performs n Miller-Rabin tests to check whether x is prime. If it returns true, x is prime with
// probability 1 - 1/4^n. If it returns false, x is not prime.
func isPrime(x int) bool {
    if x <= 1 {
        return false
    }
    y := big.NewInt(int64(x))
    return y.ProbablyPrime(1)
}

// Checks to see if a number is a square root.
func isSquareRoot(x int) bool {
    i := math.Sqrt(float64(x))
    if i == float64(int64(i)) {
        return true
    }
    return false
}

// Checks to see if a number is a cube root.
func isCubeRoot(x int) bool {
    i := math.Cbrt(float64(x))
    if i == float64(int64(i)) {
        return true
    }
    return false
}

// Checks to see if a number is a quad root.
func isQuadRoot(x int) bool {
    i := math.Pow(float64(x), 0.25)
    if i == float64(int64(i)) {
        return true
    }
    return false
}

// Checks to see if a number is a multiple of the sum of it's digits.
func isMultipleSum(x int) bool {
    if x < 10 {
        return true
    } else {
        y := x
        sum := 0
        for y != 0 {
            sum += y % 10
            y = y / 10
        }

        if x % sum == 0 {
            return true
        }
    }
    return false
}

// Checks to see if a number is a multiple of the product of its digits.
func isMultipleMultiple(x int) bool {
    if x < 10 {
        return true
    } else {
        y := x
        product := 1
        for y != 0 {
            product *= y % 10
            y = y / 10
        }

        if product != 0 && x % product == 0 {
            return true
        }
    }
    return false
}

// Checks to see if the number is a factor of any other number in the list.
func isFactor(x int, list []int) bool {
    for _, y := range list {
        if x != y && y % x == 0 {
            return true
        }
    }
    return false
}

// Checks to see if the number is a multiple of any other number in the list.
func isMultiple(x int, list []int) bool {
    for _, y := range list {
        if x != y && x % y == 0 {
            return true
        }
    }
    return false
}

// Helper method for all of the power methods.
func inList(x int, list[]int, power int) bool {
    for _, i := range list {
        y := math.Pow(float64(i), float64(power))
        if x != i && x == int(y) {
            return true
        }
    }
    return false
}

// Is the number a square of another number in the list?
func otherSquare(x int, list []int) bool {
    return inList(x, list, 2)
}

// Is the number a cube of another number in the list?
func otherCube(x int, list[]int) bool {
    return inList(x, list, 3)
}

// Is the number a quad of another number in the list?
func otherQuad(x int, list[]int) bool {
    return inList(x, list, 4)
}

// Is the number a multiple of the sum of the digits of another number in the list (wut?).
func otherDigitSum(x int, list []int) bool {
    for _, i := range list {
        if x != i {
            sum := 0
            num := i
            for num > 0 {
                sum += num % 10
                num = num / 10
            }
            if x % sum == 0{
                return true
            }
        }
    }
    return false
}

// Is the number a mulitple of the product of another number's digits in the list.
func otherDigitMultiple(x int, list []int) bool {
    for _, i := range list {
        if x != i {
            product := 1
            num := i
            for num > 0 {
                product *= num % 10
                num = num / 10
            }
            if product != 0 && x % product == 0 {
                return true
            }
        }
    }
    return false
}

// Main function.
func main() {
    // Open file.
    file, err := os.Open("../one.in")
    check(err)

    // Don't forget to close the file
    defer file.Close()

    // Gets the number of datasets.
    scanner := bufio.NewScanner(file)

    scanner.Scan()
    dataSets := toInt(scanner.Text())

    // Loop over all of the numbers in a dataset.
    for i := 0; i < dataSets; i++ {
        fmt.Printf("DATA SET #%d\n", i+1)
        var mostInteresting []int

        scanner.Scan()
        quantity := toInt(scanner.Text())

        collection := make([]int, quantity)

        // Put the numbers in a slice.
        for j := 0; j < quantity; j++ {
            scanner.Scan()
            num := toInt(scanner.Text())

            collection[j] = num
        }

        // max == interestingness
        max := 0
        // For ever number in the list, see what interesting qualities it has.
        for _, x := range collection {
            sum := 0
            if otherDigitMultiple(x, collection) {
                sum++
            }
            if otherDigitSum(x, collection) {
                sum++
            }
            if otherQuad(x, collection) {
                sum++
            }
            if otherCube(x, collection) {
                sum++
            }
            if otherSquare(x, collection) {
                sum++
            }
            if isMultiple(x, collection) {
                sum++
            }
            if isFactor(x, collection) {
                sum++
            }
            if isMultipleMultiple(x) {
                sum++
            }
            if isMultipleSum(x) {
                sum++
            }
            if isQuadRoot(x) {
                sum++
            }
            if isCubeRoot(x) {
                sum++
            }   
            if isSquareRoot(x) {
                sum++
            }
            if isPrime(x) {
                sum++
            }
            // If the number is as interesting as the most interesting number so far,
            // add it to the list.
            if sum == max {
                mostInteresting = append(mostInteresting, x)
            }
            // If it is more interesting, set it as the most interesting and forget the 
            // second most interesting number, because who cares...
            if sum > max {
                max = sum
                mostInteresting = nil
                mostInteresting = append(mostInteresting, x)
            }
        }
        // Sorts the list "they must be printed in ascending order"
        sort.Sort(sort.IntSlice(mostInteresting))
        // Print the interesting numbers.
        for _, best := range mostInteresting {
            fmt.Println(best)
        }
    }
}
