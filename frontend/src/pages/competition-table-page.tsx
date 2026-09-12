import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { getTable } from "../api/table";
import { CompetitionTitle } from "../components/competition-title";
import { Header } from "../components/header";
import { Table } from "../components/table";
import type { Club } from "../config/club";
import { getEditionFromStrings } from "../config/editions";

export function CompetitionTablePage() {
  const { competitionID, startYear } = useParams();
  const edition = getEditionFromStrings(competitionID, startYear);
  const [clubs, setClubs] = useState<Club[]>([]);

  useEffect(() => {
    if (!edition) {
      return;
    }

    getTable(edition).then(setClubs).catch(console.error);
  }, [edition]);

  if (!edition) {
    return <div>Nie znaleziono rozgrywek</div>;
  }

  return (
    <main className="min-h-screen bg-dark-background px-[50px] pb-[50px] pt-[25px]">
      <Header />
      <div className="my-[20px] flex justify-center">
        <CompetitionTitle {...edition} />
      </div>
      <div className="flex justify-center">
        <Table clubs={clubs} />
      </div>
    </main>
  );
}
