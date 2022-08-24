// https://leetcode.com/problems/longest-common-prefix/

#include <string>
#include <vector>
using namespace std;

class Solution {
public:
  string longestCommonPrefix(vector<string> &strs) {
    string common{""};

    if (strs.size() < 1) {
      return common;
    }

    for (int j = 0;; j++) {

      if (j >= strs[0].size()) {
        return common;
      }

      char c = strs[0][j];

      for (int i = 1; i < strs.size(); i++) {
        if (j >= strs[i].size() || strs[i][j] != c) {
          return common;
        }
      }
      common.push_back(c);
    }
  }
};
