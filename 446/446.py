import collections
import random
import math
import pprint

class Solution(object):
    def numberOfArithmeticSlices(self, A):
        """
        :type A: List[int]
        :rtype: inte
        """
        size = len(A)
        total = 0
        container = [collections.defaultdict(int) for x in range(size)]
        for i in range(size):
            for j in range(i):
                diff = A[i] - A[j]
                container[i][diff] += 1
                if diff in container[j]:
                    container[i][diff] += container[j][diff]
                    total += container[j][diff]
        return total

s = Solution()
print s.numberOfArithmeticSlices([1,2,3,4,5,6,5,4,3,2,1])
