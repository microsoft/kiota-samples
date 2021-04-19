import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Message} from './message';

export class MessagesResponse implements Parsable<MessagesResponse> {
    public value?: Message[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<Message>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: MessagesResponse, node: ParseNode) => void> {
        return new Map<string, (item: MessagesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<Message>(Message); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
