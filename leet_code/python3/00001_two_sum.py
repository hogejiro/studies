class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_map = {}
        for idx, num in enumerate(nums):
            n = target - num
            if n in hash_map:
                return sorted([idx, hash_map[n]])
            hash_map[num] = idx
