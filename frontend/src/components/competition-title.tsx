type CompetitionTitleProps = {
  leagueName: string;
  season: string;
};

export function CompetitionTitle({
  leagueName,
  season,
}: CompetitionTitleProps) {
  return (
    <h1 className="w-fit rounded-lg bg-card px-[50px] py-[5px] text-heading uppercase">
      {leagueName} {season}
    </h1>
  );
}
