# Stripe Application ID Parser

This problem is inspired by Stripe interview questions and is divided into two parts.

## Part 1: Parse Application IDs
You are given a string representing application IDs in the following format:
- Each application ID is prefixed by its length (number of characters in the ID).
- The format is: lengthOfApplicationId + APPLICATION_ID + ... + 0 (ends with a 0).

**Example:**
Input: `10A13414124218B124564356434567430`
Output: `["A134141242", "B12456435643456743"]`

## Part 2: Whitelist Filter
Filter the application IDs obtained from Part 1 to return only the "whitelisted" application IDs.

**Example:**
Input: `10A13414124218B124564356434567430`, `["A134141242"]`
Output: `["A134141242"]`

---

You may use any programming language. Focus on correctness and code quality.
