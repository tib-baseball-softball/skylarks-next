import type {Toast} from "$lib/types/Toast.ts";

class ToastController {
  public toastQueue: Toast[] = $state([]);
  public readonly toastDuration = 5000;

  public trigger(toast: Toast) {
    // Generate a UUID if one is not provided
    const toastWithId = {
      ...toast,
      id: toast.id || crypto.randomUUID()
    };

    this.toastQueue.push(toastWithId);

    setTimeout(() => {
      const index = this.toastQueue.findIndex(t => t.id === toastWithId.id);
      if (index !== -1) {
        this.toastQueue.splice(index, 1);
      }
    }, this.toastDuration);
  }

  public clear() {
    this.toastQueue = [];
  }
}

export const toastController = new ToastController();
