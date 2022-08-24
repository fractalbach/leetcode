#include <string>
#include <vector>
using namespace std;

class Solution {
public:
  bool isValid(string s) {
    vector<char> stack;
    for (char c : s) {
      switch (c) {
      case '(':
        stack.push_back(')');
        continue;
      case '[':
        stack.push_back(']');
        continue;
      case '{':
        stack.push_back('}');
        continue;
      }
      if (stack.size() < 1 || c != stack[stack.size() - 1]) {
        return false;
      }
      stack.pop_back();
    }
    return stack.size() == 0;
  }
};