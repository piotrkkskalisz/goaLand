import { useEffect, useState } from "react";
import { getFeaturedMatches } from "../api/matches";
import { CompetitionMatchCard } from "../components/competition-match-card";
import { Header } from "../components/header";
import { LeagueHeader } from "../components/league-header";
import { Sidebar } from "../components/sidebar";
import { mainEditions, type Edition } from "../config/editions";
import type { MatchData } from "../config/matches";

function FeaturedMatchesCard({ edition }: { edition: Edition }) {
  const [matches, setMatches] = useState<MatchData[]>([]);

  useEffect(() => {
    getFeaturedMatches(edition)
      .then(setMatches)
      .catch(console.error);
  }, [edition]);

  return (
    <CompetitionMatchCard
      header={<LeagueHeader {...edition} />}
      matches={matches}
    />
  );
}

export function MainPage() {
  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="flex items-start">
        <Sidebar />
        <div className="flex flex-col gap-[15px] pt-[50px]">
          {mainEditions.map((edition) => (
            <FeaturedMatchesCard
              key={`${edition.id}-${edition.startYear}`}
              edition={edition}
            />
          ))}
        </div>
      </div>
    </main>
  );
}