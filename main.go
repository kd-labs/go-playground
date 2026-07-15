package main

import (
	"log"
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

func twoSum(nums []int, target int) []int {
	htable := make(map[int]int)
	for i, a := range nums {
		b := target - a
		k, exists := htable[b]
		if exists {
			return []int{i, k}
		}
		htable[a] = i
	}
	return nil
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
