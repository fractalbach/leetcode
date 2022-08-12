// https://leetcode.com/problems/reverse-integer/
class Solution {
public:
  static constexpr std::array<int, 10> getTens() {
    std::array<int, 10> ten{};
    for (int i = 0; i < 10; i++) {
      ten[i] = pow(10, i);
    }
    return ten;
  }

  const std::array<int, 10> ten = getTens();

  int reverse(int x) {
    int result = 0;
    int digit = 0;
    bool padding_done = false;
    int current_exponent = 0;

    for (int i = 0; i < 10; i++) {

      digit = x / ten[9 - i];

      if (!padding_done && digit != 0)
        padding_done = true;

      if (padding_done) {

        // avoid over/underflow when multiplying by 10^9
        if (current_exponent == 9 && (digit > 2 || digit < -2))
          return 0;

        int y = digit * ten[current_exponent];

        // check for overflow and underflow
        if (result > 0 && y > INT32_MAX - result)
          return 0;
        if (result < 0 && y < INT32_MIN - result)
          return 0;

        x -= digit * ten[9 - i];
        result += y;
        current_exponent++;
      }
    }

    return result;
  }
};
