import type { Edition } from "../config/editions";
import { get } from "./client";

export function GetCompetitionEditions(competitionId: number): Promise<Edition[]> {
  return get<Edition[]>(
    `/competitions/${competitionId}/editions`,
  );
}

export function GetActiveEdition(): Promise<Edition[]> {
  return get<Edition[]>(
    `/competitions`,
  );
}

export function GetCompetitionData(competitionID: number, startYear: number): Promise<Edition> {
  return get<Edition>(
    `/competitions/${competitionID}/${startYear}/data`,
  );
}