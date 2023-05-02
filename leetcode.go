package leetcode

import (
"fmt"
)

func maximumDifference(nums []int) int {
    arrLen := len(nums)
    maxDiff := -1

    indexOfMinUpTo := make([]int, arrLen)
    indexOfMaxUpTo := make([]int, arrLen)

    indexOfMinUpTo[0] = 0
    indexOfMaxUpTo[arrLen-1] = arrLen-1

    for i := 1; i < arrLen; i++ {
        if nums[i] < nums[indexOfMinUpTo[i-1]] {
            indexOfMinUpTo[i] = i
        } else {
            indexOfMinUpTo[i] = indexOfMinUpTo[i-1]
        }
    }

    for i := arrLen-2; i >= 0; i-- {
        if nums[i] > nums[indexOfMaxUpTo[i+1]] {
            indexOfMaxUpTo[i] = i
        } else {
            indexOfMaxUpTo[i] = indexOfMaxUpTo[i+1]
        }
    }

    for i := 0; i < arrLen; i++ {
        if( indexOfMaxUpTo[i] == indexOfMinUpTo[i] ) { continue }

        currentDiff := nums[indexOfMaxUpTo[i]] - nums[indexOfMinUpTo[i]]

        if currentDiff == 0  { continue }

        if currentDiff > maxDiff {
            maxDiff = currentDiff
        }
    }

    return maxDiff
}