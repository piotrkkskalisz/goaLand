import { Outlet, useParams } from "react-router";
import { CompetitionTitle } from "../components/competition-title";
import { Header } from "../components/header";
import { NavigatorButtonList } from "../components/navigator-button-list";
import { getEditionFromStrings } from "../config/editions";


export function CompetitionLayout() {
  const { competitionID, startYear } = useParams();  
  const edition = getEditionFromStrings(competitionID, startYear)

  if (!edition) {
    return <div>Nie znaleziono rozgrywek</div>;
  }
  
  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="my-[20px] flex justify-center">
        <CompetitionTitle {...edition} />
      </div>
      <div className="flex items-start">
        <NavigatorButtonList />
        <Outlet context={edition}/>
      </div>
    </main>
  );
}
