package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func getMinimumPenalty(quantity []int, m int) int64 {
    // Sort ascending agar index 0 selalu berisi stok paling kecil
    sort.Ints(quantity)

    var penalty int64 = 0

    for m > 0 && len(quantity) > 0 {
        // Ambil terkecil (index 0)
        smallest := quantity[0]

        // Tambahkan penalty
        penalty += int64(smallest)

        // Kurangi stok terkecil
        smallest--

        if smallest == 0 {
            // Hapus elemen 0
            quantity = quantity[1:]
        } else {
            // Update lalu sort lagi
            quantity[0] = smallest
            sort.Ints(quantity)
        }

        m--
    }

    return penalty
}

func main() {
    in := bufio.NewReader(os.Stdin)

    var n, m int
    fmt.Fscan(in, &n, &m)

    quantity := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &quantity[i])
    }

    fmt.Println(getMinimumPenalty(quantity, m))
}
