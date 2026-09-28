import type { ReactNode } from "react";
import type { MatchData } from "../config/matches";
import { Match } from "./match";


type CompetitionMatchCardProps = {
  header: ReactNode;
  matches: MatchData[];
  isCurrent?: boolean;
};

export function CompetitionMatchCard({ header, matches, isCurrent = true }: CompetitionMatchCardProps) {
  return (
    <section className="flex w-[910px] flex-col items-center gap-[15px] rounded-lg bg-sections p-[10px]">
      {header}

      {matches.map((match) => (
        <Match key={match.id} {...match} isCurrent={isCurrent}
        />
      ))}
    </section>
  );
}
