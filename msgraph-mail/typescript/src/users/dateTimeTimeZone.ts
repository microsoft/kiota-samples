import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class DateTimeTimeZone implements Parsable<DateTimeTimeZone> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).  */
    public dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.  */
    public timeZone?: string | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: DateTimeTimeZone, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: DateTimeTimeZone, node: ParseNode) => void> {
        return new Map<string, (item: DateTimeTimeZone, node: ParseNode) => void>([
            ["dateTime", (o, n) => { o.dateTime = n.getStringValue(); }],
            ["timeZone", (o, n) => { o.timeZone = n.getStringValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("dateTime", this.dateTime);
        writer.writeStringValue("timeZone", this.timeZone);
        writer.writeAdditionalData(this.additionalData);
    };
}
