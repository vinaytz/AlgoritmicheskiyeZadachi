/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l:=0
    r:=n

    for true{
        m:= (l+r)/2
        g:= guess(m)
        if g == 0{
            return m
        }else if g == -1{
            r = m -1
        }else{
            l = m + 1
        }  
    }
    return -1 
}