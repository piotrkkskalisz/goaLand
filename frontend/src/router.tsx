import { createBrowserRouter } from "react-router";
import { MainPage } from "./pages/main-page";
import { CompetitionResultsPage } from "./pages/competition-results-page";
import { CompetitionLayout } from "./layouts/competition-layout";
import { CompetitionMatchesPage } from "./pages/competition-matches-page";
import { CompetitionTablePage } from "./pages/competition-table-page";
import { CompetitionGoalScorersPage } from "./pages/competition-goal-scorers-page";
import { basePath, routePath, routes } from "./config/routes";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: MainPage,
  },
  {
    path: routePath(routes.table),
    Component: CompetitionTablePage,
  },
  {
    path: basePath,
    Component: CompetitionLayout,
    children: [
      { 
        path: routes.results,
        Component: CompetitionResultsPage,
      },
      {
        path: routes.matches,
        Component: CompetitionMatchesPage,
      },
      {
        path: routes.goalScorers,
        Component: CompetitionGoalScorersPage,
      },
    ],
  },
  {
    path: "*",
    element: <div>404 Not Found</div>,
  },
]);
