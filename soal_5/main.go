package main

import (
    "fmt"
    "strings"
)

var memo map[string]int64

func normalize(arr []int) string {
    // serialize array for memo key
    parts := make([]string, len(arr))
    for i, v := range arr {
        parts[i] = fmt.Sprintf("%d", v)
    }
    return strings.Join(parts, ",")
}

func getDemolitionScore(arr []int, k int) int64 {
    memo = make(map[string]int64)
    return dfs(arr, k)
}

func dfs(arr []int, k int) int64 {
    if k == 0 || len(arr) == 0 {
        return 0
    }

    key := fmt.Sprintf("%s|%d", normalize(arr), k)
    if v, ok := memo[key]; ok {
        return v
    }

    var best int64 = 0

    n := len(arr)

    for i := 0; i < n; i++ {
        destroyed := arr[i]

        left := arr[:i]
        right := arr[i+1:]

        var next []int

        if len(left) > len(right) {
            next = append([]int{}, left...)
        } else if len(right) > len(left) {
            next = append([]int{}, right...)
        } else {
            // equal size → discard left
            next = append([]int{}, right...)
        }

        // weaken surviving partition
        for j := range next {
            next[j] -= destroyed
            if next[j] < 0 {
                next[j] = 0
            }
        }

        // recursion
        score := int64(destroyed) + dfs(next, k-1)

        if score > best {
            best = score
        }
    }

    memo[key] = best
    return best
}

func main() {
    arr := []int{10, 2, 8, 5}
    k := 2
    fmt.Println(getDemolitionScore(arr, k)) // expected 10
}
