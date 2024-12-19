# 6. Zigzag Conversion

<a href="https://github.com/whateverzpy/LeetCode-Markdown"><img src="https://img.shields.io/badge/Markdown-FFA116?logo=leetcode&labelColor=555"/></a>
![](https://img.shields.io/badge/Difficulty-Medium-orange)
![](https://img.shields.io/badge/Topics-String-blue)

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: \(you may want to display this pattern in a fixed font for better legibility\)

> P   A   H   N <br>
> A P L S I I G <br>
> Y   I   R <br>

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

> string convert\(string s, int numRows\); <br>

<br>

**Example 1:**

> **Input:**  s = "PAYPALISHIRING", numRows = 3 <br>
> **Output:**  "PAHNAPLSIIGYIR" <br>

**Example 2:**

> **Input:**  s = "PAYPALISHIRING", numRows = 4 <br>
> **Output:**  "PINALSIGYAHRPI" <br>
> **Explanation:** <br>
> P     I    N <br>
> A   L S  I G <br>
> Y A   H R <br>
> P     I <br>

**Example 3:**

> **Input:**  s = "A", numRows = 1 <br>
> **Output:**  "A" <br>

<br>

**Constraints:**

*   `1 <= s.length <= 1000`
*   `s` consists of English letters \(lower\-case and upper\-case\), `','` and `'.'`\.
*   `1 <= numRows <= 1000`
