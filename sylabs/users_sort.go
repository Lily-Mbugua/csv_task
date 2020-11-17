package sylabs

import (
	"sort"
)

// Sorts Method, calculates the frequency a user shows up first, then sorts by frequency count
// in accending order. It returns a new Users value that is sorted.
func (u *Users) Sort() Users {
	// lets calculate frequency before sorting by frequency
	// frequency is just how often a user shows up as a count
	for _, user := range *u {
		// bump the frequency when
		user.frequency++
	}

	// create a variable to store return
	usr := make(Users, len(*u))

	copy(usr,*u) // copy our users over

	// sort the users by frequncy and return this works because line 20 -22 works
	sort.Sort(sort.Reverse(usr))
	return usr
}

// implement the sort interface for the users, so we can sort by frequency
func (u Users) Len() int           { return len(u) }
func (u Users) Less(i, j int) bool { return u[i].frequency < u[j].frequency }
func (u Users) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

