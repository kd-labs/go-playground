package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScoreOfString(t *testing.T) {
	// table tests
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "score of code",
			input:    "code",
			expected: 24,
		},
		{
			name:     "score of neetcode",
			input:    "neetcode",
			expected: 65,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := scoreOfString(tC.input)
			require.Equal(t, tC.expected, actual)
		})
	}
}

func TestConcatenation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "for [1,4,1,2]",
			input:    []int{1, 4, 1, 2},
			expected: []int{1, 4, 1, 2, 1, 4, 1, 2},
		},
		{
			name:     "for [22,21,20,1]",
			input:    []int{22, 21, 20, 1},
			expected: []int{22, 21, 20, 1, 22, 21, 20, 1},
		},
	}

	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual := getConcatenation(tC.input)
			require.Equal(t, tC.expected, actual)
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		t      string
		expect bool
	}{
		{
			name:   "it should return true for 'racecar' and 'carrace'",
			s:      "racecar",
			t:      "carrace",
			expect: true,
		},
		{
			name:   "it should return false for 'jam' and 'jar'",
			s:      "jam",
			t:      "jar",
			expect: false,
		},
		{
			name:   "it should return false for 'xx' and 'bb'",
			s:      "xx",
			t:      "bb",
			expect: false,
		},
	}

	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual := isAnagram(tC.s, tC.t)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestReplaceElements(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		expect []int
	}{
		{
			name:   "it should return [5, 5, 3, 2, 2, -1]",
			arr:    []int{2, 4, 5, 3, 1, 2},
			expect: []int{5, 5, 3, 2, 2, -1},
		},
		{
			name:   "it should return [3,-1]",
			arr:    []int{3, 3},
			expect: []int{3, -1},
		},
		{
			name:   "it should return [-1]",
			arr:    []int{1},
			expect: []int{-1},
		},
	}

	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual := replaceElements(tC.arr)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestAppendCharacters(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		t      string
		expect int
	}{
		{
			name:   "it should return 4 for s 'coaching' and t 'coding'",
			s:      "coaching",
			t:      "coding",
			expect: 4,
		},
		{
			name:   "it should return 0 for s 'abcde' and t 'a'",
			s:      "abcde",
			t:      "a",
			expect: 0,
		},
		{
			name:   "it should return 5 for s 'z' and t 'abcde'",
			s:      "z",
			t:      "abcde",
			expect: 5,
		},
		{
			name:   "it should return 3 for s 'rabbit' and t 'rabbbit'",
			s:      "rabbit",
			t:      "rabbbit",
			expect: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := appendCharacters(tC.s, tC.t)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestLengthOfLastWords(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		expect int
	}{
		{
			name:   "it should 5 for `Hello World`",
			s:      "Hello World",
			expect: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := lengthOfLastWord(tC.s)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestValidWordSquare(t *testing.T) {
	testCases := []struct {
		name   string
		words  []string
		expect bool
	}{
		{
			name:   "case 1: it should return true",
			words:  []string{"abcd", "bnrt", "crmy", "dtye"},
			expect: true,
		},
		{
			name:   "case 2: it should return false",
			words:  []string{"abcd", "bnrt", "crm", "dt"},
			expect: true,
		},
		{
			name:   "case 3: it should return false",
			words:  []string{"ball", "area", "read", "lady"},
			expect: false,
		},
		{
			name:   "case 4: it should return false",
			words:  []string{"ball", "asee", "let", "lep"},
			expect: false,
		},
		{
			name:   "case 5: it should return false",
			words:  []string{"abc", "bde", "cefg"},
			expect: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := validWordSquare(tC.words)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestConfusingNumber(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		expect bool
	}{
		{
			name:   "it should return true for 6",
			n:      6,
			expect: true,
		}, {
			name:   "it should return true for 89",
			n:      89,
			expect: true,
		}, {
			name:   "it should return false for 11",
			n:      11,
			expect: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := confusingNumber(tC.n)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestAnagramMappings(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect []int
	}{
		{
			name:   "example1",
			nums1:  []int{12, 28, 46, 32, 50},
			nums2:  []int{50, 12, 32, 46, 28},
			expect: []int{1, 4, 3, 2, 0},
		},
		{
			name:   "example2",
			nums1:  []int{84, 46},
			nums2:  []int{84, 46},
			expect: []int{0, 1},
		},
		{
			name:   "example3",
			nums1:  []int{12, 28, 12, 46, 32, 50},
			nums2:  []int{50, 12, 32, 46, 28, 12},
			expect: []int{1, 4, 5, 3, 2, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := anagramMappings(tC.nums1, tC.nums2)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestCountSeniors(t *testing.T) {
	testCases := []struct {
		name    string
		details []string
		expect  int
	}{
		{
			name:    "it should return 2",
			details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			expect:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := countSeniors(tC.details)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "TC1: it should return 3",
			nums:   []int{1, 1, 0, 1, 1, 1},
			expect: 3,
		}, {
			name:   "TC2: it should return 3",
			nums:   []int{1, 1, 1, 0, 1, 1},
			expect: 3,
		}, {
			name:   "TC3: it should return 6",
			nums:   []int{1, 1, 1, 1, 1, 1},
			expect: 6,
		}, {
			name:   "TC4: it should return 2",
			nums:   []int{1, 0, 1, 1, 0, 1},
			expect: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := findMaxConsecutiveOnes(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name   string
		strs   []string
		expect string
	}{
		{
			name:   "it should return 'ba'",
			strs:   []string{"bat", "bag", "bang", "band"},
			expect: "ba",
		}, {
			name:   "it should return 'da'",
			strs:   []string{"dance", "dag", "dang", "damage"},
			expect: "da",
		}, {
			name:   "it should return ''",
			strs:   []string{"neet", "leet"},
			expect: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := longestCommonPrefix(tC.strs)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestStringMatching(t *testing.T) {
	testCases := []struct {
		name   string
		words  []string
		expect []string
	}{
		{
			name:   "it should return '[as, hero]'",
			words:  []string{"mass", "as", "hero", "superhero"},
			expect: []string{"as", "hero"},
		}, {
			name:   "it should return '[neet, code]'",
			words:  []string{"neetcode", "neeet", "neet", "code"},
			expect: []string{"neet", "code"},
		}, {
			name:   "it should return '[]'",
			words:  []string{"blue", "green", "bu"},
			expect: []string{},
		}, {
			name:   "it should return '[cat, cats, dog, rat]'",
			words:  []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			expect: []string{"cat", "cats", "dog", "rat"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := stringMatching(tC.words)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name   string
		strs   []string
		expect [][]string
	}{
		{
			name:   "TC1",
			strs:   []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expect: [][]string{{"act", "cat"}, {"pots", "tops", "stop"}, {"hat"}},
		}, {
			name:   "TC2",
			strs:   []string{"x"},
			expect: [][]string{{"x"}},
		}, {
			name:   "TC3",
			strs:   []string{""},
			expect: [][]string{{""}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := groupAnagrams(tC.strs)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		val    int
		expect int
	}{
		{
			name:   "it should return 5",
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		}, {
			name:   "it should return 3",
			nums:   []int{1, 1, 2, 3, 4},
			val:    1,
			expect: 3,
		}, {
			name:   "it should return 0",
			nums:   []int{1},
			val:    1,
			expect: 0,
		}, {
			name:   "it should return 0",
			nums:   []int{1, 1},
			val:    1,
			expect: 0,
		}, {
			name:   "it should return 1",
			nums:   []int{2},
			val:    1,
			expect: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := removeElements(tC.nums, tC.val)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "it should return 5",
			nums:   []int{5, 5, 1, 1, 1, 5, 5},
			expect: 5,
		}, {
			name:   "it should return 2",
			nums:   []int{2, 2, 2},
			expect: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := majorityElement(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestTopKFrequest(t *testing.T) {
	testCases := []struct {
		name   string
		k      int
		nums   []int
		expect []int
	}{
		{
			name:   "it should return [2,0]",
			k:      2,
			nums:   []int{2, 1, 2, 0, 0, 2},
			expect: []int{2, 0},
		}, {
			name:   "it should return [7]",
			k:      1,
			nums:   []int{7, 7},
			expect: []int{7},
		}, {
			name:   "it should return [2,1]",
			k:      2,
			nums:   []int{1, 1, 1, 2, 2, 3},
			expect: []int{1, 2},
		}, {
			name:   "it should return [1,2]",
			k:      2,
			nums:   []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			expect: []int{1, 2},
		}, {
			name:   "it should return [1]",
			k:      1,
			nums:   []int{1},
			expect: []int{1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := topKFrequent(tC.nums, tC.k)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSortArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "it should return [1,1,1,1,2,3,9,10]",
			nums:   []int{10, 9, 1, 1, 1, 2, 3, 1},
			expect: []int{1, 1, 1, 1, 2, 3, 9, 10},
		}, {
			name:   "it should return [1,2,3,5,10]",
			nums:   []int{5, 10, 2, 1, 3},
			expect: []int{1, 2, 3, 5, 10},
		}, {
			name:   "it should return [-50000, 0, 50000]",
			nums:   []int{50000, -50000, 0},
			expect: []int{-50000, 0, 50000},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := sortArray(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name   string
		colors []int
		expect []int
	}{
		{
			name:   "it should return [0,1,1,2]",
			colors: []int{1, 0, 1, 2},
			expect: []int{0, 1, 1, 2},
		}, {
			name:   "it should return [0,1,2]",
			colors: []int{2, 1, 0},
			expect: []int{0, 1, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := sortColors(tC.colors)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestValidSudoku(t *testing.T) {
	testCases := []struct {
		name   string
		board  [][]byte
		expect bool
	}{
		{
			name: "it should return true",
			board: [][]byte{
				{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
				{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
				{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
				{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expect: true,
		}, {
			name: "it should return false",
			board: [][]byte{
				{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
				{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
				{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
				{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
				{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expect: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := isValidSudoku(tC.board)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "it should return 4",
			input:  []int{2, 20, 4, 10, 3, 4, 5},
			expect: 4,
		}, {
			name:   "it should return 7",
			input:  []int{0, 3, 2, 5, 4, 6, 1, 1},
			expect: 7,
		}, {
			name:   "it should return 0",
			input:  []int{},
			expect: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := longestConsecutive(tC.input)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMaxProfitII(t *testing.T) {
	testCases := []struct {
		name   string
		prices []int
		expect int
	}{
		{
			name:   "it should return 7",
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 7,
		}, {
			name:   "it should return 4",
			prices: []int{1, 2, 3, 4, 5},
			expect: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := maxProfitII(tC.prices)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMajorityElementII(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "it should return '[2,5]'",
			nums:   []int{5, 2, 3, 2, 2, 2, 2, 5, 5, 5},
			expect: []int{2, 5},
		}, {
			name:   "it should return '[4]'",
			nums:   []int{4, 4, 4, 4, 4},
			expect: []int{4},
		}, {
			name:   "it should return '[]'",
			nums:   []int{1, 2, 3},
			expect: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := majorityElementII(tC.nums)
			slices.Sort(actual)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "it should return 1",
			input:  []int{-1, 0, -2},
			expect: 1,
		}, {
			name:   "it should return 3",
			input:  []int{2, 4, 1},
			expect: 3,
		}, {
			name:   "it should return 7",
			input:  []int{1, 2, 4, 5, 6, 3, 1},
			expect: 7,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := firstMissingPositive(tC.input)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name   string
		input  []byte
		expect []byte
	}{
		{
			name:   "it should return teen",
			input:  []byte{'n', 'e', 'e', 't'},
			expect: []byte{'t', 'e', 'e', 'n'},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			reverseString(tC.input)
			require.Equal(t, tC.expect, tC.input)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		expect bool
	}{
		{
			name:   "it should return true",
			s:      "Was it a car or a cat I saw?",
			expect: true,
		}, {
			name:   "it should return true",
			s:      ",.",
			expect: true,
		}, {
			name:   "it should return false",
			s:      "tab a cat",
			expect: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := isPalindrome(tC.s)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		expect bool
	}{
		{
			name:   "0: it should return true",
			s:      "aca",
			expect: true,
		},
		{
			name:   "1: it should return false",
			s:      "abbadc",
			expect: false,
		},
		{
			name:   "2: it should return true",
			s:      "abbda",
			expect: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			chkr := PalindromeChecker{}
			actual := chkr.validPalindrome(tC.s)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		name   string
		word1  string
		word2  string
		expect string
	}{
		{
			name:   "it should return axbycz",
			word1:  "abc",
			word2:  "xyz",
			expect: "axbycz",
		}, {
			name:   "it should return aabbbxxc",
			word1:  "ab",
			word2:  "abbxxc",
			expect: "aabbbxxc",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := mergeAlternately(tC.word1, tC.word2)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "it should return [1,2,10,20,20,40]",
			nums1:  []int{10, 20, 20, 40, 0, 0},
			m:      4,
			nums2:  []int{1, 2},
			n:      2,
			expect: []int{1, 2, 10, 20, 20, 40},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			merge(tC.nums1, tC.m, tC.nums2, tC.n)
			require.Equal(t, tC.expect, tC.nums1)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "it should return 4",
			nums:   []int{1, 1, 2, 3, 4},
			expect: 4,
		}, {
			name:   "it should return 3",
			nums:   []int{2, 10, 10, 30, 30, 30},
			expect: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := removeDuplicates(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name    string
		numbers []int
		target  int
		expect  []int
	}{
		{
			name:    "it should return [1,2]",
			numbers: []int{1, 2, 3, 4},
			target:  3,
			expect:  []int{1, 2},
		}, {
			name:    "it should return [2,4]",
			numbers: []int{2, 3, 4},
			target:  6,
			expect:  []int{1, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := twoSum(tC.numbers, tC.target)
			slices.Sort(actual)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect [][]int
	}{
		{
			name:   "it should return [[-1,-1,2], [-1,0,1]]",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		}, {
			name:   "it should return [[-2,0,2]",
			nums:   []int{-2, 0, 0, 2, 2},
			expect: [][]int{{-2, 0, 2}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := threeSum(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFourSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		expect [][]int
	}{
		{
			name:   "it should return [[-3,0,3,3], [-3,1,2,3]",
			nums:   []int{3, 2, 3, -3, 1, 0},
			target: 3,
			expect: [][]int{{-3, 0, 3, 3}, {-3, 1, 2, 3}},
		}, {
			name:   "it should return [[2,2,2,2]]",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			expect: [][]int{{2, 2, 2, 2}},
		}, {
			name:   "it should return [[-4,0,1,2], [-1,-1,0,1]",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			target: -1,
			expect: [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := fourSum(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMaxWaterContainer(t *testing.T) {
	testCases := []struct {
		desc    string
		heights []int
		expect  int
	}{
		{
			desc:    "it should return 36",
			heights: []int{1, 7, 2, 5, 4, 7, 3, 6},
			expect:  36,
		}, {
			desc:    "it should return 1",
			heights: []int{1, 2},
			expect:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := maxWaterContainer(tC.heights)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		k      int
		expect []int
	}{
		{
			desc:   "it should return [5,6,7,8,1,2,3,4]",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:      4,
			expect: []int{5, 6, 7, 8, 1, 2, 3, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rotate(tC.nums, tC.k)
			require.Equal(t, tC.expect, tC.nums)
		})
	}
}

func TestNumRescueBoats(t *testing.T) {
	testCases := []struct {
		desc   string
		people []int
		limit  int
		expect int
	}{
		{
			desc:   "it should return 2",
			people: []int{5, 1, 4, 2},
			limit:  6,
			expect: 2,
		}, {
			desc:   "it should return 4",
			people: []int{1, 3, 2, 3, 2},
			limit:  3,
			expect: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := numRescueBoats(tC.people, tC.limit)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestCalPoints(t *testing.T) {
	testCases := []struct {
		desc   string
		ops    []string
		expect int
	}{
		{
			desc:   "it should return 18",
			ops:    []string{"1", "2", "+", "C", "5", "D"},
			expect: 18,
		}, {
			desc:   "it should return 15",
			ops:    []string{"5", "D", "+", "C"},
			expect: 15,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := calPoints(tC.ops)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		desc   string
		path   string
		expect string
	}{
		{
			desc:   "it should return '/neetcode/practice/courses'",
			path:   "/neetcode/practice//...///../courses",
			expect: "/neetcode/practice/courses",
		}, {
			desc:   "it should return '/'",
			path:   "/..//",
			expect: "/",
		}, {
			desc:   "it should return '/..//_home/a/b/..///'",
			path:   "/neetcode/practice//...///../courses",
			expect: "/_home/a",
		}, {
			desc:   "it should return '/c'",
			path:   "/a/./b/../../c/",
			expect: "/c",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := simplifyPath(tC.path)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		desc          string
		encodedString string
		decodedString string
	}{
		{
			desc:          "it should return abbbabbbc",
			encodedString: "2[a3[b]]c",
			decodedString: "abbbabbbc",
		}, {
			desc:          "it should return kaaaaaaaaaaab",
			encodedString: "k11[a]b",
			decodedString: "kaaaaaaaaaaab",
		}, {
			desc:          "it should return axbzzzcccc",
			encodedString: "axb3[z]4[c]",
			decodedString: "axbzzzcccc",
		}, {
			desc:          "it should return abccdddx",
			encodedString: "ab2[c]3[d]1[x]",
			decodedString: "abccdddx",
		}, {
			desc:          "it should return aaabcbc",
			encodedString: "3[a]2[bc]",
			decodedString: "aaabcbc",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := decodeString(tC.encodedString)
			require.Equal(t, tC.decodedString, actual)
		})
	}
}

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		desc         string
		temperatures []int
		expect       []int
	}{
		{
			desc:         "it should return [1,4,1,2,1,0,0]",
			temperatures: []int{30, 38, 30, 36, 35, 40, 28},
			expect:       []int{1, 4, 1, 2, 1, 0, 0},
		}, {
			desc:         "it should return [0,0,0]",
			temperatures: []int{22, 21, 20},
			expect:       []int{0, 0, 0},
		}, {
			desc:         "it should return [8,1,5,4,3,2,1,1,0,0]",
			temperatures: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			expect:       []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := dailyTemperatures(tC.temperatures)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestStockSpanner(t *testing.T) {
	expect := []int{1, 1, 1, 2, 1, 4, 6}
	actual := make([]int, 0)
	stkSpanner := NewStockSpanner()
	actual = append(actual, stkSpanner.Next(100))
	actual = append(actual, stkSpanner.Next(80))
	actual = append(actual, stkSpanner.Next(60))
	actual = append(actual, stkSpanner.Next(70))
	actual = append(actual, stkSpanner.Next(60))
	actual = append(actual, stkSpanner.Next(75))
	actual = append(actual, stkSpanner.Next(85))

	require.Equal(t, expect, actual)
}

func TestMinStack(t *testing.T) {
	minStk := NewMinStack()

	minStk.Push(1)
	minStk.Push(2)
	minStk.Push(0)
	m := minStk.GetMin()
	require.Equal(t, 0, m)
	minStk.Pop()
	k := minStk.Top()
	require.Equal(t, 2, k)
	m = minStk.GetMin()
	require.Equal(t, 1, m)
}

func TestEvaluateReversePolishNotation(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []string
		expect int
	}{
		{
			desc:   "it should return 5",
			tokens: []string{"1", "2", "+", "3", "*", "4", "-"},
			expect: 5,
		}, {
			desc:   "it should return 6",
			tokens: []string{"4", "13", "5", "/", "+"},
			expect: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := evalRPN(tC.tokens)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestCarFleet(t *testing.T) {
	testCases := []struct {
		desc     string
		target   int
		position []int
		speed    []int
		expect   int
	}{
		{
			desc:     "it should return 1",
			target:   10,
			position: []int{1, 4},
			speed:    []int{3, 2},
			expect:   1,
		}, {
			desc:     "it should return 3",
			target:   10,
			position: []int{4, 1, 0, 7},
			speed:    []int{2, 2, 1, 1},
			expect:   3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := carFleet(tC.target, tC.position, tC.speed)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestLargestRectangleArea(t *testing.T) {
	testCases := []struct {
		desc    string
		heights []int
		expect  int
	}{
		{
			desc:    "it should return 8",
			heights: []int{7, 1, 7, 2, 2, 4},
			expect:  8,
		}, {
			desc:    "it should return 7",
			heights: []int{1, 3, 7},
			expect:  7,
		}, {
			desc:    "it should return 0",
			heights: []int{0, 0},
			expect:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := largestRectangleArea(tC.heights)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFrequencyStack(t *testing.T) {
	// TC1
	freqStack := NewFreqStack()

	freqStack.Push(5)
	freqStack.Push(7)
	freqStack.Push(5)
	freqStack.Push(7)
	freqStack.Push(4)
	freqStack.Push(5)

	require.Equal(t, 5, freqStack.Pop())
	require.Equal(t, 7, freqStack.Pop())
	require.Equal(t, 5, freqStack.Pop())
	require.Equal(t, 4, freqStack.Pop())

	// TC2
	freqStack = NewFreqStack()

	freqStack.Push(4)
	freqStack.Push(0)
	require.Equal(t, 0, freqStack.Pop())
	freqStack.Push(9)
	freqStack.Push(3)
	freqStack.Push(9)
	require.Equal(t, 9, freqStack.Pop())
	require.Equal(t, 3, freqStack.Pop())
	require.Equal(t, 9, freqStack.Pop())
}

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect int
	}{
		{
			desc:   "it should return 3",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 4,
			expect: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := binarySearch(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect int
	}{
		{
			desc:   "it should return 4",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 5,
			expect: 4,
		}, {
			desc:   "it should return 6",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 10,
			expect: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchInsert(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		desc   string
		x      int
		expect int
	}{
		{
			desc:   "it should return 3",
			x:      9,
			expect: 3,
		}, {
			desc:   "it should return 3",
			x:      12,
			expect: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := mySqrt(tC.x)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		desc   string
		piles  []int
		target int
		expect int
	}{
		{
			desc:   "it should return 2",
			piles:  []int{1, 4, 3, 2},
			target: 9,
			expect: 2,
		}, {
			desc:   "it should return 25",
			piles:  []int{25, 10, 23, 4},
			target: 4,
			expect: 25,
		}, {
			desc:   "it should return 23",
			piles:  []int{30, 11, 23, 4, 20},
			target: 6,
			expect: 23,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := minEatingSpeed(tC.piles, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestShipWithinDays(t *testing.T) {
	testCases := []struct {
		desc    string
		weights []int
		days    int
		expect  int
	}{
		{
			desc:    "it should return 15",
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:    5,
			expect:  15,
		}, {
			desc:    "it should return 6",
			weights: []int{3, 2, 2, 4, 1, 4},
			days:    3,
			expect:  6,
		}, {
			desc:    "it should return 3",
			weights: []int{1, 2, 3, 1, 1},
			days:    4,
			expect:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := shipWithinDays(tC.weights, tC.days)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect []int
	}{
		{
			desc:   "it should return [3,4]",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		}, {
			desc:   "it should return [-1,-1]",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			expect: []int{-1, -1},
		}, {
			desc:   "it should return [-1,-1]",
			nums:   []int{},
			target: 0,
			expect: []int{-1, -1},
		}, {
			desc:   "it should return [0,0]",
			nums:   []int{1},
			target: 1,
			expect: []int{0, 0},
		}, {
			desc:   "it should return [0,2]",
			nums:   []int{3, 3, 3},
			target: 3,
			expect: []int{0, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchRange(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFindMin(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "it should return 1",
			nums:   []int{3, 4, 5, 6, 1, 2},
			expect: 1,
		}, {
			desc:   "it should return 0",
			nums:   []int{4, 5, 0, 1, 2, 3},
			expect: 0,
		}, {
			desc:   "it should return 1",
			nums:   []int{4, 5, 1, 2},
			expect: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findMin(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect int
	}{
		{
			desc:   "it should return 4",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		},
		{
			desc:   "it should return -1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			expect: -1,
		},
		{
			desc:   "it should return -1",
			nums:   []int{1},
			target: 0,
			expect: -1,
		},
		{
			desc:   "it should return 4",
			nums:   []int{3, 4, 5, 6, 1, 2},
			target: 1,
			expect: 4,
		},
		{
			desc:   "it should return -1",
			nums:   []int{3, 5, 6, 0, 1, 2},
			target: 4,
			expect: -1,
		},
		{
			desc:   "it should return 1",
			nums:   []int{3, 4, 5, 6, 0, 1, 2},
			target: 4,
			expect: 1,
		},
		{
			desc:   "it should return 1",
			nums:   []int{1, 3},
			target: 3,
			expect: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchInSortedArray(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestMinIdx(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "it should return min idx 2",
			nums:   []int{6, 7, 0, 1, 2, 3, 4},
			expect: 2,
		}, {
			desc:   "it should return min idx 0",
			nums:   []int{1, 3},
			expect: 0,
		}, {
			desc:   "it should return min idx 3",
			nums:   []int{5, 6, 7, 0, 1, 2, 3},
			expect: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findMinIdx(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestSingleNonDuplicate(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "it should return 2",
			nums:   []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expect: 2,
		}, {
			desc:   "it should return 4",
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 8, 8},
			expect: 4,
		}, {
			desc:   "it should return 10",
			nums:   []int{3, 3, 7, 7, 10, 11, 11},
			expect: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := singleNonDuplicate(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFindKRotation(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "it should return 2",
			nums:   []int{6, 9, 2, 4},
			expect: 2,
		}, {
			desc:   "it should return 1",
			nums:   []int{5, 1, 2, 3, 4},
			expect: 1,
		}, {
			desc:   "it should return 0",
			nums:   []int{1, 2},
			expect: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findKRotation(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}

func TestFindKthPositive(t *testing.T) {
	testCases := []struct {
		desc   string
		arr    []int
		k      int
		expect int
	}{
		{
			desc:   "it should return 9",
			arr:    []int{2, 3, 4, 7, 11},
			k:      5,
			expect: 9,
		}, {
			desc:   "it should return 6",
			arr:    []int{1, 2, 3, 4},
			k:      2,
			expect: 6,
		}, {
			desc:   "it should return 1",
			arr:    []int{2},
			k:      1,
			expect: 1,
		}, {
			desc:   "it should return 14",
			arr:    []int{1, 10, 21, 22, 25},
			k:      12,
			expect: 14,
		}, {
			desc:   "it should return 24",
			arr:    []int{1, 7, 11, 14, 29, 31, 40, 44},
			k:      20,
			expect: 24,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findKthPositive(tC.arr, tC.k)
			require.Equal(t, tC.expect, actual)
		})
	}
}
