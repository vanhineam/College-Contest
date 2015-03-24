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

func toInt(s string) int {
    x, err := strconv.Atoi(s)
    check(err)
    return x
}

func isPrime(x int) bool {
    if x <= 1 {
        return false
    }
    y := big.NewInt(int64(x))
    return y.ProbablyPrime(1)

}

func isSquareRoot(x int) bool {
    i := math.Sqrt(float64(x))
    if i == float64(int64(i)) {
        return true
    }
    return false
}

func isCubeRoot(x int) bool {
    i := math.Cbrt(float64(x))
    if i == float64(int64(i)) {
        return true
    }
    return false
}

func isQuadRoot(x int) bool {
    i := math.Pow(float64(x), 0.25)
    if i == float64(int64(i)) {
        return true
    }
    return false
}

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

func isFactor(x int, list []int) bool {
    for _, y := range list {
        if x != y && y % x == 0 {
            return true
        }
    }
    return false
}

func isMultiple(x int, list []int) bool {
    for _, y := range list {
        if x != y && x % y == 0 {
            return true
        }
    }
    return false
}

func inList(x int, list[]int, power int) bool {
    for _, i := range list {
        y := math.Pow(float64(i), float64(power))
        if x != i && x == int(y) {
            return true
        }
    }
    return false
}

func otherSquare(x int, list []int) bool {
    return inList(x, list, 2)
}

func otherCube(x int, list[]int) bool {
    return inList(x, list, 3)
}

func otherQuad(x int, list[]int) bool {
    return inList(x, list, 4)
}

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

func main() {
    file, err := os.Open("../one.in")
    check(err)

    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    dataSets := toInt(scanner.Text())

    for i := 0; i < dataSets; i++ {
        fmt.Printf("DATA SET #%d\n", i+1)
        var mostInteresting []int

        scanner.Scan()
        quantity := toInt(scanner.Text())

        collection := make([]int, quantity)

        for j := 0; j < quantity; j++ {
            scanner.Scan()
            num := toInt(scanner.Text())

            collection[j] = num
        }

        max := 0
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
            if sum == max {
                mostInteresting = append(mostInteresting, x)
            }
            if sum > max {
                max = sum
                mostInteresting = nil
                mostInteresting = append(mostInteresting, x)
            }
        }
        sort.Sort(sort.IntSlice(mostInteresting))
        for _, best := range mostInteresting {
            fmt.Println(best)
        }
    }
}
