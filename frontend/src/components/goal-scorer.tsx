import type { GoalScorerData } from "../config/goalscorer";

type GoalScorerProps = {
  goalScorer: GoalScorerData;
};

export function GoalScorer({ goalScorer }: GoalScorerProps) {
  return (
    <tr className="h-[58px] text-center">
      <td>{goalScorer.position}.</td>
      <td className="text-left">{goalScorer.playerName}</td>
      <td className="text-left">{goalScorer.teamName}</td>
      <td>{goalScorer.goals}</td>
      <td>{goalScorer.assists}</td>
    </tr>
  );
}
