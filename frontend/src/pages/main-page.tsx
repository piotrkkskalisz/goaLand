import { CompetitionMatchCard } from "../components/competition-match-card";
import { Header } from "../components/header";
import { LeagueHeader } from "../components/league-header";
import { Sidebar } from "../components/sidebar";

export function MainPage() {
  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="flex items-start">
        <Sidebar />
        <div className="flex flex-col gap-[15px] pt-[50px]">
          <CompetitionMatchCard
            header={
              <LeagueHeader
                leagueName="Premier League"
                flag="linear-gradient(90deg, transparent 42%, #ce1124 42% 58%, transparent 58%), linear-gradient(transparent 35%, #ce1124 35% 65%, transparent 65%), #fff"
              />
            }
            matches={[
              {
                status: "live",
                date: "13.07.2026",
                time: "18:00",
                homeTeam: "Arsenal",
                awayTeam: "Chelsea",
                homeScore: 0,
                awayScore: 0,
              },
              {
                status: "finished",
                date: "13.07.2026",
                time: "18:00",
                homeTeam: "Liverpool",
                awayTeam: "Everton",
                homeScore: 2,
                awayScore: 1,
              },
              {
                status: "scheduled",
                date: "13.07.2026",
                time: "18:00",
                homeTeam: "Manchester City",
                awayTeam: "Tottenham",
              },
            ]}
          />
          <CompetitionMatchCard
            header={
              <LeagueHeader
                leagueName="La Liga"
                flag="linear-gradient(#aa151b 0 25%, #f1bf00 25% 75%, #aa151b 75%)"
              />
            }
            matches={[
              {
                status: "finished",
                date: "14.07.2026",
                time: "18:00",
                homeTeam: "Real Madrid",
                awayTeam: "Barcelona",
                homeScore: 2,
                awayScore: 2,
              },
              {
                status: "scheduled",
                date: "14.07.2026",
                time: "20:30",
                homeTeam: "Atletico Madrid",
                awayTeam: "Sevilla",
              },
            ]}
          />
        </div>
      </div>
    </main>
  );
}
