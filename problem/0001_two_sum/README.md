# 1. Two Sum

<a href="https://github.com/whateverzpy/LeetCode-Markdown"><img src="https://img.shields.io/badge/Markdown-FFA116?logo=leetcode&labelColor=555"/></a>
![](https://img.shields.io/badge/Difficulty-Easy-green)
![](https://img.shields.io/badge/Topic-Array,_Hash_Table-blue)

Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_\.

You may assume that each input would have **_exactly_ one solution** , and you may not use the _same_ element twice\.

You can return the answer in any order\.

<br>

**Example 1:**

> **Input:**  nums = \[2,7,11,15], target = 9 <br>
> **Output:**  \[0,1] <br>
> **Explanation:**  Because nums\[0] \+ nums\[1] == 9, we return \[0, 1]\. <br>

**Example 2:**

> **Input:**  nums = \[3,2,4], target = 6 <br>
> **Output:**  \[1,2] <br>

**Example 3:**

> **Input:**  nums = \[3,3], target = 6 <br>
> **Output:**  \[0,1] <br>

<br>

**Constraints:**

*   `2 <= nums.length <= 10^4`
*   `-10^9 <= nums[i] <= 10^9`
*   `-10^9 <= target <= 10^9`
*   **Only one valid answer exists\.**

<br>

**Follow\-up:**  Can you come up with an algorithm that is less than `O(n^2)` time complexity?
