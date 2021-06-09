package main

import (
	"fmt"
	"math"
	"sort"
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
			} else if idx < len(second) && hashmap[char] < hashmap[rune(second[idx])] {
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

type Pair struct {
	first  int
	second int
}

func firstUniqChar(s string) int {
	hashmap := make(map[rune]Pair)
	ans := math.MaxInt32
	for idx, char := range s {
		if _, ok := hashmap[char]; !ok {
			hashmap[char] = Pair{idx, 1}
		} else {
			hashmap[char] = Pair{hashmap[char].first, hashmap[char].second + 1}
		}
	}
	for _, value := range hashmap {
		if value.second == 1 && value.first < ans {
			ans = value.first
		}
	}
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func uncommonFromSentences(s1 string, s2 string) []string {
	s1Split := strings.Split(s1, " ")
	s2Split := strings.Split(s2, " ")
	s1map := make(map[string]int)
	s2map := make(map[string]int)

	for _, val := range s1Split {
		s1map[val]++
	}
	for _, val := range s2Split {
		s2map[val]++
	}

	ans := make(map[string]bool)
	for key, value := range s1map {
		if _, ok := s2map[key]; !ok && value == 1 {
			ans[key] = true
		}
	}
	for key, value := range s2map {
		if _, ok := s1map[key]; !ok && value == 1 {
			ans[key] = true
		}
	}
	ret := []string{}
	for key, _ := range ans {
		ret = append(ret, key)
	}
	return ret
}

func sliceTest(arr *[]int) {
	*arr = append(*arr, 10)
	*arr = append(*arr, 123)
}

type Logger struct {
	hashmap map[string]int
	time    int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{make(map[string]int), 0}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.hashmap[message]; !ok || timestamp >= this.hashmap[message] {
		this.hashmap[message] = timestamp + 10
		return true
	} else if this.hashmap[message] < timestamp {
		return false
	}
	return false
}

func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)
	for _, char := range s {
		hashmap[char]++
	}

	for _, char := range t {
		if _, ok := hashmap[char]; !ok {
			return false
		} else {
			hashmap[char]--
			if hashmap[char] <= 0 {
				delete(hashmap, char)
			}
		}
	}
	return len(hashmap) == 0
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hashmap := make(map[int]int)
	for idx, val := range nums {
		if previdx, ok := hashmap[val]; ok && idx-previdx <= k {
			return true
		}
		hashmap[val] = idx
	}
	return false
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, val := range arr1 {
		map1[val] = true
	}
	for _, val := range arr2 {
		map2[val] = true
	}
	ansmap := make(map[int]bool)
	for _, val := range arr3 {
		_, ok1 := map1[val]
		_, ok2 := map2[val]
		if ok1 && ok2 {
			ansmap[val] = true
		}
	}
	ans := []int{}
	for key, _ := range ansmap {
		ans = append(ans, key)
	}
	sort.Ints(ans)
	return ans

}

func singleNumber(nums []int) int {
	val := nums[0]
	for i := 1; i < len(nums); i++ {
		val ^= nums[i]
	}
	return val
}

func sumOfUnique(nums []int) int {
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}
	sum := 0
	for key, value := range freq {
		if value == 1 {
			sum += key
		}
	}
	return sum
}

func canPermutePalindrome(s string) bool {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	odd := 0
	for _, value := range freq {
		if value%2 != 0 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}
	return true
}

func findKthPositive(arr []int, k int) int {
	set := make(map[int]bool)
	for _, value := range arr {
		set[value] = true
	}
	count := 0
	for i := 1; ; i++ {
		if _, ok := set[i]; !ok {
			count++
		}
		if count == k {
			return i
		}
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, val := range nums1 {
		set[val] = true
	}
	ansset := make(map[int]bool)
	for _, val := range nums2 {
		if set[val] {
			ansset[val] = true
		}
	}
	ans := []int{}
	for key, _ := range ansset {
		ans = append(ans, key)
	}
	return ans
}

func numJewelsInStones(jewels string, stones string) int {
	set := make(map[rune]bool)

	for _, char := range jewels {
		set[char] = true
	}
	count := 0
	for _, char := range stones {
		if set[char] {
			count++
		}
	}
	return count
}
func intersect(nums1 []int, nums2 []int) []int {
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	for _, val := range nums1 {
		freq1[val]++
	}
	for _, val := range nums2 {
		freq2[val]++
	}

	ans := []int{}

	for key, value := range freq2 {
		if val, ok := freq1[key]; ok {
			for i := 0; i < min(value, val); i++ {
				ans = append(ans, key)
			}
		}
	}

	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func isStrobogrammatic(num string) bool {
	stro := make(map[rune]rune)
	stro['0'], stro['1'], stro['6'], stro['8'], stro['9'] = '0', '1', '9', '8', '6'

	if len(num)%2 != 0 && !(num[len(num)/2] == '0' || num[len(num)/2] == '1' || num[len(num)/2] == '8') {
		return false
	}
	for i := 0; i < len(num)/2; i++ {
		curr := num[i]
		if rune(num[len(num)-i-1]) != stro[rune(curr)] {
			return false
		}
	}
	return true
}

func islandPerimeter(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 1 {
				continue
			}
			dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, dir := range dirs {
				nextR := row + dir[0]
				nextC := col + dir[1]
				if !inside(nextR, nextC, grid) || grid[nextR][nextC] == 0 {
					count++
				}

			}
		}
	}
	return count
}

func inside(row int, col int, grid [][]int) bool {
	return row >= 0 && col >= 0 && row < len(grid) && col < len(grid[row])
}
func main() {
	fmt.Println(isStrobogrammatic("181"))
}
