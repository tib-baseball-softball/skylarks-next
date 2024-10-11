/** @deprecated use enum BaseballFieldPosition instead */
export type BaseballPosition = "Pitcher" | "Catcher" | "First Base" | "Second Base" | "Third Base" | "Shortstop" |
  "Left Field" | "Right Field" | "Center Field" | "Utility" | "Infield" | "Outfield"

export enum BaseballFieldPosition {
  None = 0,
  Pitcher = 1,
  Catcher = 2,
  FirstBase = 3,
  SecondBase = 4,
  ThirdBase = 5,
  Shortstop = 6,
  LeftField = 7,
  CenterField = 8,
  RightField = 9,
  Outfield = 10,
  Infield = 11,
  Utility = 12,
  DesignatedHitter = 13,
  DesignatedPlayer = 14
}

export function getSingleBaseballPositionLabel(value: BaseballFieldPosition): string {
  switch (value) {
    case BaseballFieldPosition.None:
      return 'None';
    case BaseballFieldPosition.Pitcher:
      return 'Pitcher';
    case BaseballFieldPosition.Catcher:
      return 'Catcher';
    case BaseballFieldPosition.FirstBase:
      return 'First Base';
    case BaseballFieldPosition.SecondBase:
      return 'Second Base';
    case BaseballFieldPosition.ThirdBase:
      return 'Third Base';
    case BaseballFieldPosition.Shortstop:
      return 'Shortstop';
    case BaseballFieldPosition.LeftField:
      return 'Left Field';
    case BaseballFieldPosition.CenterField:
      return 'Center Field';
    case BaseballFieldPosition.RightField:
      return 'Right Field';
    case BaseballFieldPosition.Outfield:
      return 'Outfield';
    case BaseballFieldPosition.Infield:
      return 'Infield';
    case BaseballFieldPosition.Utility:
      return 'Utility';
    case BaseballFieldPosition.DesignatedHitter:
      return 'Designated Hitter';
    case BaseballFieldPosition.DesignatedPlayer:
      return 'Designated Player';
    default:
      return 'Unknown';
  }
}

export function getAllBaseballPositionStringValues(): string[] {
  return Object.values(BaseballFieldPosition)
    .filter(value => typeof value === 'string') as string[];
}

export function positionKeysToEnumStringValues(keys: string[]): string[] {
  return keys.map(key => BaseballFieldPosition[parseInt(key)]);
}

export function positionEnumStringValuesToKeys(values: string[]): string[] {
  return values.map(value => {
    //@ts-ignore
    const key = Object.keys(BaseballFieldPosition).find(k => BaseballFieldPosition[k] === value);
    return key ? parseInt(key).toString() : "Unknown";
  });
}
