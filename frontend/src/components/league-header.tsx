import type { Edition } from "../config/editions";
import { toFlag } from "../config/editions-with-flags";

export function LeagueHeader(props: Edition) {
  return (
    <header className="flex h-[40px] w-[800px] items-center gap-[30px] rounded-lg bg-green-750  px-[30px] ">
      <span
        aria-hidden="true"
        className="h-[25px] w-[35px]"
        style={{ background: toFlag(props) }}
      />
      <span>{props.competitionName}</span>
    </header>
  );
}
