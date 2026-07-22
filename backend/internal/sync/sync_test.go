package sync

import (
	"errors"
	"testing"
	"time"

	"backend/internal/api"
	"backend/internal/database"
	"backend/internal/sync/mocks"
	"backend/internal/testutils"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func strPtr(s string) *string { return &s }

func testSync(t *testing.T) (*gomock.Controller, *mocks.MockAPI, *mocks.MockDatabase, *Sync) {
	t.Helper()

	ctrl := gomock.NewController(t)

	apiMock := mocks.NewMockAPI(ctrl)
	dbMock := mocks.NewMockDatabase(ctrl)

	return ctrl, apiMock, dbMock, New(apiMock, dbMock)
}

func TestInitAreas(t *testing.T) {
	ctrl, apiMock, dbMock, s := testSync(t)
	defer ctrl.Finish()

	england := testutils.England()
	poland := testutils.Poland()

	apiMock.EXPECT().FetchAreas().Return([]api.Area{
		{ID: testutils.PolandAreaID, Name: "Poland", CountryCode: testutils.PremierLeagueCode, ParentArea: strPtr("Europe")},
		{ID: testutils.EnglandAreaID, Name: testutils.EnglandName, CountryCode: "EN", ParentArea: strPtr("Europe")},
	}, nil)

	dbMock.EXPECT().SaveAreas([]database.Area{
		poland,
		england,
	}).Return(nil)

	require.NoError(t, s.initAreas())

	require.Equal(t, s.areasByName["Poland"], testutils.PolandAreaID)
	require.Equal(t, s.areasByName[testutils.EnglandName], testutils.EnglandAreaID)
}

func TestInitCompetition(t *testing.T) {
	ctrl, apiMock, dbMock, s := testSync(t)
	defer ctrl.Finish()

	apiMock.EXPECT().FetchCompetition(testutils.PremierLeagueCode).Return(api.Competition{
		ID:   2021,
		Name: "Premier League",
		Code: testutils.PremierLeagueCode,
		Type: "LEAGUE",
		Area: struct {
			ID int `json:"id"`
		}{ID: testutils.EnglandAreaID},
	}, nil)

	s.areasByName[testutils.EnglandName] = testutils.EnglandAreaID

	competition := testutils.NewCompetition()

	dbMock.EXPECT().SaveCompetitions([]database.Competition{competition}).Return(nil)

	id, err := s.initCompetition(competition.Code)
	require.NoError(t, err)

	require.Equal(t, id, competition.CompetitionID)
}

func TestInitCompetitionFetchError(t *testing.T) {
	ctrl, apiMock, _, s := testSync(t)
	defer ctrl.Finish()

	expectedErr := errors.New("api error")

	apiMock.EXPECT().
		FetchCompetition(testutils.PremierLeagueCode).
		Return(api.Competition{}, expectedErr)

	_, err := s.initCompetition(testutils.PremierLeagueCode)

	require.ErrorIs(t, err, expectedErr)
}

func TestInitCompetitions(t *testing.T) {
	ctrl, apiMock, dbMock, s := testSync(t)
	defer ctrl.Finish()

	premierLeague := testutils.PremierLeague()
	ekstraklasa := testutils.Ekstraklasa()

	apiMock.EXPECT().FetchCompetitions().Return([]api.Competition{
		{
			ID:   premierLeague.CompetitionID,
			Name: premierLeague.Name,
			Code: premierLeague.Code,
			Type: premierLeague.CompetitionType,
			Area: struct {
				ID int `json:"id"`
			}{
				ID: premierLeague.AreaID,
			},
		}, {
			ID:   ekstraklasa.CompetitionID,
			Name: ekstraklasa.Name,
			Code: ekstraklasa.Code,
			Type: ekstraklasa.CompetitionType,
			Area: struct {
				ID int `json:"id"`
			}{
				ID: ekstraklasa.AreaID,
			},
		},
	}, nil)

	dbMock.EXPECT().
		SaveCompetitions([]database.Competition{premierLeague}).
		Return(nil)

	err := s.initCompetitions(map[string]struct{}{
		premierLeague.Code: {},
	})
	require.NoError(t, err)
}

func TestInitCompetitionsMissingCodes(t *testing.T) {
	ctrl, apiMock, _, s := testSync(t)
	defer ctrl.Finish()

	apiMock.EXPECT().FetchCompetitions().Return([]api.Competition{}, nil)

	err := s.initCompetitions(map[string]struct{}{
		testutils.PremierLeagueCode: {},
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "competitions not found")
}

func TestInitMatches(t *testing.T) {
	ctrl, apiMock, dbMock, s := testSync(t)
	defer ctrl.Finish()

	match := testutils.NewMatch()

	apiMatch := api.Match{
		ID:      match.MatchID,
		UtcDate: match.StartTime.Format(time.RFC3339),
		Status:  match.Status,
	}
	apiMatch.HomeTeam.ID = match.HomeTeamID
	apiMatch.AwayTeam.ID = match.AwayTeamID

	apiMatch.Score.FullTime.Home = match.HomeGoals
	apiMatch.Score.FullTime.Away = match.AwayGoals

	apiMock.EXPECT().FetchMatches(testutils.PremierLeagueCode, testutils.Year).Return([]api.Match{apiMatch}, nil)

	dbMock.EXPECT().SaveMatches([]database.Match{match}).Return(nil)

	err := s.initMatches(Season{
		CompetitionID:   testutils.PremierLeagueID,
		CompetitionCode: testutils.PremierLeagueCode,
		StartYear:       testutils.Year,
	})

	require.NoError(t, err)
}

func TestInitGoalScorers(t *testing.T) {
	ctrl, apiMock, dbMock, s := testSync(t)
	defer ctrl.Finish()

	scorer := testutils.NewGoalScorer()

	assists := scorer.Assists

	apiScorer := api.GoalScorer{
		Goals:     scorer.Goals,
		Assists:   &assists,
		Penalties: nil,
	}

	apiScorer.Player.ID = scorer.GoalScorerID
	apiScorer.Player.Name = scorer.Name
	apiScorer.Player.Nationality = testutils.EnglandName

	apiScorer.Team.ID = scorer.TeamID

	s.areasByName[testutils.EnglandName] = testutils.EnglandAreaID

	expected := scorer

	apiMock.EXPECT().
		FetchGoalScorers(testutils.PremierLeagueCode, testutils.Year, defaultGoalScorerLimit).
		Return([]api.GoalScorer{apiScorer}, nil)

	dbMock.EXPECT().
		SaveGoalScorers([]database.GoalScorer{expected}).
		Return(nil)

	err := s.initGoalScorers(Season{
		CompetitionID:   testutils.PremierLeagueID,
		CompetitionCode: testutils.PremierLeagueCode,
		StartYear:       testutils.Year,
	}, defaultGoalScorerLimit)

	require.NoError(t, err)
}
