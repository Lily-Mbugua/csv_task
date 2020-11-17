package sylabs

// TopUsers will return which users are using the remote build service the most
// Sorted by acceding frequency
func (u *Users) TopUsers() Users {
	sortedUsrs := u.Sort()
	return sortedUsrs
}


