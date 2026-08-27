import type { ReactNode } from "react";
import type { MatchData } from "../config/matches";
import { Match } from "./match";


type CompetitionMatchCardProps = {
  header: ReactNode;
  matches: MatchData[];
};

export function CompetitionMatchCard({ header, matches}: CompetitionMatchCardProps) {
  return (
    <section className="flex w-[910px] flex-col items-center gap-[15px] rounded-lg bg-sections p-[10px]">
      {header}

      {matches.map((match) => (
        <Match key={match.id} {...match}
        />
      ))}
    </section>
  );
}
