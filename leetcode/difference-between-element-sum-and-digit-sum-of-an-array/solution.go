func differenceOfSum(nums []int) int {
    sum1 :=0
    for _,x :=range nums{
        sum1+=x
    }
    sum2 :=0
    for _,x:=range nums{
        if x<10{
            sum2+=x
        }else{
            for x>0{
                digit := x%10
                sum2+=digit
                x = x/10
            }
        }
    }
    if sum1-sum2<0{
        return sum2-sum1
    }
    return sum1-sum2
}
