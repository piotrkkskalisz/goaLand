import type { Club } from "../config/club";
import type { Edition } from "../config/editions";
import { get } from "./client";

export function getTable(edition: Edition): Promise<Club[]> {
  return get<Club[]>(
    `/competitions/${edition.id}/${edition.startYear}/table`,
  );
}
