#include <iostream>
#include <cstdlib>

class Solution {
public:
    double myPow(double x, long n) {
        if (x == 0) {
            if (n > 0) {
                return 0;
            }
            if (n == 0) {
                return 1;
            }
            if (n < 0) {
                return 1.0/0.0;
            }
        }

        if (n < 0) {
            return myPow(1/x, -n);
        }
        if (n == 0) {
            return 1;
        }
        if (n == 1) {
            return x;
        }
        if (n % 2 == 0) {
            return myPow(x*x, n/2);
        }
        return x * myPow(x*x, (n - 1)/2);
    }
};

int main(int argc, char** argv) {
    Solution* solution = new Solution();
    double x = atof(argv[1]);
    long n = atol(argv[2]);
    std::cout << "base = " << x << ", "
              << "exp = " << n << ", "
              << "pow = " << solution->myPow(x, n)
              << std::endl;
    return 0;
}
