package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamStats struct {
	MP, W, D, L, P int
}

type Team struct {
	Name string
	*TeamStats
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*TeamStats)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		matches := strings.Split(line, ";")

		if line == "" || string(line[0]) == "#" {
			continue
		}

		if len(matches) != 3 {
			return errors.New("invalid input")
		}

		if teams[matches[0]] == nil {
			teams[matches[0]] = &TeamStats{}
		}
		if teams[matches[1]] == nil {
			teams[matches[1]] = &TeamStats{}
		}

		teams[matches[0]].MP++
		teams[matches[1]].MP++
		result := matches[2]
		if result == "win" {
			teams[matches[0]].W++
			teams[matches[0]].P += 3
			teams[matches[1]].L++
		} else if result == "loss" {
			teams[matches[1]].W++
			teams[matches[1]].P += 3
			teams[matches[0]].L++
		} else if result == "draw" {
			teams[matches[0]].D++
			teams[matches[0]].P++
			teams[matches[1]].D++
			teams[matches[1]].P++
		} else {
			return errors.New("invalid input")
		}
	}

	var sortedTeams []Team
	for name, stats := range teams {
		sortedTeams = append(sortedTeams, Team{name, stats})
	}

	sort.Slice(
		sortedTeams, func(i, j int) bool {
			if sortedTeams[i].P == sortedTeams[j].P {
				return sortedTeams[i].Name < sortedTeams[j].Name
			}
			return sortedTeams[i].P > sortedTeams[j].P
		},
	)

	_, err := fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range sortedTeams {
		_, err = fmt.Fprintf(
			writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.Name, team.MP, team.W, team.D, team.L, team.P,
		)
	}

	if err != nil {
		return err
	}

	return nil
}
