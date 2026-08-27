import type { Edition } from "../config/editions";


export function CompetitionTitle(props: Edition) {
  return (
    <h1 className="w-fit rounded-lg bg-card px-[50px] py-[5px] text-heading uppercase">
      {props.competitionName} {props.startYear}/{props.startYear+1}
    </h1>
  );
}
