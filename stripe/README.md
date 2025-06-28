# Stripe Currency Conversion & Shipping Problem

This problem is based on a Stripe technical phone screen. The task is divided into several parts, each building on the previous one.

## Problem Description

You are given a string representing direct currency conversion rates and associated shipping methods, in the format:

```
USD:CAD:DHL:5,USD:GBP:FEDX:10
```

Each entry means you can convert from the source currency to the target currency using the specified shipping method at the given cost. Only direct conversions are allowed unless otherwise specified.

### Part 1
Write a method to convert a given amount from one currency to another using only direct conversions. Return the converted amount or indicate if conversion is not possible.

### Part 2
Write a method that returns the cost and shipping methods involved, allowing at most one hop (i.e., at most one intermediate currency) in the conversion from one currency to another.

### Part 3
Write a method that returns the minimum cost and the involved shipping methods, allowing at most one hop for the conversion.

---

You may use any programming language. Input parsing and output formatting are up to you. Focus on correctness and code quality.
