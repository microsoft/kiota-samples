import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class DateTimeTimeZone implements Parsable<DateTimeTimeZone> {
    public dateTime?: string | undefined;
    public timeZone?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("dateTime", this.dateTime);
        writer.writeStringValue("timeZone", this.timeZone);
    };
    public deserializeFields () : Map<string, (item: DateTimeTimeZone, node: ParseNode) => void> {
        return new Map<string, (item: DateTimeTimeZone, node: ParseNode) => void>([
            ["dateTime", (o, n) => { o.dateTime = n.getStringValue(); }],
            ["timeZone", (o, n) => { o.timeZone = n.getStringValue(); }],
        ]);
    };
}
