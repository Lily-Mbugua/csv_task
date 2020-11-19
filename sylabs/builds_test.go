package sylabs

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBuildsExecutedSince(t *testing.T) {
	// total number of builds expected
	expectedBuilds := 100

	// get path to current working directory
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get current path")
	}

	// create duration
	duration, err := time.ParseDuration("35000h")
	if err != nil {
		t.Errorf("failed to set duration")
	}

	// pass path to csv to Load
	singularityBuilds := Load(fmt.Sprintf("%s/test.csv", path))
	result := singularityBuilds.BuildsExecutedSince(duration)

	// validate response
	if result.Executed != expectedBuilds {
		t.Errorf("test failed; expected=%d, got=%d", expectedBuilds, result.Executed)
	}

}

func TestBuildsExecutedBetween(t *testing.T) {
	// total number of builds expected
	expectedBuilds := 35

	// get path to current working directory
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get current path")
	}

	startDate, err := time.Parse(time.RFC3339, strings.TrimSpace("2018-10-31T01:54:32-04:00"))
	if err != nil {
		fmt.Println(err)
		return
	}
	endDate, err := time.Parse(time.RFC3339, strings.TrimSpace("2018-11-09T00:18:59-05:00"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// pass path to csv to Load
	singularityBuilds := Load(fmt.Sprintf("%s/test.csv", path))
	result := singularityBuilds.BuildsExcutedBetween(startDate, endDate)

	// validate response
	if result.Executed != expectedBuilds {
		t.Errorf("test failed; expected=%d, got=%d", expectedBuilds, result.Executed)
	}

}

func TestTopUsersBuilds(t *testing.T) {
	expectedUserId := "5c00a8f685db9ec46dbc13d1"

	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get current path")
	}

	// pass path to csv to Load
	singularityBuilds := Load(fmt.Sprintf("%s/test.csv", path))
	result := singularityBuilds.TopUsersBuilds(5)

	highestUser := result.TopUsers[4]

	// validate response
	if highestUser.UserID != expectedUserId {
		t.Errorf("test failed; expected=%s, got=%s", expectedUserId, highestUser.UserID)
	}

}

func TestBuildsSuccessRate(t *testing.T) {
	expectedRate := 100.0

	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get current path")
	}

	// pass path to csv to Load
	singularityBuilds := Load(fmt.Sprintf("%s/test.csv", path))
	result := singularityBuilds.BuildsSuccessRate()

	// validate response
	if result.SuccessRate != expectedRate {
		t.Errorf("test failed; expected=%v, got=%v", expectedRate, result.SuccessRate)
	}

}
