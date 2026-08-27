import { getRoundMatches} from "../api/matches";
import { CompetitionRoundMatchesPage } from "./competition-round-matches-page";

export function CompetitionMatchesPage(){
    return < CompetitionRoundMatchesPage
      fetchMatches={getRoundMatches}
    />
}