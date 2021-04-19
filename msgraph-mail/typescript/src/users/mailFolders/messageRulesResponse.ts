import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {MessageRule} from './messageRule';

export class MessageRulesResponse implements Parsable<MessageRulesResponse> {
    public value?: MessageRule[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<MessageRule>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: MessageRulesResponse, node: ParseNode) => void> {
        return new Map<string, (item: MessageRulesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<MessageRule>(MessageRule); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
