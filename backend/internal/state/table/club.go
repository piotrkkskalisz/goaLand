package table

type MatchResult string

const (
	Win  MatchResult = "win"
	Draw MatchResult = "draw"
	Loss MatchResult = "loss"

	FormMaxLength = 5
)

type Club struct {
	TeamID        int           `json:"teamId"`
	TeamName      string        `json:"teamName"`
	Points        int           `json:"points"`
	Wins          int           `json:"wins"`
	Draws         int           `json:"draws"`
	Losses        int           `json:"losses"`
	GoalsScored   int           `json:"goalsScored"`
	GoalsConceded int           `json:"goalsConceded"`
	Form          []MatchResult `json:"form"`

	IsLive bool `json:"isLive"`
}

func (c *Club) goalDiffrent() int {
	return c.GoalsScored - c.GoalsConceded
}

func (c *Club) addResult(teamGoals, rivalGoals int) {
	if teamGoals > rivalGoals {
		c.Wins++
	} else if teamGoals == rivalGoals {
		c.Draws++
	} else {
		c.Losses++
	}
}

func (c *Club) addResultToForm(teamGoals, rivalGoals int) {
	if len(c.Form) >= FormMaxLength {
		return
	}
	if teamGoals > rivalGoals {
		c.Form = append(c.Form, Win)
	} else if teamGoals == rivalGoals {
		c.Form = append(c.Form, Draw)
	} else {
		c.Form = append(c.Form, Loss)
	}
}
