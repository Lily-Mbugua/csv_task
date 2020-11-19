package sylabs

import "testing"

func TestUserSort(t *testing.T) {

	expectedFrequency := 50

	users := Users{
		User{
			UserID:    "5c00a8f685db9ec46dbc13d5",
			frequency: 10,
		},

		User{
			UserID:    "5c00a8f685db9ec46dbc13d3",
			frequency: 20,
		},

		User{
			UserID:    "5c00a8f685db9ec46dbc13d0",
			frequency: 50,
		},
	}

	sorted := users.Sort()

	if sorted[0].frequency != expectedFrequency {
		t.Errorf("test failed; expected=%v, got=%v", expectedFrequency, sorted[0].frequency)
	}

}
