import type { GoalScorerData } from "../config/goalscorer";
import { GoalScorer} from "./goal-scorer";

type GoalScorersTableProps = {
  goalScorers: GoalScorerData[];
};

export function GoalScorersTable({ goalScorers }: GoalScorersTableProps) {
  return (
    <table className="text-primary bg-sections w-full table-fixed border-collapse overflow-hidden rounded-sm">
      <colgroup>
        <col className="w-[100px]" />
        <col />
        <col />
        <col className="w-[100px]" />
        <col className="w-[100px]" />
      </colgroup>

      <thead className=" text-granit-200">
        <tr className="h-[60px] text-center">
          <th className="font-normal">miejsce</th>
          <th className="font-norma text-left">zawodnik</th>
          <th className="font-norma text-left">klub</th>
          <th className="font-normal" >bramki</th>
          <th className="font-normal" >asysty</th>
        </tr>
      </thead>

      <tbody>
        {goalScorers.map((goalScorer) => (
          <GoalScorer
            key={`${goalScorer.position}-${goalScorer.playerID}`}
            goalScorer={goalScorer}
          />
        ))}
      </tbody>
    </table>
  );
}
