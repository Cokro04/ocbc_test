// go run ascii_demolition.go
package main

import (
	"fmt"
	"strings"
	"time"
)

// Memo entry storing best score and the corresponding choice path
type memoEntry struct {
	score int64
	path  []int
}

var memo map[string]memoEntry

// normalize array into string for memo key
func normalize(arr []int) string {
	parts := make([]string, len(arr))
	for i, v := range arr {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, ",")
}

// dfs returns best score and the sequence of indices chosen (relative to current array)
// Note: indices are with respect to the array at that step
func dfs(arr []int, k int) (int64, []int) {
	if k == 0 || len(arr) == 0 {
		return 0, []int{}
	}
	key := fmt.Sprintf("%s|%d", normalize(arr), k)
	if e, ok := memo[key]; ok {
		return e.score, append([]int(nil), e.path...)
	}

	var bestScore int64 = -1
	var bestPath []int

	n := len(arr)
	for i := 0; i < n; i++ {
		destroyed := arr[i]
		left := append([]int(nil), arr[:i]...)
		right := append([]int(nil), arr[i+1:]...)

		var next []int
		if len(left) > len(right) {
			next = left
		} else if len(right) > len(left) {
			next = right
		} else {
			// equal size: discard left -> keep right
			next = right
		}

		// weaken surviving partition
		for j := range next {
			next[j] -= destroyed
			if next[j] < 0 {
				next[j] = 0
			}
		}

		nextScore, nextPath := dfs(next, k-1)
		total := int64(destroyed) + nextScore
		if total > bestScore {
			bestScore = total
			// path: pick i now, then append subsequent indices (they are relative to subsequent arrays)
			bestPath = append([]int{i}, nextPath...)
		}
	}

	if bestScore < 0 {
		bestScore = 0
		bestPath = []int{}
	}
	memo[key] = memoEntry{score: bestScore, path: append([]int(nil), bestPath...)}
	return bestScore, append([]int(nil), bestPath...)
}

// ASCII drawing helpers

// clear screen (ANSI)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// drawBuildings draws vertical bars based on heights slice.
// highlightIdx = index to mark as demolished (use -1 for none)
func drawBuildings(heights []int, highlightIdx int) {
	// find max height
	maxH := 0
	for _, h := range heights {
		if h > maxH {
			maxH = h
		}
	}
	if maxH == 0 {
		fmt.Println("(all buildings height 0)")
		return
	}

	// draw rows from maxH down to 1
	for level := maxH; level >= 1; level-- {
		for i, h := range heights {
			if h >= level {
				if i == highlightIdx {
					// different glyph for demolished building
					fmt.Print(" ██ ")
				} else {
					fmt.Print(" █  ")
				}
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
	// base line with indices and values
	for range heights {
		fmt.Print("─── ")
	}
	fmt.Println()
	for _, h := range heights {
		fmt.Printf("%3d ", h)
	}
	fmt.Println()
	// indices row
	for i := range heights {
		fmt.Printf(" %d  ", i)
	}
	fmt.Println("\n")
}

// simulateAndAnimate runs through the sequence (indices) and animates each step
func simulateAndAnimate(arr []int, sequence []int) {
	score := int64(0)
	current := append([]int(nil), arr...)

	for step, pick := range sequence {
		clearScreen()
		fmt.Printf("STEP %d — before explosion (score=%d)\n\n", step+1, score)
		drawBuildings(current, pick)
		fmt.Printf("\nDemolish index %d (value %d)\n", pick, current[pick])
		time.Sleep(1200 * time.Millisecond)

		// perform demolition
		destroy := current[pick]
		score += int64(destroy)

		// partition
		left := append([]int(nil), current[:pick]...)
		right := append([]int(nil), current[pick+1:]...)
		var remaining []int
		if len(left) > len(right) {
			remaining = left
			fmt.Println("\nKeep LEFT partition, discard RIGHT")
		} else if len(right) > len(left) {
			remaining = right
			fmt.Println("\nKeep RIGHT partition, discard LEFT")
		} else {
			remaining = right
			fmt.Println("\nPartitions equal size -> discard LEFT, keep RIGHT")
		}
		time.Sleep(900 * time.Millisecond)

		// weaken remaining
		for i := range remaining {
			remaining[i] -= destroy
			if remaining[i] < 0 {
				remaining[i] = 0
			}
		}

		clearScreen()
		fmt.Printf("STEP %d — after explosion (score=%d)\n\n", step+1, score)
		drawBuildings(remaining, -1)
		fmt.Println("Press Ctrl+C to quit, or wait for next step...")
		time.Sleep(1200 * time.Millisecond)

		// update current
		current = remaining
	}

	// final frame
	clearScreen()
	fmt.Printf("FINAL RESULT — TOTAL SCORE = %d\n\n", score)
	drawBuildings(current, -1)
	fmt.Println("\nAnimation complete.")
}

// entry point with example
func main() {
	// Example: demolition problem arr = [10,2,8,5], k=2
	arr := []int{10, 2, 8, 5}
	k := 2

	fmt.Println("Computing optimal demolition sequence, please wait...")
	memo = make(map[string]memoEntry)
	bestScore, bestPath := dfs(arr, k)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Best score = %d\nSequence of picks (indices per step): %v\n", bestScore, bestPath)
	fmt.Println("Animating now...")
	time.Sleep(1200 * time.Millisecond)

	simulateAndAnimate(arr, bestPath)
}
