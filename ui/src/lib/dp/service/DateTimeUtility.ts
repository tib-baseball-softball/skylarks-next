export class DateTimeUtility {
  //@ts-ignore
  public static readonly eventDateFormat: DateTimeFormatOptions = {
    weekday: "long",
    year: "numeric",
    month: "numeric",
    day: "numeric",
  };
  //@ts-ignore
  public static readonly eventSeriesDateFormat: DateTimeFormatOptions = {
    year: "numeric",
    month: "numeric",
    day: "numeric",
  };
  //@ts-ignore
  public static readonly eventSeriesDateTimeFormat: DateTimeFormatOptions = {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
  };
  //@ts-ignore
  public static readonly eventTimeFormat: DateTimeFormatOptions = {
    hour: "numeric",
    minute: "numeric",
  };

  public static dateTimeFormatMedium(locale: string) {
    return new Intl.DateTimeFormat(locale, {
      dateStyle: "medium",
      timeStyle: "short",
    });
  }

  public static dateFormatMedium(locale: string) {
    return new Intl.DateTimeFormat(locale, {
      dateStyle: "medium",
    });
  }

  public static dateTimeFormatShort(locale: string) {
    return new Intl.DateTimeFormat(locale, {
      dateStyle: "short",
      timeStyle: "short",
    });
  }

  public static timeFormatShort(locale: string) {
    return new Intl.DateTimeFormat(locale, {
      weekday: "short",
      hour: "2-digit",
      minute: "2-digit",
    });
  }

  public static getRelativeTimeString(
    date: Date,
    locale: string = navigator.language,
  ): string {
    const now = new Date();
    const diff = date.getTime() - now.getTime(); // Difference in milliseconds

    const formatter = new Intl.RelativeTimeFormat(locale, { numeric: "auto" });

    const units: { unit: Intl.RelativeTimeFormatUnit; value: number }[] = [
      { unit: "year", value: 1000 * 60 * 60 * 24 * 365 },
      { unit: "month", value: 1000 * 60 * 60 * 24 * 30 },
      { unit: "week", value: 1000 * 60 * 60 * 24 * 7 },
      { unit: "day", value: 1000 * 60 * 60 * 24 },
      { unit: "hour", value: 1000 * 60 * 60 },
      { unit: "minute", value: 1000 * 60 },
      { unit: "second", value: 1000 },
    ];

    for (const { unit, value } of units) {
      const relativeValue = Math.round(diff / value);
      if (Math.abs(relativeValue) > 0) {
        return formatter.format(relativeValue, unit);
      }
    }

    return "just now";
  }

  /**
   * Converts an RFC3339 string (from PB backend)
   * to a format compatible with <input type="datetime-local">
   *
   * TODO: use Temporal when baseline
   *
   * @param {string} rfc3339String - e.g., "2023-10-27T14:30:00Z"
   * @returns {string} - e.g., "2023-10-27T14:30"
   */
  public static formatForDatetimeLocal(rfc3339String: string): string {
    const date = new Date(rfc3339String);

    if (isNaN(date.getTime())) {
      return "";
    }

    // Adjust the date by the timezone offset to "fake" a UTC time
    // that matches our local time, then call .toISOString()
    const offset = date.getTimezoneOffset() * 60000; // offset in milliseconds
    const localDate = new Date(date.getTime() - offset);

    // .toISOString() returns "YYYY-MM-DDTHH:mm:ss.sssZ"
    // slice(0, 16) to get "YYYY-MM-DDTHH:mm"
    return localDate.toISOString().slice(0, 16);
  }

  public static formatForDate(dateString: string): string {
    if (!dateString) {
      return "";
    }
    const newVal = new Date(dateString).toISOString();

    return newVal.slice(0, 10);
  }

  /**
   * Slightly overcomplicated method to parse the BSM "time" string into something usable.
   * BSM does not use ISO 8601 format.
   *
   * TODO: get rid of this monster. Most BSM date handling happens on the backend anyway.
   * Chrome can parse the date string unmodified directly in `new Date()`, but Safari can't.
   *
   * TODO: use Temporal when baseline
   *
   * @param {string} formattedString - e.g. "2024-04-07 12:05:00 +0200"
   */
  public static parseDateFromBSMString(formattedString: string): Date {
    const dateParts = formattedString.split(/[- :]/);

    const year = parseInt(dateParts[0], 10);
    const month = parseInt(dateParts[1], 10) - 1; // Month is zero-based
    const day = parseInt(dateParts[2], 10);
    const hour = parseInt(dateParts[3], 10);
    const minute = parseInt(dateParts[4], 10);
    const second = parseInt(dateParts[5], 10);

    const timezoneOffsetHours = parseInt(dateParts[6].substring(0, 3), 10);
    const timezoneOffsetMinutes = parseInt(dateParts[6].substring(3), 10);
    const timezoneOffset = timezoneOffsetHours * 60 + timezoneOffsetMinutes;

    const date = new Date(year, month, day, hour, minute, second);

    const utcTime = date.getTime() + date.getTimezoneOffset() * 60000;
    const adjustedTime = utcTime + timezoneOffset * 60000;

    return new Date(adjustedTime);
  }
}
