class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        missing = 0
        st = set(arr)
        counter = 0
        i = 1
        while counter <= k:
            if i not in st:
                missing = i
                counter+=1
            if counter==k:
                return missing
            i+=1