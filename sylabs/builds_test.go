package sylabs

import (
	"fmt"
	"os"
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
