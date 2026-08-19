import { CompetitionMatchCard } from "../components/competition-match-card";
import { CompetitionTitle } from "../components/competition-title";
import { Header } from "../components/header";
import { MatchWeekHeader } from "../components/match-week-header";
import { NavigatorButtonList } from "../components/navigator-button-list";

export function CompetitionResultsPage() {
  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="my-[20px] flex justify-center">
        <CompetitionTitle leagueName="Premier League" season="26/27" />
      </div>
      <div className="flex items-start">
        <NavigatorButtonList selectedButton="wyniki" />
        <CompetitionMatchCard
          header={<MatchWeekHeader text="Kolejka 1" />}
          matches={[
            {
              status: "live",
              date: "15.08.2026",
              time: "17:30",
              homeTeam: "Liverpool",
              awayTeam: "Manchester City",
              homeScore: 2,
              awayScore: 1,
              minute: 82,
            },
            {
              status: "finished",
              date: "15.08.2026",
              time: "15:00",
              homeTeam: "Arsenal",
              awayTeam: "Chelsea",
              homeScore: 3,
              awayScore: 1,
            },
            {
              status: "finished",
              date: "15.08.2026",
              time: "15:00",
              homeTeam: "Manchester United",
              awayTeam: "Newcastle United",
              homeScore: 0,
              awayScore: 0,
            },
            {
              status: "finished",
              date: "15.08.2026",
              time: "13:30",
              homeTeam: "Tottenham",
              awayTeam: "Aston Villa",
              homeScore: 1,
              awayScore: 2,
            },
            {
              status: "finished",
              date: "14.08.2026",
              time: "20:00",
              homeTeam: "Everton",
              awayTeam: "West Ham United",
              homeScore: 2,
              awayScore: 0,
            },
          ]}
        />
      </div>
    </main>
  );
}
