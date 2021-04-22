import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Extension} from '../../../extension';

export class ExtensionsResponse implements Parsable<ExtensionsResponse> {
    public value?: Extension[] | undefined;
    public nextLink?: string | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<Extension>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
        writer.writeAdditionalData(this.additionalData);
    };
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: ExtensionsResponse, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: ExtensionsResponse, node: ParseNode) => void> {
        return new Map<string, (item: ExtensionsResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<Extension>(Extension); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
