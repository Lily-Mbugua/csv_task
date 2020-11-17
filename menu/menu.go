package menu

import (
	"bufio"
	"csv_task/sylabs"
	"fmt"
	"time"
	"os"
	"strings"
)

func startPrompt(){
	fmt.Println("Please select from the follow, input number & hit return:")
	fmt.Println("1. Builds Executed within a Duration subtracted from current time.")
	fmt.Println("2. Builds Executed within Start_Time & End_Time range.")
	fmt.Println("3. Who are the top 5 Users and how many builds have they executed.")
	fmt.Println("4. Builds success rate, and for builds that are not succeeding what are the top exit codes.")
	fmt.Println("5. Exit.")
}

func Start(builds  *sylabs.SingularityBuilds) {
	startPrompt()
	scanner := bufio.NewScanner(os.Stdin)
	// Get User input, then invoke cases
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "1": // Menu Option 1
			buildsExecutedWithinDuration(builds)
			rePrompt()
			continue
		case "2": // Menu Option 2
			buildsExecutedWithinTime(builds)
			rePrompt()
			continue
		case "3": // Menu Option 3
			topFiveUser(builds)
			rePrompt()
			continue
		case "4": // Menu Option 4
			buildsSuccessRate(builds)
			rePrompt()
			continue
		case "5": // Menu Option 5
			exit()
		case "": // hidden restart
			startPrompt()
			continue
		default: // Incorrect Menu Option
			fmt.Printf("\n===================\n\nIncorrect Menu Option %q select, Restarting!\n\n===================\n\n\n\n", line)
			time.Sleep(2 * time.Second)
			continue
		}
	}
}


func buildsExecutedWithinDuration(builds *sylabs.SingularityBuilds) {
	fmt.Println("Please give me a duration of time:")
	// capture user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dur, err := time.ParseDuration(scanner.Text())
		if err != nil {
			fmt.Println("Invalid Duration:")
			fmt.Println(err)
			return
		}

		// print results
		resp := builds.BuildsExecutedSince(dur)
		if resp.Executed > 1 {
			fmt.Printf("Executed %d builds since %q\n", resp.Executed, time.Now().Add(-dur).Format(time.RFC3339))
			for _, build := range resp.Builds {
				fmt.Printf("build %q initiated by %q began %s and finished %s returned exit code %d and was  %d bytes large\n", build.UUID, build.UserID, build.BuildBeganTS.Format(time.RFC3339), build.BuildFinishedTS.Format(time.RFC3339), build.ExitCode, build.Size)
			}
		} else {
			fmt.Println("No Results!")
			return

		}

		return
	}

}

func buildsExecutedWithinTime(builds *sylabs.SingularityBuilds) {
	fmt.Println("Please give me a Start_Date and End_Date in  RFC3339 format separated by a | pipe character:")
	fmt.Println("Example: 2019-10-12T07:20:50.52Z | 2019-10-14T07:20:50.52Z")
	// capture user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dates := strings.Split(scanner.Text(), "|")
		if len(dates) != 2 {
			fmt.Println("Invalid characters given")
		}

		//parse dates
		startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(dates[0]))
		if err != nil {
			fmt.Println(err)
			return
		}
		endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(dates[1]))
		if err != nil {
			fmt.Println(err)
			return
		}

		// print results
		resp := builds.BuildsExcutedBetween(startDate, endDate)
		if resp.Executed > 1 {
			fmt.Printf("Executed %d builds since %q\n", resp.Executed, startDate.Format(time.RFC3339))
			for _, build := range resp.Builds {
				fmt.Printf("build %q initiated by %q began %s and finished %s returned exit code %d and was  %d bytes large\n", build.UUID, build.UserID, build.BuildBeganTS.Format(time.RFC3339), build.BuildFinishedTS.Format(time.RFC3339), build.ExitCode, build.Size)
			}
		} else {
			fmt.Println("No Results!")
		}
		return
	}

}

func topFiveUser(builds *sylabs.SingularityBuilds) {
	resp := builds.TopUsersBuilds(5)
	fmt.Println("LeaderBoard:")
	fmt.Printf("The 5 Top Users Executed %d builds \n", resp.Executed)
	for i, User := range resp.TopUsers {
		fmt.Printf("Rank %d User %q\n", i+1, User.UserID)
	}
	return
}

func buildsSuccessRate(builds *sylabs.SingularityBuilds) {

	resp := builds.BuildsSuccessRate()
	fmt.Printf("SuccessRate was %2.1f\n", resp.SuccessRate)
	if len(resp.TopFailureExitCodes) > 0 {
		fmt.Printf("TopExitCodes:%v\n", resp.TopFailureExitCodes)
	}
	return

}

// exit politely quits the program after telling user Goodbye
func exit() {
	fmt.Printf("Goodbye!\n\n")
	os.Exit(0)
}


// reRrun will start menu loop over again
func reRun(builds *sylabs.SingularityBuilds) {
	Start(builds)
}

// rePrompt() asks the user if they wish to continue
func rePrompt(){
	fmt.Println("\nPush Enter to continue....\n")
}