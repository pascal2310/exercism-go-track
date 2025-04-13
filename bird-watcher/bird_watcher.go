package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds = 0
	for i := range birdsPerDay {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var (
		birdCount  = 0
		startRange = (week - 1) * 7
		endRange   = week * 7
	)

	for i := startRange; i < endRange; i++ {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
