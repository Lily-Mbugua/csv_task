package sylabs

import (
	"time"
)

func (s *SingularityBuilds) BuildsExecutedSince(duration time.Duration) BuildsExecutedResponse {
	requestedTime := time.Now().Add(-duration)
	var response BuildsExecutedResponse
	for _, build := range *s {
		if build.BuildBeganTS.After(requestedTime) {
			response.Executed++
			response.Builds = append(response.Builds, build)
		}
	}
	return response
}

// BuildsExcutedBetween will return a BuildsExecutedResponse representing the total builds that where executed
// since the duration of time specified by the startTime and endTime
func (s *SingularityBuilds) BuildsExcutedBetween(startTime, endTime time.Time) BuildsExecutedResponse {
	// our returned response
	var response BuildsExecutedResponse
	// lets iterate through all singularityBuilds to calculate BuildExecutionResponse
	for _, build := range *s {
		// if the build is within the window of our start/endtime it's valid
		if build.BuildBeganTS.After(startTime) && build.BuildFinishedTS.Before(endTime) {
			// valid execution increment
			response.Executed++
			// add current build to response builds
			response.Builds = append(response.Builds, build)
		}
	}
	return response
}


func (s *SingularityBuilds) TopUsersBuilds(topFilter int) BuildTopUsers {
	//Map used to calculate frequency of each user in Singularity Builds
	usrMap := make(map[string]int)
	//calculate the frequencies
	for _, build := range *s {
		//bump if it exists
		usrMap[build.UserID]++
	}

	// will contain all of our users
	var usrs Users
	// Convert from map to slice of User
	for k, v := range usrMap {
		usrs = append(usrs, User{
			UserID:    k,
			frequency: v,
		})
	}

	// Sort the users by frequency
	usrs = usrs.Sort()

	// limit top users by topFilter
	usrs = usrs[:topFilter]

	//Map  to store our top users
	topUsrMap := make(map[string]struct{})
	for _, build := range *s {
		//it exists
		topUsrMap[build.UserID] = struct{}{}
	}

	// gather our users builds and add to builds variable
	response := BuildTopUsers{
		TopUsers: usrs, // we already calculated this, assign it to top users
	}
	// calculate executed and builds
	for _, build := range *s {
		// if the current build is among our top
		if _, ok := topUsrMap[build.UserID]; ok {
			// bump executed b/c it was done by a Top User
			response.Executed++
			// build our response builds, add current build to response builds
			response.Builds = append(response.Builds, build)
		}
	}
	return response
}

// BuildsSuccessRate will calculate the percentage of builds that have succeed and the top exit codes for those that have failed
func (s *SingularityBuilds) BuildsSuccessRate() BuildSuccessRateResponse {
	// Total builds is the number of builds that exist in the SingularityBuilds
	totalBuilds := len(*s)
	// how many failed builds did we have?
	var failedBuilds int

	response := BuildSuccessRateResponse{}
	// iterate through all builds so we can calculate successfulBuilds, failedBuilds and gather ExitCodes if there are any
	for _, build := range *s {
		//if build ExitCode is greater than 0, that implies failure
		if build.ExitCode > 0 {
			// lets capture the failure
			failedBuilds++
			// add exit code to TopFailureExitCodes
			response.TopFailureExitCodes = append(response.TopFailureExitCodes, build.ExitCode )
		}

	}
	// calculate SuccessRate
	successfulBuilds := float64(totalBuilds - failedBuilds)
	response.SuccessRate = successfulBuilds / float64(totalBuilds) * 100.00
	return response
}
