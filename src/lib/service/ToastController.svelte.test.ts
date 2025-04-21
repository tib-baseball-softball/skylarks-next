import {afterEach, beforeEach, describe, expect, it, vi} from "vitest";
import {toastController} from "./ToastController.svelte";
import type {Toast} from "$lib/types/Toast";

describe("ToastController", () => {
  // Mock data
  const mockToast: Toast = {
    id: "test-id",
    message: "Test message",
    background: "bg-primary"
  };

  // Reset the controller before each test
  beforeEach(() => {
    toastController.clear();
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe("trigger", () => {
    it("adds a toast to the toasts array", () => {
      // Initial state should be empty
      expect(toastController.toasts.length).toBe(0);
      
      // Trigger a toast
      toastController.trigger(mockToast);
      
      // Should have one toast
      expect(toastController.toasts.length).toBe(1);
      expect(toastController.toasts[0]).toEqual(mockToast);
    });

    it("automatically removes the toast after the timeout", () => {
      // Trigger a toast
      toastController.trigger(mockToast);
      
      // Should have one toast
      expect(toastController.toasts.length).toBe(1);
      
      // Fast-forward time
      vi.advanceTimersByTime(5000);
      
      // Toast should be removed
      expect(toastController.toasts.length).toBe(0);
    });

    it("correctly identifies and removes the specific toast by id", () => {
      // Create multiple toasts with different IDs
      const toast1: Toast = { id: "id1", message: "Message 1" };
      const toast2: Toast = { id: "id2", message: "Message 2" };

      // Add both toasts
      toastController.trigger(toast1);
      vi.advanceTimersByTime(1000); // slight time offset between the two
      toastController.trigger(toast2);

      // Should have two toasts
      expect(toastController.toasts.length).toBe(2);

      // Fast-forward time to remove the first toast with some padding
      vi.advanceTimersByTime(toastController.toastDuration - 1000);

      // Should have one toast left (the second one)
      expect(toastController.toasts.length).toBe(1);
      expect(toastController.toasts.at(0))?.toEqual(toast2);
    });
  });

  describe("clear", () => {
    it("removes all toasts from the array", () => {
      // Add multiple toasts
      toastController.trigger({ id: "id1", message: "Message 1" });
      toastController.trigger({ id: "id2", message: "Message 2" });
      
      // Should have two toasts
      expect(toastController.toasts.length).toBe(2);
      
      // Clear all toasts
      toastController.clear();
      
      // Should have no toasts
      expect(toastController.toasts.length).toBe(0);
    });
  });
});