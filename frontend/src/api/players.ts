import type { Edition } from "../config/editions";
import type { GoalScorerData } from "../config/goalscorer";
import { get } from "./client";

export function getGoalScorers(edition: Edition): Promise<GoalScorerData[]> {
  return get<GoalScorerData[]>(
    `/competitions/${edition.id}/${edition.startYear}/goal-scorers`,
  );
}
