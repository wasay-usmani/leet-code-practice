#pragma once
#include <string>
#include <vector>
#include <unordered_map>
#include <algorithm>
#include <stack>

class Solution {
public:
    // Array and hashing
    static bool isAnagram(std::string s, std::string t);
    static std::vector<int> twoSum(std::vector<int>& nums, int target);
    static std::vector<std::vector<std::string>> groupAnagrams(std::vector<std::string>& strs);
    static std::vector<int> topKFrequent(std::vector<int>& nums, int k);
    std::vector<int> productExceptSelf(std::vector<int>& nums);
    bool isValid(std::string s);

    // Two pointers
    bool isPalindrome(std::string s);
};
