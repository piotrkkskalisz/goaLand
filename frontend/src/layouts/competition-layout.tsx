import { Outlet } from "react-router";
import { CompetitionTitle } from "../components/competition-title";
import { Header } from "../components/header";
import { NavigatorButtonList } from "../components/navigator-button-list";

export function CompetitionLayout() {
  return (
    <main className="min-h-screen bg-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="my-[20px] flex justify-center">
        <CompetitionTitle leagueName="Premier League" season="26/27" />
      </div>
      <div className="flex items-start">
        <NavigatorButtonList />
        <Outlet />
      </div>
    </main>
  );
}
