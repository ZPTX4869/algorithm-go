package searchs

var badVer int

func firstBadVersion(n int) int {
	left, right := 1, n+1

	for left < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) == true {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func setBadVersion(n int) {
	badVer = n
}

func isBadVersion(n int) bool {
	if n >= badVer {
		return true
	}

	return false
}
