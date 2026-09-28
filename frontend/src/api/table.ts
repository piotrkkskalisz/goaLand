import type { ClubStats } from "../config/club";
import type { Edition } from "../config/editions";
import { get } from "./client";

export function getTable(edition: Edition): Promise<ClubStats[]> {
  return get<ClubStats[]>(
    `/competitions/${edition.id}/${edition.startYear}/table`,
  );
}
