import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Message} from '../message';

export class MessagesResponse implements Parsable<MessagesResponse> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    public nextLink?: string | undefined;
    public value?: Message[] | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: MessagesResponse, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: MessagesResponse, node: ParseNode) => void> {
        return new Map<string, (item: MessagesResponse, node: ParseNode) => void>([
            ["@odata.nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<Message>(Message); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("@odata.nextLink", this.nextLink);
        writer.writeCollectionOfObjectValues<Message>("value", this.value);
        writer.writeAdditionalData(this.additionalData);
    };
}
