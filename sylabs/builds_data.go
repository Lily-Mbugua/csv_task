package sylabs

import "time"

// SingularityBuild represents /each line on the csv file
type SingularityBuild struct {
	UUID      string
	User
	RequestReceviedTS *time.Time
	BuildBeganTS      *time.Time
	BuildFinishedTS   *time.Time
	Deleted           bool
	ExitCode          uint8
	Size              uint64
}

// SingularityBuilds represents all builds from  singularityBuilds CSV file
type SingularityBuilds []SingularityBuild

// BuildsExecutedResponse represents responses from marketing departments Builds Executed Duration request
type BuildsExecutedResponse struct {
	Executed int               `json:"executed"`
	Builds   SingularityBuilds `json:"builds"`
}

// BuildTopUsers represents the top users of a singularity build, and their builds
type BuildTopUsers struct {
	TopUsers Users `json:"topUsers"`
	Executed int               `json:"executed"`
	Builds   SingularityBuilds `json:"builds"`
}
// BuildSuccessRateResponse represents a data structure containing the successRate as percentage and
// top failure exit codes as slice of string
type BuildSuccessRateResponse struct {
	SuccessRate         float64  `json:"successRate"`
	TopFailureExitCodes []uint8 `json:"topFailureExitCodes"`
}
