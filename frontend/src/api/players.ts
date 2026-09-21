import type { Edition } from "../config/editions";
import type { GoalScorerData } from "../config/goalscorer";
import type { PlayersOnPosition } from "../config/team-players";
import { get } from "./client";


export function getGoalScorers(edition: Edition): Promise<GoalScorerData[]> {
  return get<GoalScorerData[]>(
    `/competitions/${edition.id}/${edition.startYear}/goal-scorers`,
  );
}


export function getTeamPlayers(
  competitionID: number,
  clubID: number,
): Promise<PlayersOnPosition[]> {
  return get<PlayersOnPosition[]>(
    `/competitions/${competitionID}/clubs/${clubID}/players`,
  );
}
