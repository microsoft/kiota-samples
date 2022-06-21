import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface DateTimeTimeZone extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'. */
    dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values. */
    timeZone?: string | undefined;
}
