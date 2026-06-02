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
