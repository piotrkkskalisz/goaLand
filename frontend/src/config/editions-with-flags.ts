import type { Edition } from "./editions";
import {
  bundesligaEdition,
  laLigaEdition,
  ligue1Edition,
  premierLeagueEdition,
  serieAEdition,
} from "./editions";


const flags = new Map<number, string>([
  [
    premierLeagueEdition.id,
    "linear-gradient(90deg, transparent 42%, #ce1124 42% 58%, transparent 58%), linear-gradient(transparent 35%, #ce1124 35% 65%, transparent 65%), #fff",
  ], [
    laLigaEdition.id,
    "linear-gradient(#aa151b 0 25%, #f1bf00 25% 75%, #aa151b 75%)",
  ], [
    serieAEdition.id,
    "linear-gradient(90deg, #009246 0 33%, #fff 33% 67%, #ce2b37 67%)",
  ], [
    bundesligaEdition.id,
    "linear-gradient(#000 0 33%, #dd0000 33% 67%, #ffce00 67%)",
  ],[
    ligue1Edition.id,
    "linear-gradient(90deg, #002654 0 33%, #fff 33% 67%, #ce1126 67%)",
  ],
]);

export function toFlag(edition: Edition) {
  return flags.get(edition.id);
}
