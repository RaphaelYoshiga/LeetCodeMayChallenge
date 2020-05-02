package main;

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
 func isBadVersion(i int) bool{
	 return true;
 }
    
func firstBadVersion(n int) int {
    
    leftIndex := 1;
    rightIndex := n;
    check := 0; 
    
    if n <= 2{
        if isBadVersion(n-1){
            return n-1;
        }else{
            return n;
        }
    }
        
    for rightIndex != leftIndex{
        check = ((rightIndex -leftIndex) / 2) + leftIndex

        if isBadVersion(check){
            rightIndex = check;
        }else{
            leftIndex = check+1;
        }
    }
    return leftIndex;
}