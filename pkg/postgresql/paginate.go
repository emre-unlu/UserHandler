package postgresql

func CalculateOffset(page uint, limit uint) int {
	return int((page - 1) * limit)
}
