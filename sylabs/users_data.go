package sylabs

// User describes a User's Unique Identifier and ID from CSV file
type User struct {
	UserID    string `json:"userid"` // User ID references the user that submitted the build request
	frequency int    `json:"-"`      // used to calculate rank
}

// Users represents all users from a singularityBuilds CSV file
type Users []User
