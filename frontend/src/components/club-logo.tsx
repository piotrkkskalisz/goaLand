
type ClubLogoProps = {
  url?: string;
  size?: number;
};


export function ClubLogo({ url, size }: ClubLogoProps) {
  const logoSize = size ? `${size}px` : "full"
  return url && (
      <img
        className={`h-[${logoSize}] w-[${logoSize}] shrink-0 object-contain`}
        src={url}
        alt=""
        onError={(event) => event.currentTarget.remove()}
      />
  );
}

