import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {MessageRule} from '../messageRule';

export class MessageRulesResponse implements Parsable<MessageRulesResponse> {
    public value?: MessageRule[] | undefined;
    public nextLink?: string | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<MessageRule>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: MessageRulesResponse, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: MessageRulesResponse, node: ParseNode) => void> {
        return new Map<string, (item: MessageRulesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<MessageRule>(MessageRule); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
