import type { Competition} from "./editions";
import {
  premierLeagueCompetition,
  laLigaCompetition,
  bundesligaCompetition,
  serieACompetition,
  ligue1Competition,
} from "./editions";


const flags = new Map<number, string>([
  [
    premierLeagueCompetition.id,
    "linear-gradient(90deg, transparent 42%, #ce1124 42% 58%, transparent 58%), linear-gradient(transparent 35%, #ce1124 35% 65%, transparent 65%), #fff",
  ], [
    laLigaCompetition.id,
    "linear-gradient(#aa151b 0 25%, #f1bf00 25% 75%, #aa151b 75%)",
  ], [
    serieACompetition.id,
    "linear-gradient(90deg, #009246 0 33%, #fff 33% 67%, #ce2b37 67%)",
  ], [
    bundesligaCompetition.id,
    "linear-gradient(#000 0 33%, #dd0000 33% 67%, #ffce00 67%)",
  ],[
    ligue1Competition.id,
    "linear-gradient(90deg, #002654 0 33%, #fff 33% 67%, #ce1126 67%)",
  ],
]);

export function toFlag(competition: Competition) {
  if (flags.has(competition.id)) {
    return flags.get(competition.id);
  }
  return "linear-gradient(#808080, #808080)";
}
