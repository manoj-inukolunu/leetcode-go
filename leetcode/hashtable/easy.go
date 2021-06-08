package main

import (
	"fmt"
	"strings"
)

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for idx, val := range nums {
		diff := target - val
		if value, ok := hashmap[diff]; ok {
			return []int{idx, value}
		} else {
			hashmap[val] = idx
		}
	}
	return []int{-1, -1}
}

func isAlienSorted(words []string, order string) bool {
	hashmap := make(map[rune]int)
	for idx, char := range order {
		hashmap[char] = idx
	}
	for i := 0; i+1 < len(words); i++ {
		first := words[i]
		second := words[i+1]
		if strings.HasPrefix(first, second) && len(first) != len(second) {
			return false
		}
		for idx, char := range first {
			if idx < len(second) && rune(second[idx]) != char && hashmap[char] > hashmap[rune(second[idx])] {
				return false
			} else if idx < len(second) && char < rune(second[idx]) {
				break
			}
		}
	}
	return true
}

func countPrimes(n int) int {
	sieve := make([]bool, n)
	for idx := range sieve {
		sieve[idx] = true
	}
	for i := 2; i*i < n; i++ {
		if !sieve[i] {
			for j := i * i; j < n; j += i {
				sieve[j] = false
			}
		}
	}
	count := 0
	for _, val := range sieve {
		if val {
			count++
		}
	}
	return count
}

func sieve(n int) []bool {
	primes := make([]bool, n)
	for idx := range primes {
		primes[idx] = true
	}
	for i := 2; i*i < n; i++ {
		if !primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	return primes
}

func longestWord(words []string) string {
	set := make(map[string]bool)

	for _, str := range words {
		set[str] = true
	}
	var ans string
	length := 0
	for _, str := range words {
		found := true
		for i := 0; i <= len(str); i++ {
			substr := str[:i]
			if substr != "" && !set[substr] {
				found = false
				break
			}
		}
		if found && len(str) >= length {
			if len(str) == length {
				if str < ans {
					ans = str
				}
			} else {
				ans = str
			}
			length = len(str)
		}
	}
	return ans
}

func isIsomorphic(s string, t string) bool {
	mapping := make(map[rune]rune)
	mapped := make(map[rune]bool)
	for idx, char := range s {
		if _, exists := mapping[char]; !exists {
			if mapped[rune(t[idx])] {
				return false
			}
			mapping[char] = rune(t[idx])
			mapped[rune(t[idx])] = true
		} else if mapping[char] != rune(t[idx]) {
			return false
		}
	}
	return true
}

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	hashmap := make(map[string]map[string]bool)
	for _, arr := range similarPairs {
		if _, ok := hashmap[arr[0]]; !ok {
			hashmap[arr[0]] = make(map[string]bool)
		}
		hashmap[arr[0]][arr[1]] = true
		if _, ok := hashmap[arr[1]]; !ok {
			hashmap[arr[1]] = make(map[string]bool)
		}
		hashmap[arr[1]][arr[0]] = true
	}
	for idx, word := range sentence1 {
		if word == sentence2[idx] {
			continue
		}
		if _, ok := hashmap[word][sentence2[idx]]; !ok {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	collector := []int{}
	dfs(root, &collector)
	return collector
}

func dfs(root *TreeNode, collector *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, collector)
	*collector = append(*collector, root.Val)
	dfs(root.Right, collector)
}

func main() {
	arr := []int{}
	sliceTest(&arr)
	fmt.Println(arr)
}

func sliceTest(arr *[]int) {
	*arr = append(*arr, 10)
	*arr = append(*arr, 123)
}
