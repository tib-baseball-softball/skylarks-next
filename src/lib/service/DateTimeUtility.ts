export class DateTimeUtility {
    public static readonly timeFormatShort = new Intl.DateTimeFormat("de-DE", {
        weekday: "short",
        hour: "2-digit",
        minute: "2-digit"
    })

    public static readonly dateTimeFormatMedium = new Intl.DateTimeFormat("de-DE", {
        dateStyle: "medium",
        timeStyle: "short"
    })

    /**
     * Slightly overcomplicated method to parse the BSM "time" string into something usable, with help from generative AI.
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