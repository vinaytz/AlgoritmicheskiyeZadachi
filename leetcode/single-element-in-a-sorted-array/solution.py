class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        if len(nums) == 1 or nums[0]!=nums[1]:
            return nums[0]
        for i in range(0, len(nums)-1, 2):
            if nums[i]== nums[i+1]:
                continue
            else:
                return nums[i]
        return nums[-1]