export type FieldValidationError = {
  code?: string;
  message?: string;
};

export enum PBErrorCode {
  ValidationNotUnique = "validation_not_unique",
  SignupKeyInvalid = "signup_key_invalid",
}
