import { useEffect, useState } from "react";
import { getAllMatches, getFeaturedMatches } from "../api/matches";
import { CompetitionMatchCard } from "../components/competition-match-card";
import { Header } from "../components/header";
import { LeagueHeader } from "../components/league-header";
import { Sidebar } from "../components/sidebar";
import { TopMatches } from "../components/top-matches";
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
  const [topMatches, setTopMatches] = useState<MatchData[]>([]);

  useEffect(() => {
    getAllMatches()
      .then(([upcoming, live, finished]) => {
        setTopMatches([...finished, ...live, ...upcoming]);
      })
      .catch(console.error);
  }, []);

  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px] gap-[20px]">
      <Header />
      <TopMatches matches={topMatches} />
      <div className="flex items-start">
        <Sidebar />
        <div className="flex flex-col gap-[15px] pt-[20px]">
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
