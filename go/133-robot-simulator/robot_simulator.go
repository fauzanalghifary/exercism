package robot

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

// Step 2
// Define Action type here.
type Action struct {
	name  string
	robot Step2Robot
}

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		switch cmd {
		case 'R':
			action <- Action{name: "R"}
		case 'L':
			action <- Action{name: "L"}
		case 'A':
			action <- Action{name: "A"}
		}
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		switch act.name {
		case "R":
			robot.Dir = (robot.Dir + 1) % 4
		case "L":
			robot.Dir = (robot.Dir + 3) % 4
		case "A":
			newPos := robot.Pos
			switch robot.Dir {
			case N:
				newPos.Northing++
			case E:
				newPos.Easting++
			case S:
				newPos.Northing--
			case W:
				newPos.Easting--
			}
			// Check if new position is within room bounds
			if newPos.Easting >= extent.Min.Easting && newPos.Easting <= extent.Max.Easting &&
				newPos.Northing >= extent.Min.Northing && newPos.Northing <= extent.Max.Northing {
				robot.Pos = newPos
			}
		}
	}
	report <- robot
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	name  string
	robot string
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() {
		action <- Action3{name: "done", robot: name}
	}()

	for _, cmd := range script {
		switch cmd {
		case 'R', 'L', 'A':
			action <- Action3{name: string(cmd), robot: name}
		default:
			log <- name + ": undefined command " + string(cmd)
			return
		}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	// Validate initial robot positions
	robotMap := make(map[string]*Step3Robot)
	positions := make(map[Pos]string)
	activeRobots := make(map[string]bool)

	for i := range robots {
		robot := &robots[i]

		// Check for empty name
		if robot.Name == "" {
			log <- "Robot without a name"
			// Drain the action channel
			go func() {
				for range action {
				}
			}()
			rep <- robots
			return
		}

		// Check for duplicate names
		if _, exists := robotMap[robot.Name]; exists {
			log <- "Duplicate robot name: " + robot.Name
			// Drain the action channel
			go func() {
				for range action {
				}
			}()
			rep <- robots
			return
		}

		// Check if robot is outside room
		if robot.Pos.Easting < extent.Min.Easting || robot.Pos.Easting > extent.Max.Easting ||
			robot.Pos.Northing < extent.Min.Northing || robot.Pos.Northing > extent.Max.Northing {
			log <- robot.Name + " placed outside of room"
			// Drain the action channel
			go func() {
				for range action {
				}
			}()
			rep <- robots
			return
		}

		// Check for duplicate positions
		if otherName, exists := positions[robot.Pos]; exists {
			log <- robot.Name + " and " + otherName + " at same position"
			// Drain the action channel
			go func() {
				for range action {
				}
			}()
			rep <- robots
			return
		}

		robotMap[robot.Name] = robot
		positions[robot.Pos] = robot.Name
		activeRobots[robot.Name] = true
	}

	expectedRobots := len(robots)
	doneCount := 0

	// Process actions
	for act := range action {
		// Check for done signal
		if act.name == "done" {
			doneCount++
			if doneCount == expectedRobots {
				break
			}
			continue
		}

		// Check if action is from a valid robot
		if _, exists := robotMap[act.robot]; !exists {
			log <- "Action from unknown robot: " + act.robot
			continue
		}

		robot := robotMap[act.robot]

		switch act.name {
		case "R":
			robot.Dir = (robot.Dir + 1) % 4
		case "L":
			robot.Dir = (robot.Dir + 3) % 4
		case "A":
			newPos := robot.Pos
			switch robot.Dir {
			case N:
				newPos.Northing++
			case E:
				newPos.Easting++
			case S:
				newPos.Northing--
			case W:
				newPos.Easting--
			}

			// Check if new position is within room bounds
			if newPos.Easting < extent.Min.Easting || newPos.Easting > extent.Max.Easting ||
				newPos.Northing < extent.Min.Northing || newPos.Northing > extent.Max.Northing {
				log <- robot.Name + " attempting to advance into wall"
				continue
			}

			// Check if new position is occupied by another robot
			if otherName, occupied := positions[newPos]; occupied {
				log <- robot.Name + " attempting to advance into " + otherName
				continue
			}

			// Move robot
			delete(positions, robot.Pos)
			robot.Pos = newPos
			positions[newPos] = robot.Name
		}
	}

	rep <- robots
}
