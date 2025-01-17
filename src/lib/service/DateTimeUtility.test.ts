import {describe, expect, it} from "vitest";
import {DateTimeUtility} from "./DateTimeUtility.ts";

/**
 * Massively flawed due to DST changes during the year. A failing test may not need the code is broken!
 */
describe("Time conversion helpers", () => {
  describe("convertToUTC", () => {
    it("converts local time to UTC", () => {
      const date = new Date();
      date.getTimezoneOffset()
      expect(DateTimeUtility.convertTimeToUTC("10:00")).toBe("09:00"); // UTC + 1 for Berlin Winter
    });

    it("throws an error for invalid input format", () => {
      expect(() => DateTimeUtility.convertTimeToUTC("invalid")).toThrow("Invalid time format. Expected 'HH:mm'.");
    });
  });

  describe("convertFromUTC", () => {
    it("converts UTC time to local time", () => {
      expect(DateTimeUtility.convertTimeFromUTC("09:00")).toBe("10:00");
    });

    it("throws an error for invalid input format", () => {
      expect(() => DateTimeUtility.convertTimeFromUTC("invalid")).toThrow("Invalid time format. Expected 'HH:mm'.");
    });
  });
});
