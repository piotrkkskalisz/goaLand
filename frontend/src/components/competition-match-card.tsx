import type { ReactNode } from "react";
import { Match, type MatchProps } from "./match";

type CompetitionMatchCardProps = {
  header: ReactNode;
  matches: MatchProps[];
};

export function CompetitionMatchCard({ header, matches,
}: CompetitionMatchCardProps) {
  return (
    <section className="flex w-[910px] flex-col items-center gap-[15px] rounded-lg bg-sections p-[10px]">
      {header}

      {matches.map((match) => (
        <Match
          key={`${match.date}-${match.time}-${match.homeTeam}-${match.awayTeam}`}
          {...match}
        />
      ))}
    </section>
  );
}
