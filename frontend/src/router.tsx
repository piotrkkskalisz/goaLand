import { createBrowserRouter } from "react-router";
import { MainPage } from "./pages/main-page";
import { CompetitionResultsPage } from "./pages/competition-results-page";
import { CompetitionLayout } from "./layouts/competition-layout";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: MainPage,
  },
  {
    path: "/:competitionCode/:startYear",
    Component: CompetitionLayout,
    children: [
      {
        path: "wyniki",
        Component: CompetitionResultsPage,
      },
      {
        path: "mecze",
        Component: CompetitionResultsPage,
      },
      {
        path: "tabela",
        Component: CompetitionResultsPage,
      },
      {
        path: "strzelcy",
        Component: CompetitionResultsPage,
      },
      {
        path: "asystenci",
        Component: CompetitionResultsPage,
      },
    ],
  },
  {
    path: "*",
    element: <div>404 Not Found</div>,
  },
]);
