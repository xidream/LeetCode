# 4. Median of Two Sorted Arrays

<a href="https://github.com/whateverzpy/LeetCode-Markdown"><img src="https://img.shields.io/badge/Markdown-FFA116?logo=leetcode&labelColor=555"/></a>
![](https://img.shields.io/badge/Difficulty-Hard-red)
![](https://img.shields.io/badge/Topics-Array,_Binary_Search,_Divide_and_Conquer-blue)

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median**  of the two sorted arrays\.

The overall run time complexity should be `O(log (m+n))`\.

<br>

**Example 1:**

> **Input:**  nums1 = \[1,3], nums2 = \[2]  
> **Output:**  2\.00000  
> **Explanation:**  merged array = \[1,2,3] and median is 2\.
>

**Example 2:**

> **Input:**  nums1 = \[1,2], nums2 = \[3,4]  
> **Output:**  2\.50000  
> **Explanation:**  merged array = \[1,2,3,4] and median is \(2 \+ 3\) / 2 = 2\.5\.
>

<br>

**Constraints:**

*   `nums1.length == m`
*   `nums2.length == n`
*   `0 <= m <= 1000`
*   `0 <= n <= 1000`
*   `1 <= m + n <= 2000`
*   `-10^6 <= nums1[i], nums2[i] <= 10^6`
