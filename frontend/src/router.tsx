import { createBrowserRouter } from "react-router";
import { MainPage } from "./pages/main-page";
import { CompetitionResultsPage } from "./pages/competition-results-page";
import { CompetitionLayout } from "./layouts/competition-layout";
import { CompetitionMatchesPage } from "./pages/competition-matches-page";
import { CompetitionTablePage } from "./pages/competition-table-page";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: MainPage,
  },
  {
    path: "/:competitionID/:startYear/tabela",
    Component: CompetitionTablePage,
  },
  {
    path: "/:competitionID/:startYear",
    Component: CompetitionLayout,
    children: [
      {
        path: "wyniki",
        Component: CompetitionResultsPage,
      },
      {
        path: "mecze",
        Component: CompetitionMatchesPage,
      },
      {
        path: "strzelcy",
        Component: CompetitionResultsPage,
      },
    ],
  },
  {
    path: "*",
    element: <div>404 Not Found</div>,
  },
]);
