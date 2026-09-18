export type Toast = {
  id?: string; // must be unique, will be generated if not provided
  message: string;
  background?:
    | "preset-filled-success-500"
    | "preset-filled-warning-500"
    | "preset-filled-error-500"
};
