package status

const (
	// API response
	RetCode400 = "Unauthorized Request"
	RetCode401 = "Invalid Request"
	RetCode404 = "Bad Request"
	RetCode419 = "Authentication Timeout"
	RetCode500 = "Internal Server Error"

	// User status
	UserStatusNew    = "New"
	UserStatusActive = "Active"
	UserStatusLocked = "Locked"
)
