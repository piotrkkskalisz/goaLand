import { useEffect, useState } from "react";
import { Link } from "react-router";
import { GetActiveEdition } from "../api/edition";
import { toFlag } from "../config/flags";
import type { Edition } from "../config/editions";

export function Sidebar() {
  const [editions, setEditions] = useState<Edition[]>([]);

  useEffect(() => {
    GetActiveEdition()
      .then(setEditions)
      .catch(console.error);
  }, []);

  return (
    <aside className="flex h-[650px] w-[370px] flex-col items-center gap-[30px] px-[10px] py-[50px]">
      {editions.map((edition) => (
        <Link
          className="flex h-[60px] w-[350px] items-center gap-[30px] rounded-lg px-[10px] text-left transition-colors hover:bg-granit-800"
          key={edition.id}
          to={`/${edition.id}/${edition.startYear}/wyniki`}
        >
          {toFlag(edition) && (
            <span
              aria-hidden="true"
              className="h-[40px] w-[60px]"
              style={{ background: toFlag(edition) }}
            />
          )}
          <span className="text-primary text-text-primary">
            {edition.competitionName}
          </span>
        </Link>
      ))}
    </aside>
  );
}
