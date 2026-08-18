import { Search } from "lucide-react";

export function Header() {
  return (
    <header className="flex h-[109px] w-full items-center justify-center gap-[50px] rounded-lg bg-header px-[77px] py-[25px]">
      <h1 className="w-[760px] text-center text-heading-bold">
        Goaland
      </h1>

      <label className="flex h-[60px] w-[300px] items-center gap-[10px] rounded-full bg-search px-[20px]">
        <Search aria-hidden="true" size={20} />
        <input
          aria-label="Wyszukaj"
          className="min-w-0 flex-1 bg-transparent text-secondary outline-none placeholder:text-text-secondary [&::-webkit-search-cancel-button]:hidden"
          placeholder="wyszukaj"
          type="search"
        />
      </label>
    </header>
  );
}
