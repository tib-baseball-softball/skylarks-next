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
  public static readonly eventSeriesDateTimeFormat: DateTimeFormatOptions = {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
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