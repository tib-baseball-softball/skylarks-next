import type {Options} from "flatpickr/dist/types/options";

export class DateTimeUtility {
  public static readonly timeFormatShort = new Intl.DateTimeFormat("de-DE", {
    weekday: "short",
    hour: "2-digit",
    minute: "2-digit"
  });

  public static readonly dateTimeFormatMedium = new Intl.DateTimeFormat("de-DE", {
    dateStyle: "medium",
    timeStyle: "short"
  });

  //@ts-ignore
  public static readonly eventDateFormat: DateTimeFormatOptions = {
    weekday: 'long',
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
  };

  //@ts-ignore
  public static readonly eventSeriesDateFormat: DateTimeFormatOptions = {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
  };

  //@ts-ignore
  public static readonly eventTimeFormat: DateTimeFormatOptions = {
    hour: 'numeric',
    minute: 'numeric',
  };

  public static readonly datePickerOptions: Options = {
    enableTime: true,
    dateFormat: "c", // ISO 8601
    altInput: true,
    altFormat: "j F Y - H:i",
    time_24hr: true,
  };

  public static readonly datePickerOptionsNoTime: Options = {
    enableTime: false,
    dateFormat: "c", // ISO 8601
    altInput: true,
    altFormat: "j F Y",
  };

  public static getRelativeTimeString(date: Date, locale: string = navigator.language): string {
    const now = new Date();
    const diff = date.getTime() - now.getTime(); // Difference in milliseconds

    const formatter = new Intl.RelativeTimeFormat(locale, {numeric: 'auto'});

    const units: { unit: Intl.RelativeTimeFormatUnit; value: number }[] = [
      {unit: 'year', value: 1000 * 60 * 60 * 24 * 365},
      {unit: 'month', value: 1000 * 60 * 60 * 24 * 30},
      {unit: 'week', value: 1000 * 60 * 60 * 24 * 7},
      {unit: 'day', value: 1000 * 60 * 60 * 24},
      {unit: 'hour', value: 1000 * 60 * 60},
      {unit: 'minute', value: 1000 * 60},
      {unit: 'second', value: 1000}
    ];

    for (const {unit, value} of units) {
      const relativeValue = Math.round(diff / value);
      if (Math.abs(relativeValue) > 0) {
        return formatter.format(relativeValue, unit);
      }
    }

    return "just now";
  }

  /**
   * Time snippets are entered as strings from HTML time picker (`10:00`).
   * Convert these to UTC before sending to backend.
   * @param time
   */
  public static convertTimeToUTC(time: string): string {
    if (!/^\d{2}:\d{2}$/.test(time)) {
      throw new Error("Invalid time format. Expected 'HH:mm'.");
    }

    const [hours, minutes] = time.split(":").map(Number);

    // Create a Date object for the current date and the given time in the user's local time zone
    const now = new Date();
    const localDate = new Date(now.getFullYear(), now.getMonth(), now.getDate(), hours, minutes);

    const utcHours = localDate.getUTCHours();
    const utcMinutes = localDate.getUTCMinutes();

    return `${utcHours.toString().padStart(2, "0")}:${utcMinutes.toString().padStart(2, "0")}`;
  }

  public static convertTimeFromUTC(utcTime: string): string {
    if (!/^\d{2}:\d{2}$/.test(utcTime)) {
      throw new Error("Invalid time format. Expected 'HH:mm'.");
    }

    const [utcHours, utcMinutes] = utcTime.split(":").map(Number);

    const utcDate = new Date();
    utcDate.setUTCHours(utcHours, utcMinutes, 0, 0);

    const localHours = utcDate.getHours();
    const localMinutes = utcDate.getMinutes();

    return `${localHours.toString().padStart(2, "0")}:${localMinutes.toString().padStart(2, "0")}`;
  }

  /**
   * Slightly overcomplicated method to parse the BSM "time" string into something usable, with help from generative AI.
   * BSM does not use ISO 8601 format.
   *
   * Format example: 2024-04-07 12:05:00 +0200
   */
  public static parseDateFromBSMString(formattedString: string): Date {
    const dateParts = formattedString.split(/[- :]/);

    // Extracting date parts
    const year = parseInt(dateParts[0], 10);
    const month = parseInt(dateParts[1], 10) - 1; // Month is zero-based
    const day = parseInt(dateParts[2], 10);
    const hour = parseInt(dateParts[3], 10);
    const minute = parseInt(dateParts[4], 10);
    const second = parseInt(dateParts[5], 10);

    // Extracting time zone offset
    const timezoneOffsetHours = parseInt(dateParts[6].substring(0, 3), 10);
    const timezoneOffsetMinutes = parseInt(dateParts[6].substring(3), 10);
    const timezoneOffset = (timezoneOffsetHours * 60) + timezoneOffsetMinutes;

    // Creating date object
    const date = new Date(year, month, day, hour, minute, second);

    // Adjusting for time zone offset
    const utcTime = date.getTime() + (date.getTimezoneOffset() * 60000);
    const adjustedTime = utcTime + (timezoneOffset * 60000);

    return new Date(adjustedTime);
  }
}