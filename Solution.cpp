//
// Created by Wasay Usmani on 1/29/24.
//

#include "Solution.h"
#include <cmath>
#include <iostream>

bool Solution::isAnagram(std::string s, std::string t) {
    std::unordered_map<char, int> map;
    for (int i = 0; i < s.size(); i++) {
        if (map.find(s[i]) == map.end()) {
            map[s[i]] = 1;
        } else {
            map[s[i]] += 1;
        }
    }

    for (char &c : t) {
        auto v = map.find(c);
        if (v != map.end()) {
            if (map[c] > 1) {
                map[c] -= 1;
            } else {
                map.erase(c);
            }
        } else
            return false;
    }
    if (map.size() == 0) {
        return true;
    } else
        return false;
}

std::vector<int> Solution::twoSum(std::vector<int>& nums, int target) {
    std::vector<int> result;
    for (int i = 0; i < nums.size(); i++) {
        for (int j = i + 1; j < nums.size(); j++) {
            if (nums[i] + nums[j] == target) {
                result.push_back(i);
                result.push_back(j);
            }
        }
    }
    return result;
}

std::vector<std::vector<std::string>> Solution::groupAnagrams(std::vector<std::string>& strs) {
    std::unordered_map<std::string, std::vector<std::string>> groups;
    std::vector<std::vector<std::string>> result;
    for (auto &s : strs) {
        auto sorted = s;
        std::sort(sorted.begin(), sorted.end());
        groups[sorted].emplace_back(s);
    }
    for (auto &it : groups) {
        result.emplace_back(it.second);
    }
    return result;
}

std::vector<int> Solution::topKFrequent(std::vector<int>& nums, int k) {
    std::unordered_map<int, int> dict;
    std::vector<int> result;
    for (auto &num : nums) {
        if (dict.find(num) != dict.end()) {
            dict[num] += 1;
        } else {
            dict[num] = 1;
        }
    }
    std::vector<std::pair<int, int>> sortedCount(dict.begin(), dict.end());
    std::sort(sortedCount.begin(), sortedCount.end(), [](const auto &a, const auto &b) {
        return a.second > b.second;
    });
    for (int i = 0; i < k; i++) {
        result.emplace_back(sortedCount[i].first);
    }
    return result;
}

std::vector<int> Solution::productExceptSelf(std::vector<int>& nums) {
    std::vector<int> result(nums.size());
    int multi = 1;
    int ans = std::count(nums.begin(), nums.end(), 0);
    if (ans == nums.size() || ans > 1)
        return result;

    for (const auto &e : nums) {
        if (e == 0)
            continue;
        multi *= e;
    }

    for (int i = 0; i < nums.size(); i++) {
        if (ans > 0) {
            if (nums[i] == 0)
                result[i] = multi;
            else
                result[i] = 0;
        } else {
            result[i] = multi * std::pow(nums[i], -1);
        }
    }
    return result;
}

bool Solution::isValid(std::string s) {
    std::stack<char> st;
    for (const auto &a : s) {
        if (a == '{' || a == '[' || a == '(')
            st.push(a);
        else if (!st.empty()) {
            if (a == '}') {
                if (st.top() == '{')
                    st.pop();
                else
                    return false;
            } else if (a == ']') {
                if (st.top() == '[')
                    st.pop();
                else
                    return false;
            } else if (a == ')') {
                if (st.top() == '(')
                    st.pop();
                else
                    return false;
            }
        } else
            return false;
    }

    return st.empty();
}


bool Solution::isPalindrome(std::string s) {
    for (int i = 0; i < s.size(); ++i) {
        if (isalpha(s[i])) {
            if (tolower(s[i]) != tolower(s[s.size() - 1 - i]))
                return false;
        } else if (isdigit(s[i])) {
            if (s[i] != s[s.size() - 1 - i])
                return false;
        }
    }

    return true;
}
