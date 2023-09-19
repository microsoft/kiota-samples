import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createDateTimeTimeZoneFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoDateTimeTimeZone;
}
export interface DateTimeTimeZone extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).
     */
    dateTime?: string | undefined;
    /**
     * Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
     */
    timeZone?: string | undefined;
}
export function deserializeIntoDateTimeTimeZone(dateTimeTimeZone: DateTimeTimeZone | undefined = {} as DateTimeTimeZone) : Record<string, (node: ParseNode) => void> {
    return {
        "dateTime": n => { dateTimeTimeZone.dateTime = n.getStringValue(); },
        "timeZone": n => { dateTimeTimeZone.timeZone = n.getStringValue(); },
    }
}
export function serializeDateTimeTimeZone(writer: SerializationWriter, dateTimeTimeZone: DateTimeTimeZone | undefined = {} as DateTimeTimeZone) : void {
        writer.writeStringValue("dateTime", dateTimeTimeZone.dateTime);
        writer.writeStringValue("timeZone", dateTimeTimeZone.timeZone);
        writer.writeAdditionalData(dateTimeTimeZone.additionalData);
}
