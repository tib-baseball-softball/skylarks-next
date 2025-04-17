export type Handedness = "Left" | "Right" | "Switch" | "Unknown"

export function convertIntKeyToHandedness(key: number): Handedness {
  switch (key) {
    case 1:
      return "Left"
    case 2:
      return "Right"
    case 3:
      return "Switch"
  }

  return "Unknown"
}