package twosum

func GetAddendIndics(target int, nums []int) (found bool, indices [2]int) {
	addendLookUpTable := make(map[int]int)
	for enquiredAddendIndex, addendToEnquire := range nums {
		if complementIndex, complementAvailable := addendLookUpTable[target-addendToEnquire]; complementAvailable {
			indices[0] = complementIndex
			indices[1] = enquiredAddendIndex
			found = true
			return
		}
		addendLookUpTable[addendToEnquire] = enquiredAddendIndex
	}
	return
}
