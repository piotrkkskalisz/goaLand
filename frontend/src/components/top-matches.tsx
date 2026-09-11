import { ChevronLeft, ChevronRight } from "lucide-react";
import { useEffect, useRef } from "react";
import type { MatchData } from "../config/matches";
import { FinishedMatch, LiveMatch, UpcomingMatch } from "./top-match";

type TopMatchesProps = {
  matches: MatchData[];
};

function TopMatchItem({ match }: { match: MatchData }) {
  if (match.status === "scheduled") {
    return ( <UpcomingMatch {...match}/> );
  }

  if (match.status === "live") {
    return ( <LiveMatch {...match}/> );
  }
  return ( <FinishedMatch {...match}/> );
}

export function TopMatches({ matches }: TopMatchesProps) {
  const matchesRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const currentIndex = matches.findIndex(
      (match) => match.status === "live" || match.status === "scheduled",
    );
    matchesRef.current?.scrollTo({ left: Math.max(currentIndex, 0) * 138 });
  }, [matches]);

  function scroll(direction: -1 | 1) {
    matchesRef.current?.scrollBy({
      left: direction * 138,
      behavior: "smooth",
    });
  }

  return (
    <div className="flex w-full items-center gap-[10px] px-[10px] py-[15px]">
      <button
        type="button"
        aria-label="Poprzednie mecze"
        className="shrink-0 cursor-pointer text-text hover:text-green-500"
        onClick={() => scroll(-1)}
      >
        <ChevronLeft size={20} strokeWidth={2.5} />
      </button>

      <div
        ref={matchesRef}
        className="flex flex-1 items-center gap-[10px] overflow-hidden"
      >
        {matches.map((match) => (
          <TopMatchItem key={match.id} match={match} />
        ))}
      </div>

      <button
        type="button"
        aria-label="Następne mecze"
        className="shrink-0 cursor-pointer text-text hover:text-green-500"
        onClick={() => scroll(1)}
      >
        <ChevronRight size={20} strokeWidth={2.5} />
      </button>
    </div>
  );
}
