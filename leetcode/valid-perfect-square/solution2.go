func isPerfectSquare(num int) bool {
    l:= 1
    r := num

    for l<=r{
        m:= (l+r)/2
        if int(m*m) == num{
            return true
        }
        if int(m*m) > num{
            r = m -1
        }else{
            l = m + 1
        }
    }

    return false
}