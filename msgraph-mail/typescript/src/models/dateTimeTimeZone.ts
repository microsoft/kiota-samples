import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface DateTimeTimeZone extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000). */
    dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values. */
    timeZone?: string | undefined;
}
