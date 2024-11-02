package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/xidream/leetcode/problem/0005_longest_palindromic_substring"
)

var _ = DescribeTable("LongestPalindrome",
	func(s string, expected string) {
		Expect(LongestPalindrome(s)).To(Equal(expected))
	},
	Entry("Example 1", "babad", "aba"),
	Entry("Example 2", "cbbd", "bb"),
	Entry("a", "a", "a"),
	Entry("aa", "aa", "aa"),
	Entry("ab", "ab", "a"),
	Entry("aaa", "aaa", "aaa"),
	Entry("aba", "aba", "aba"),
	Entry("baa", "baa", "aa"),
	Entry("aaaaaa", "aaaaaa", "aaaaaa"),
	Entry("abaaaa", "abaaaa", "aaaa"),
	Entry("aaabaa", "aaabaa", "aabaa"),
	Entry("aaaaaaa", "aaaaaaa", "aaaaaaa"),
	Entry("abaaaaa", "abaaaaa", "aaaaa"),
	Entry("aabaaaa", "aabaaaa", "aabaa"),
)

var _ = DescribeTable("GetLongestPalindromeFromMiddle",
	func(s, expected string) {
		Expect(GetLongestPalindromeFromMiddle(s)).To(Equal(expected))
	},
	Entry("a", "a", "a"),
	Entry("aa", "aa", "aa"),
	Entry("ab", "ab", "a"),
	Entry("aaa", "aaa", "aaa"),
	Entry("aba", "aba", "aba"),
	Entry("baa", "baa", "b"),
	Entry("aaaaaa", "aaaaaa", "aaaaaa"),
	Entry("abaaaa", "abaaaa", "aa"),
	Entry("aaabaa", "aaabaa", "a"),
	Entry("aaaaaaa", "aaaaaaa", "aaaaaaa"),
	Entry("abaaaaa", "abaaaaa", "aaa"),
	Entry("aabaaaa", "aabaaaa", "a"),
)
