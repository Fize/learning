#!/usr/bin/env python

class Solution:
    def twoSum(self, nums, target):
        d = {}
        for k, v in enumerate(nums):
            if target-v in d:
                return [d[target-v], k]
            else:
                d[v] = k
        return []


if __name__ == '__main__':
    nums_0 = [2, 7, 11, 15]
    target_0 = 9
    nums_1 = [3, 5, 11, 11, 16]
    target_1 = 22
    t = Solution()
    print(t.twoSum(nums_0, target_0))
    print(t.twoSum(nums_1, target_1))
