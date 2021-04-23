import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {MailFolder} from '../mailFolder';

export class ChildFoldersResponse implements Parsable<ChildFoldersResponse> {
    public value?: MailFolder[] | undefined;
    public nextLink?: string | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<MailFolder>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
        writer.writeAdditionalData(this.additionalData);
    };
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: ChildFoldersResponse, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: ChildFoldersResponse, node: ParseNode) => void> {
        return new Map<string, (item: ChildFoldersResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<MailFolder>(MailFolder); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
