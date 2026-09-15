import { useEffect, useState } from "react";
import { useOutletContext } from "react-router";
import { getGoalScorers } from "../api/players";
import { GoalScorersTable } from "../components/goal-scorers-table";
import type { Edition } from "../config/editions";
import type { GoalScorerData } from "../config/goalscorer";

export function CompetitionGoalScorersPage() {
  const [goalScorers, setGoalScorers] = useState<GoalScorerData[]>([]);
  const edition = useOutletContext<Edition>();

  useEffect(() => {
    getGoalScorers(edition).then(setGoalScorers).catch(console.error);
  }, [edition]);

  return (
    <div className="pt-[50px] w-[1210px]">
      <GoalScorersTable goalScorers={goalScorers} />
    </div>
  );
}
