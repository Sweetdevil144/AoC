package Arrays

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

func Recon() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    n, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal("Failed to read the number of soldiers:", err)
    }

    scanner.Scan()
    heightsInput := scanner.Text()
    heightsStr := strings.Fields(heightsInput)
    if len(heightsStr) != n {
        log.Fatalf("Expected %d heights but got %d", n, len(heightsStr))
    }

    heights := make([]int, n)
    for i, h := range heightsStr {
        heights[i], err = strconv.Atoi(h)
        if err != nil {
            log.Fatalf("Failed to parse height: %v", err)
        }
    }

    minDiff := math.MaxInt32
    var soldier1, soldier2 int
    for i := 0; i < n; i++ {
        next := (i + 1) % n
        diff := abs(heights[i] - heights[next])
        if diff < minDiff {
            minDiff = diff
            soldier1, soldier2 = i, next
        }
    }

    fmt.Printf("%d %d\n", soldier1+1, soldier2+1)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}