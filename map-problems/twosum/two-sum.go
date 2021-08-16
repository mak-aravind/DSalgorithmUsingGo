package twosum

func GetAddendIndics(target int, nums []int) (found bool, indices [2]int) {
	addendLookUpTable := make(map[int]int)
	for addendIndex, addendInEnquiry := range nums {
		if complementIndex, complementAvailable := addendLookUpTable[target-addendInEnquiry]; complementAvailable {
			indices[0] = complementIndex
			indices[1] = addendIndex
			found = true
			return
		}
		addendLookUpTable[addendInEnquiry] = addendIndex
	}
	return
}
