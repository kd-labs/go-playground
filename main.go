package main

import (
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func scoreOfString(str string) int {
	i := 1
	var score, tmp int
	for {
		if i == len(str) {
			break
		}
		tmp = int(str[i]) - int(str[i-1])
		if tmp < 0 {
			tmp = -1 * tmp
		}
		score += tmp
		i++
	}
	return score
}

func main() {
	logger := log.Default()
	logger.Println("Running main...")
}

func getConcatenation(nums []int) []int {
	l := len(nums)
	res := make([]int, 2*l, 2*l)
	for i := 0; i < l; i++ {
		res[i] = nums[i]
		res[i+l] = nums[i]
	}
	return res
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	htable := make(map[rune]int)

	for _, r := range s {
		htable[r] = htable[r] + 1
	}
	for _, r := range t {
		if val, exists := htable[r]; !exists {
			return false
		} else {
			val = val - 1
			if val == 0 {
				delete(htable, r)
			} else {
				htable[r] = val
			}
		}
	}

	return len(htable) == 0
}

func replaceElements(arr []int) []int {
	lgst := -1
	res := make([]int, len(arr))
	i := len(arr) - 1
	for i >= 0 {
		res[i] = lgst
		lgst = max(lgst, arr[i])
		i = i - 1
	}

	return res
}

func appendCharacters(s string, t string) int {
	var i, idx int
	for _, r := range t {
		if idx = strings.IndexRune(s, r); idx == -1 {
			break
		}
		s = s[idx+1:]
		i++
	}
	return len(t) - i
}

type ErrorBucket string

const (
	BucketSVC ErrorBucket = "SVC"
	BucketUSR ErrorBucket = "USR"
)

type AppError struct {
	HTTPStatus int
	Identifier string // e.g. "0001"
	Bucket     ErrorBucket
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrUnauthorized = &AppError{
		HTTPStatus: 401,
		Bucket:     BucketUSR,
		Identifier: "0001",
		Message:    "unauthorized",
	}
	ErrForbidden = &AppError{
		HTTPStatus: 403,
		Bucket:     BucketUSR,
		Identifier: "0002",
		Message:    "forbidden",
	}
	ErrNotFound = &AppError{
		HTTPStatus: 404,
		Bucket:     BucketUSR,
		Identifier: "0003",
		Message:    "not found",
	}
)

type ServiceConfig struct {
	DomainCode  string
	ServiceCode string
}

type errorResponse struct {
	Status    int
	Title     string
	Detail    string
	Instance  string
	ErrorCode string
}

func lengthOfLastWord(s string) uint16 {
	runes := []rune(s)
	n := len(runes) - 1
	var l uint16

	for i := n; i >= 0; i-- {
		if runes[i] != ' ' {
			l += 1
		} else {
			if l > 0 {
				break
			}
		}
	}
	return l
}

func validWordSquare(words []string) bool {
	rows := len(words)
	var cols int
	if rows == 0 {
		return false
	}
	if rows > 0 {
		cols = len(words[0])
	}
	if rows != cols {
		return false
	}

	for i, word := range words {
		for j, r := range word {
			if j >= len(words) || i >= len(words[j]) ||
				r != rune(words[j][i]) {
				return false
			}
		}
	}
	return true
}

func confusingNumber(k int) bool {
	confusing := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	n := k
	var m int
	for ; n > 0; n = n / 10 {
		r := n % 10
		c, exists := confusing[r]
		if !exists {
			return false
		}
		m = (m * 10) + c
	}

	return k != m
}

func anagramMappings(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i, n1 := range nums1 {
		for j, n2 := range nums2 {
			if n1 == n2 {
				res[i] = j
				nums2[j] = -1
				break
			}
		}
	}
	return res
}

func makeSimilarPairsMap(similarPairs [][]string) map[string]string {
	similarPairsMap := make(map[string]string)
	for _, pair := range similarPairs {
		similarPairsMap[pair[0]] = pair[1]
	}

	return similarPairsMap
}

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}

	similarMap := makeSimilarPairsMap(similarPairs)

	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] != sentence2[i] {
			// if the words aren't same, check if they are similar pairs

			if similarMap[sentence1[i]] != sentence2[i] && similarMap[sentence2[i]] != sentence1[i] {
				return false
			}
		}
	}

	return true
}

func countSeniors(details []string) int {
	var res int

	for _, perDeet := range details {
		age, _ := strconv.ParseInt(perDeet[11:13], 10, 64)
		if age > 60 {
			res++
		}
	}
	return res
}

func findMaxConsecutiveOnes(nums []int) int {
	var res, curr int

	for _, n := range nums {
		if n == 1 {
			curr++
		} else {
			res = max(res, curr)
			curr = 0
		}
	}
	// Edge case when nums has all 1
	res = max(res, curr)

	return res
}

func longestCommonPrefix(strs []string) string {
	commonPre := strs[0]

	for i := 1; i < len(strs); i++ {
		commonPre = getPrefix(commonPre, strs[i])
	}

	return commonPre
}

func getPrefix(a, b string) string {
	var p string
	l := min(len(a), len(b))

	if l == 0 {
		return ""
	}

	var i int
	for i = 0; i < l; i++ {
		if a[i] != b[i] {
			break
		}
	}
	p = a[:i]

	return p
}

// O(n^2)
func stringMatching(words []string) []string {
	res := make(map[string]struct{})

	for _, word := range words {
		for _, n := range words {
			if word == n {
				continue
			}
			if strings.Index(n, word) != -1 {
				if _, ok := res[word]; !ok {
					res[word] = struct{}{}
				}
			}
		}
	}

	substr := make([]string, 0, len(res))

	for s := range res {
		substr = append(substr, s)
	}

	return substr
}

func groupAnagrams(strs []string) [][]string {
	anagramsGroupings := make(map[[26]int][]string)
	for _, str := range strs {
		var ht [26]int
		for _, r := range str {
			ht[r-'a']++
		}
		anagramsGroupings[ht] = append(anagramsGroupings[ht], str)
	}

	var res [][]string
	for _, anagrams := range anagramsGroupings {
		res = append(res, anagrams)
	}
	return res
}

func removeElements(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// initially position k from end at non-val
	k := len(nums) - 1
	for nums[k] == val && k > 0 {
		k--
	}

	for i := 0; i < k; i++ {
		if nums[i] == val {
			// swap the vals at i and k
			nums[i], nums[k] = nums[k], nums[i]
			k--
			// again position k at next non-val
			for nums[k] == val && k > 0 {
				k--
			}
		}
	}

	if k == 0 && nums[k] == val {
		return k
	}
	return k + 1
}

func majorityElement(nums []int) int {
	winner := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == winner {
			count++
		} else {
			count--
		}

		if count < 0 {
			winner = nums[i]
			count = 1
		}
	}
	return winner
}

func topKFrequent(nums []int, k int) []int {
	maxLen := 2000
	res := make([]int, 0, k)
	m := func(i int) int {
		return i + 1000
	}
	rev_m := func(i int) int {
		return i - 1000
	}
	sorted := bucketSort(nums, maxLen, m)
	ht := make(map[int][]int)
	for n, count := range sorted {
		ht[count] = append(ht[count], n)
	}
	countsArr := make([]int, len(nums)+1)

	for count, nums := range ht {
		countsArr[count] = len(nums)
	}

	for m := len(nums); m >= 0 && k > 0; m-- {
		if countsArr[m] > 0 {
			for _, n := range ht[m] {
				res = append(res, rev_m(n))
				k--
			}
		}
	}
	return res
}

func bucketSort(nums []int, len int, idxFn func(int) int) []int {
	bucket := make([]int, len)

	for _, n := range nums {
		bucket[idxFn(n)]++
	}

	return bucket
}

func sortArray(nums []int) []int {
	numToIdxMapper := func(num int) int {
		return num + 50000
	}
	bucketSorted := bucketSort(nums, 100001, numToIdxMapper)

	var res []int
	for idx, val := range bucketSorted {
		for val > 0 {
			res = append(res, idx-50000)
			val--
		}
	}
	return res
}

func sortColors(colors []int) []int {
	quickSort(colors, 0, len(colors)-1)
	return colors
}

func quickSort(nums []int, start int, end int) {
	// Base Case
	if end-start+1 <= 1 {
		return
	}

	pivot := nums[end]
	partitionIdx := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			// swap value at i with sortIdx
			nums[partitionIdx], nums[i] = nums[i], nums[partitionIdx]
			partitionIdx++
		}
	}

	nums[end], nums[partitionIdx] = nums[partitionIdx], pivot

	// recursive call for left and right partitions
	quickSort(nums, start, partitionIdx-1)
	quickSort(nums, partitionIdx+1, end)
}

func isValidSudoku(board [][]byte) bool {
	type coord struct {
		r int
		c int
	}

	rowSet := make(map[int]map[byte]struct{})
	colSet := make(map[int]map[byte]struct{})
	subMatrixMap := make(map[coord]map[byte]struct{})

	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[byte]struct{})
		colSet[i] = make(map[byte]struct{})

		if i%3 == 0 {
			for k := 0; k < 3; k++ {
				subMatrixMap[coord{i / 3, k}] = make(map[byte]struct{})
			}
		}
	}

	for r, row := range board {
		for c, n := range row {
			if n == '.' {
				continue
			}
			if _, found := rowSet[r][n]; found {
				return false
			}
			if _, found := colSet[c][n]; found {
				return false
			}

			subMatrixCoord := coord{r / 3, c / 3}
			if _, found := subMatrixMap[subMatrixCoord][n]; found {
				return false
			}

			rowSet[r][n] = struct{}{}
			colSet[c][n] = struct{}{}
			subMatrixMap[subMatrixCoord][n] = struct{}{}
		}
	}
	return true
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]struct{})
	longest := 1

	for _, n := range nums {
		set[n] = struct{}{}
	}

	for _, n := range nums {
		// check if 'n' is in middle of a sequence
		if _, found := set[n-1]; found {
			continue
		}

		// at this point, 'n' is established to be start of sequence
		// check if subsequent n+1 exists
		k := n
		l := 1
		for {
			if _, exists := set[k+1]; exists {
				k = k + 1
				l++
			} else {
				break
			}
		}
		longest = max(l, longest)
	}
	return longest
}

func maxProfitII(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit := 0
	buyPrice := prices[0]
	prices = prices[1:]
	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
		} else {
			profit += price - buyPrice
			buyPrice = price
		}
	}
	return profit
}

func majorityElementII(nums []int) []int {
	majorityEls := make(map[int]int)
	partitionSize := len(nums) / 3

	for _, n := range nums {
		majorityEls[n] += 1

		if len(majorityEls) == 3 {
			for k, count := range majorityEls {
				if count == 1 {
					delete(majorityEls, k)
				}
				majorityEls[k] = count - 1
			}
		}
	}

	res := make([]int, 0)
	for k := range majorityEls {
		count := 0
		for _, n := range nums {
			if k == n {
				count++
			}
		}
		if count > partitionSize {
			res = append(res, k)
		}
	}
	return res
}

func firstMissingPositive(nums []int) int {
	hasOne := false

	size := len(nums)

	for i, n := range nums {
		if n == 1 {
			hasOne = true
		} else if n <= 0 || n > size {
			nums[i] = 1
		}
	}

	if !hasOne {
		return 1
	}

	for _, n := range nums {
		i := abs(n) - 1
		if nums[i] > 0 {
			nums[i] = -nums[i]
		}
	}

	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}
	return size + 1
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}

	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	s = sanitize(s)

	if len(s) == 0 || len(s) == 1 {
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNumeric(c rune) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func sanitize(s string) string {
	var newS []rune
	for _, c := range s {
		if isAlphaNumeric(c) {
			newS = append(newS, c)
		}
	}
	return string(newS)
}

type PalindromeChecker struct {
	deleted bool
}

func (p PalindromeChecker) validPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return p.validPalindrome(s[1 : len(s)-1])
	} else {
		if p.deleted {
			return false
		} else {
			p.deleted = true
		}
		return p.validPalindrome(s[1:]) || p.validPalindrome(s[:len(s)-1])
	}
}

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	sb := strings.Builder{}

	for i < len(word1) && j < len(word2) {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		sb.WriteByte(word1[i])
		i++
	}

	for j < len(word2) {
		sb.WriteByte(word2[j])
		j++
	}

	return sb.String()
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m)

	for i := range m {
		tmp[i] = nums1[i]
	}

	k := 0
	p := 0
	j := 0

	for k < m && p < n {
		if tmp[k] < nums2[p] {
			nums1[j] = tmp[k]
			k++
		} else {
			nums1[j] = nums2[p]
			p++
		}
		j++
	}

	for k < m {
		nums1[j] = tmp[k]
		j++
		k++
	}

	for p < n {
		nums1[j] = nums2[p]
		j++
		p++
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	uniqueIdx := 0
	i := 1

	for i < len(nums) {
		if nums[i] != nums[i-1] {
			uniqueIdx++
			nums[uniqueIdx] = nums[i]
		}
		i++
	}

	return uniqueIdx + 1
}

func twoSum(numbers []int, target int) []int {
	var res []int

	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			res = append(res, l+1, r+1)
			break
		} else if sum > target {
			r--
		} else {
			l++
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	s := len(nums)

	for i, a := range nums {

		if i > 0 && a == nums[i-1] {
			continue
		}

		target := -a
		l := i + 1
		r := s - 1

		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	var res [][]int

	sort.Ints(nums)
	s := len(nums)

	for i := 0; i < s-3; i++ {

		a := nums[i]

		if i > 0 && a == nums[i-1] {
			continue
		}

		for k := i + 1; k < len(nums)-2; k++ {
			b := nums[k]

			if k-i > 1 && b == nums[k-1] {
				continue
			}

			twoSumTarget := target - a - b
			l := k + 1
			r := s - 1

			for l < r {
				sum := nums[l] + nums[r]
				if sum == twoSumTarget {
					res = append(res, []int{a, b, nums[l], nums[r]})
					l++

					for nums[l] == nums[l-1] && l < r {
						l++
					}
				} else if sum > twoSumTarget {
					r--
				} else {
					l++
				}
			}

		}
	}
	return res
}

func maxWaterContainer(heights []int) int {
	if len(heights) == 2 {
		smallColumn := min(heights[0], heights[1])
		return smallColumn * smallColumn
	}

	var maxWater int
	l := 0
	r := len(heights) - 1

	for l < r {
		water := min(heights[l], heights[r]) * (r - l)
		maxWater = max(maxWater, water)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxWater
}

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	k = k % len(nums)

	// reverse the whole slice
	l := 0
	r := len(nums) - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	// revers the first k elements
	l = 0
	r = k - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	// reverse the remaining elements
	l = k
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	var res int
	if len(people) == 1 {
		return 1
	}

	l := 0
	r := len(people) - 1

	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		res++
		r--
	}

	return res
}

func calPoints(ops []string) int {
	records := make([]int, 1000)
	i := -1

	for _, op := range ops {
		switch op {
		case "+":
			op1 := records[i]
			op2 := records[i-1]
			i++
			records[i] = op1 + op2
		case "D":
			i++
			prevScore := records[i-1]
			records[i] = 2 * prevScore
		case "C":
			records[i] = 30001
			i--
		default:
			score, _ := strconv.Atoi(op)
			i++
			records[i] = score
		}
	}

	var total int
	for k := 0; k <= i; k++ {
		if records[k] == 30001 {
			continue
		}
		total += records[k]
	}

	return total
}

type Stack[T any] struct {
	arr []T
	tos int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		arr: make([]T, 10000),
		tos: -1,
	}
}

func (s *Stack[T]) Push(val T) {
	s.tos++
	s.arr[s.tos] = val
}

func (s *Stack[T]) Pop() T {
	if s.tos < 0 {
		var a T
		return a
	}
	res := s.arr[s.tos]
	s.tos--
	return res
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

func (s *Stack[T]) Peek() T {
	var p T
	if s.tos >= 0 {
		return s.arr[s.tos]
	}
	return p
}

func (s *Stack[T]) ToArray() []T {
	res := make([]T, 0, 100)
	for i := range s.tos + 1 {
		res = append(res, s.arr[i])
	}
	return res
}

func (s *Stack[T]) isEmpty() bool {
	return s.tos < 0
}

func asteroidCollision(asteroids []int) []int {
	stk := NewStack[int]()

	var asteroidDestroyed bool
	for _, asteroid := range asteroids {
		asteroidDestroyed = false
		if stk.isEmpty() || stk.Peek() < 0 || asteroid > 0 {
			stk.Push(asteroid)
		} else {
			// control will enter only when incoming asteroid is -ve
			for !stk.isEmpty() && stk.Peek() > 0 {
				o := stk.Pop()
				if o+asteroid < 0 {
					continue
				} else if o+asteroid == 0 {
					asteroidDestroyed = true
					break
				} else {
					asteroidDestroyed = true
					stk.Push(o)
					break
				}
			}

			if !asteroidDestroyed {
				stk.Push(asteroid)
			}
		}
	}

	return stk.ToArray()
}

func modInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func simplifyPath(path string) string {
	prevDir := ".."
	pathFrags := strings.Split(path, "/")

	stk := NewStack[string]()

	for _, frag := range pathFrags {
		if len(frag) == 0 || frag == "." {
			continue
		} else if frag == prevDir {
			stk.Pop()
		} else {
			stk.Push(frag)
		}
	}

	res := strings.Builder{}
	res.WriteString("/")
	fragArr := stk.ToArray()
	res.WriteString(strings.Join(fragArr, "/"))

	return res.String()
}

func decodeString(s string) string {
	if len(s) < 4 {
		return s
	}

	stk := NewStack[any]()

	var res strings.Builder

	var multiplier int
	for _, c := range s {

		digit, err := strconv.ParseInt(string(c), 10, 8)
		if err == nil {
			// rune is a digit
			multiplier = multiplier*10 + int(digit)
		} else if c == '[' {
			stk.Push(multiplier)
			multiplier = 0
		} else if c == ']' {
			var s string
			var n int

			for {
				o := stk.Pop()
				tmp, ok := o.(string)
				if ok {
					s = tmp + s
				} else {
					n, _ = o.(int)
					break
				}
			}

			r := strings.Repeat(s, n)

			if stk.isEmpty() {
				res.WriteString(r)
			} else {
				stk.Push(r)
			}
		} else if stk.isEmpty() {
			// rune is a letter but outside brackets
			res.WriteRune(c)
		} else {
			stk.Push(string(c))
		}
	}

	return res.String()
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	i := len(temperatures) - 1

	stk := NewStack[int]()

	stk.Push(i)
	res[i] = 0
	i--

	for ; i >= 0; i-- {
		currTemp := temperatures[i]

		for !stk.isEmpty() && currTemp >= temperatures[stk.Peek()] {
			_ = stk.Pop()
		}

		if stk.isEmpty() {
			res[i] = 0
		} else {
			res[i] = stk.Peek() - i
		}

		stk.Push(i)
	}

	return res
}

///////// Stock Spanner /////////////////////

type StockSpanner struct {
	stk *Stack[int]
}

func NewStockSpanner() StockSpanner {
	return StockSpanner{
		stk: NewStack[int](),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := this.span(price)
	this.stk.Push(price)
	return span
}

func (this *StockSpanner) span(price int) int {
	res := 1

	arr := this.stk.ToArray()
	i := len(arr) - 1
	for i >= 0 && arr[i] <= price {
		res++
		i--
	}

	return res
}

/////////////////////////////////////////////////

type Pair struct {
	val int
	min int
}

type MinStack struct {
	arr []Pair
	tos int
}

func NewMinStack() MinStack {
	return MinStack{
		arr: make([]Pair, 1000),
		tos: -1,
	}
}

func (this *MinStack) Push(val int) {
	if this.tos == -1 {
		this.tos++
		this.arr[this.tos] = Pair{val: val, min: val}
		return
	}

	min := this.arr[this.tos].min
	if min > val {
		min = val
	}
	this.tos++
	this.arr[this.tos] = Pair{val, min}
}

func (this *MinStack) Pop() {
	this.tos--
}

func (this *MinStack) Top() int {
	return this.arr[this.tos].val
}

func (this *MinStack) GetMin() int {
	return this.arr[this.tos].min
}

func evalRPN(tokens []string) int {
	ops := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	stk := NewStack[string]()

	var res int
	for _, token := range tokens {
		if _, exists := ops[token]; exists {
			op2, _ := strconv.Atoi(stk.Pop())
			op1, _ := strconv.Atoi(stk.Pop())
			res = eval(token, op1, op2)
			stk.Push(strconv.Itoa(res))
		} else {
			stk.Push(token)
		}
	}
	res, _ = strconv.Atoi(stk.Pop())
	return res
}

func eval(operator string, op1, op2 int) int {
	switch operator {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	default:
		return op1 / op2
	}
}

func carFleet(target int, position []int, speed []int) int {
	type Car struct {
		position int
		speed    int
	}

	carsArr := make([]Car, 0, len(position))

	for i, p := range position {
		carsArr = append(carsArr, Car{position: p, speed: speed[i]})
	}

	// slices.SortFunc(carsArr, func(a, b Car) int { return a.position - b.position })
	sort.SliceStable(carsArr, func(i, j int) bool { return carsArr[i].position < carsArr[j].position })

	stk := NewStack[float64]()

	for i := len(carsArr) - 1; i >= 0; i-- {
		car := carsArr[i]
		time := float64(target-car.position) / float64(car.speed)

		if stk.isEmpty() {
			stk.Push(time)
			continue
		}

		for !stk.isEmpty() && time > stk.Peek() {
			stk.Push(time)
		}
	}

	return stk.tos + 1
}

func largestRectangleArea(heights []int) int {
	var maxArea int

	type pair struct {
		idx    int
		height int
	}

	stk := NewStack[pair]()

	for i, h := range heights {
		if stk.isEmpty() {
			stk.Push(pair{i, h})
			continue
		}

		p := pair{-1, -1}
		for !stk.isEmpty() && stk.Peek().height >= h {
			// pop from Stack
			p = stk.Pop()

			// calculate area
			tmp := p.height * (i - p.idx)
			// compare with maxArea
			maxArea = max(maxArea, tmp)
		}
		// push current height with idx to which extends back
		if p.idx > -1 {
			stk.Push(pair{p.idx, h})
		} else {
			stk.Push(pair{i, h})
		}
	}

	// cal area remaining in stack
	s := len(heights)
	for !stk.isEmpty() {
		p := stk.Pop()
		tmp := p.height * (s - p.idx)
		maxArea = max(maxArea, tmp)
	}

	return maxArea
}

type FreqStack struct {
	arr          []int
	tos          int
	numFreqCount map[int]int
	maxCount     int
}

func NewFreqStack() FreqStack {
	return FreqStack{
		arr:          make([]int, 20000),
		tos:          -1,
		numFreqCount: make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.tos++
	this.arr[this.tos] = val
	this.numFreqCount[val]++
	this.maxCount = max(this.maxCount, this.numFreqCount[val])
}

func (this *FreqStack) Pop() int {
	k := []int{}
	for n, c := range this.numFreqCount {
		if c == this.maxCount {
			k = append(k, n)
		}
	}

	var i, res int
	if len(k) == 1 {

		// adjust the internal storage
		for i = this.tos; i >= 0; i-- {
			if this.arr[i] == k[0] {
				break
			}
		}
		for k := i + 1; k <= this.tos; k++ {
			this.arr[i] = this.arr[k]
			i++
		}

		this.tos--
		this.maxCount--
		res = k[0]
	} else {
	outer:
		for i = this.tos; i >= 0; i-- {
			for _, l := range k {
				if this.arr[i] == l {
					res = l
					break outer
				}
			}
		}

		for d := i + 1; d <= this.tos; d++ {
			this.arr[i] = this.arr[d]
			i++
		}

		this.tos--
	}
	this.numFreqCount[res]--

	return res
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	lo, hi := 0, len(nums)-1

	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func mySqrt(x int) int {
	var res int

	lo, hi := 0, x

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid*mid <= x {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return res
}

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, getMax(piles)

	ans := hi
	for lo <= hi {

		rate := lo + (hi-lo)/2

		var totalTime int
		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(rate)))
		}

		if totalTime <= h {
			ans = rate
			hi = rate - 1
		} else {
			lo = rate + 1
		}
	}
	return ans
}

func getMax(arr []int) int {
	var maxEl int
	for _, n := range arr {
		maxEl = max(maxEl, n)
	}

	return maxEl
}

func shipWithinDays(weights []int, days int) int {
	lo := 0

	var hi int
	for _, wt := range weights {
		// maximum ship capacity is sum of all weigghts to carry all cargo together in one day
		hi += wt
		// creative way to set the lower bound of search range to max cargo weight
		// because, need minimum ship capacity to carry the heaviest cargo.
		if wt > lo {
			lo = wt
		}
	}

	minWt := hi
	for lo <= hi {
		cap := lo + (hi-lo)/2
		d := 1
		var cmltvWt int
		for _, wt := range weights {
			cmltvWt += wt
			if cmltvWt > cap {
				d++
				cmltvWt = wt
			}
		}

		if d <= days {
			minWt = cap
			hi = cap - 1
		} else {
			lo = cap + 1
		}
	}

	return minWt
}

func searchRange(nums []int, target int) []int {
	defaultRes := []int{-1, -1}

	if len(nums) == 0 {
		return defaultRes
	}

	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			leftBound, rightBound := mid, mid
			// find left boundary
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					leftBound = i
				}
			}

			// find right bound
			for i := mid; i < len(nums); i++ {
				if nums[i] == target {
					rightBound = i
				}
			}

			return []int{leftBound, rightBound}
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return defaultRes
}
