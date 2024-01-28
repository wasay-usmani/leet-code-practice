#include <iostream>
#include "Solution.h"

using namespace std;

int main()
{
    Solution s;
    std::vector<string> test = {"eat", "tea", "tan", "ate", "nat", "bat"};
    s.groupAnagrams(test);
    vector<int> test2 = {1, 2, 2, 3, 3, 3};
    s.topKFrequent(test2, 2);
    vector<int> v = {0, 4, 0};
    s.productExceptSelf(v);
    s.isValid("[");
    s.isPalindrome("A man, a plan, a canal: Panama");
    return 0;
}
