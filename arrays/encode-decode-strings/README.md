# 271. Encode and Decode Strings

[NeetCode Problem Link](https://neetcode.io/problems/string-encode-and-decode?list=neetcode150)

## Description

Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Implement the `Encode` and `Decode` methods:

- `Encode(strs []string) string`: Encodes a list of strings to a single string.
- `Decode(s string) []string`: Decodes a single string to a list of strings.

**Example:**

```
Input: ["neet","code","love","you"]
Output: ["neet","code","love","you"]
```

**Note:**
- The string may contain any possible characters out of 256 valid ascii characters.
- Do not use any built-in serialization functions such as `json.dumps` or `join`.
