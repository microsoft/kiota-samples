import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class DateTimeTimeZone implements Parsable<DateTimeTimeZone> {
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.  */
    public dateTime?: string | undefined;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.  */
    public timeZone?: string | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("dateTime", this.dateTime);
        writer.writeStringValue("timeZone", this.timeZone);
    };
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
}
