import { Match, type MatchProps } from "./match";

type CompetitionMatchCardProps = {
  leagueName: string;
  flag: string;
  matches: MatchProps[];
};

export function CompetitionMatchCard({ leagueName, flag, matches,
}: CompetitionMatchCardProps) {
  return (
    <section className="flex w-[910px] flex-col items-center gap-[15px] rounded-lg bg-sections p-[10px]">
      <header className="flex h-[40px] w-[800px] items-center gap-[30px] rounded-lg bg-green-750 pl-[15px] pr-[10px]">
        <span
          aria-hidden="true"
          className="h-[25px] w-[35px]"
          style={{ background: flag }}
        />
        <span>{leagueName}</span>
      </header>

      {matches.map((match) => (
        <Match
          key={`${match.date}-${match.time}-${match.homeTeam}-${match.awayTeam}`}
          {...match}
        />
      ))}
    </section>
  );
}
