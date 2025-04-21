import type {Toast} from "$lib/types/Toast.ts";

class ToastController {
  public toasts: Toast[] = $state([])
  public readonly toastDuration = 5000

  public trigger(toast: Toast) {
    this.toasts.push(toast)

    setTimeout(() => {
      const index = this.toasts.findIndex(t => t.id === toast.id)
      if (index !== -1) {
        this.toasts.splice(index, 1)
      }
    }, this.toastDuration)

  }

  public clear() {
    this.toasts = []
  }
}

export const toastController = new ToastController()