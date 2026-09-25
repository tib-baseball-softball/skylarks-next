import type { Toast } from "$lib/dp/types/Toast.ts";
import type { ClientResponseError } from "pocketbase";
import type { FieldValidationError } from "$lib/dp/types/Error";

class ToastController {
  public toastQueue: Toast[] = $state([]);
  public readonly toastDuration = 5000;

  public trigger(toast: Toast): void {
    // Generate a UUID if one is not provided
    const toastWithId = {
      ...toast,
      id: toast.id || crypto.randomUUID(),
    };

    if (!toast.background) {
      toast.background = "preset-filled-error-500";
    }

    this.toastQueue.push(toastWithId);

    setTimeout(() => {
      const index = this.toastQueue.findIndex((t) => t.id === toastWithId.id);
      if (index !== -1) {
        this.toastQueue.splice(index, 1);
      }
    }, this.toastDuration);
  }

  public clear(): void {
    this.toastQueue = [];
  }

  /**
   * Generic handler function that extracts validation errors from backend responses.
   * Prefer more specific handling in critical areas.
   *
   * @TODO: wire up with localization
   *
   * @param error the error instance thrown by the Pocketbase SDK
   */
  public handleClientResponseError(error: ClientResponseError): void {
    switch (error.status) {
      case 400:
        const fieldsWithErrors: [string, FieldValidationError][] =
          Object.entries(error.response?.data);

        // concrete field errors exist
        if (fieldsWithErrors?.length > 0) {
          for (const [key, innerError] of fieldsWithErrors) {
            this.trigger({
              message: `Validation error for field "${key}": ${innerError.message}`,
              background: "preset-filled-error-500",
            });
          }
          return;
        }

        // no additional data - display message directly
        this.trigger({
          message: error.message,
          background: "preset-filled-error-500",
        });

        break;
      case 403:
        this.trigger({
          message:
            "You don't have sufficient permissions to perform this action.",
          background: "preset-filled-error-500",
        });
        break;
      case 404:
        this.trigger({
          message: "The requested resource wasn't found.", // PB default message, but this way it gets translated
          background: "preset-filled-error-500",
        });
        break;
      case 500:
      case 501:
      case 502:
      case 503:
        this.trigger({
          message:
            "A server error occurred. This is most likely not user-actionable. Please try again later or contact support.",
          background: "preset-filled-error-500",
        });
        break;
      default:
        this.triggerGenericErrorMessage();
    }
  }

  public triggerGenericFormSuccessMessage(modelName: string): void {
    this.trigger({
      message: `${modelName} data saved successfully.`,
      background: "preset-filled-success-500",
    });
  }

  public triggerGenericFormErrorMessage(modelName: string): void {
    this.trigger({
      message: `An error occurred while saving ${modelName} data.`,
      background: "preset-filled-error-500",
    });
  }

  /**
   * Used for generic errors that could occur anywhere
   * and cannot be handled with the information at the call site.
   * Most likely server errors (500 range) or JS errors that have nothing to do with my code.
   */
  public triggerGenericErrorMessage(): void {
    this.trigger({
      message: "An unknown error occurred. Please try again later.",
      background: "preset-filled-error-500",
    });
  }

  /**
   * Generic enough, but still scoped to auth requests.
   * Should only used when no other helpful feedback can be given.
   */
  public triggerAuthErrorMessage(): void {
    this.trigger({
      message: "There was an error processing your authentication request.",
      background: "preset-filled-error-500",
    });
  }
}

export const toastController = new ToastController();
