import { get } from "./client";
import type { Edition } from "../config/editions";
import type { RoundMatches, MatchData} from "../config/matches";


export function getFeaturedMatches(edition: Edition) {
  return get<MatchData[]>(
    `/competitions/${edition.id}/${edition.startYear}/featured-matches`,
  );
}

export function getRoundMatches(edition: Edition) {
  return get<RoundMatches[]>(
    `/competitions/${edition.id}/${edition.startYear}/matches`,
  );
}

export function getRoundResults(edition: Edition): Promise<RoundMatches[]> {
  return get<RoundMatches[]>(
    `/competitions/${edition.id}/${edition.startYear}/results`,
  );
}
