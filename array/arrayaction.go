package main

func pivotIndex(nums []int) int {
    
    length := len(nums)
    
    if length <= 2 {
        return -1
    }
    
    lsum, rsum := 0, 0
    
    lidx,ridx := 0, length -1
    
    for i:=0;ridx != lidx;i++{
        
        if lsum > rsum {
            rsum += nums[ridx]
            ridx -= 1
        }else{
            lsum += nums[lidx]
            lidx ++
        }
	}

	if lsum == rsum{
		return lidx
	}
	
	return -1

}