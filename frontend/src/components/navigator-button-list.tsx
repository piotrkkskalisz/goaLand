import { NavigatorButton } from "./navigator-button";

const buttonNames = ["wyniki", "mecze", "tabela", "strzelcy", "asystenci"] as const;

type NavigatorButtonName = (typeof buttonNames)[number];

type NavigatorButtonListProps = {
  selectedButton: NavigatorButtonName;
};

export function NavigatorButtonList({
  selectedButton,
}: NavigatorButtonListProps) {
  return (
    <nav
      aria-label="Nawigacja rozgrywek"
      className="sticky top-[20px] flex w-[370px] flex-col items-center gap-[20px] pt-[15px] self-start"
    >
      {buttonNames.map((buttonName) => (
        <NavigatorButton
          isSelected={buttonName === selectedButton}
          key={buttonName}
          text={buttonName}
        />
      ))}
    </nav>
  );
}
