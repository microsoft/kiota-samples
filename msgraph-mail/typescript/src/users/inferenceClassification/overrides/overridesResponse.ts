import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {InferenceClassificationOverride} from '../inferenceClassificationOverride';

export class OverridesResponse implements Parsable<OverridesResponse> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    public nextLink?: string | undefined;
    public value?: InferenceClassificationOverride[] | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: OverridesResponse, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: OverridesResponse, node: ParseNode) => void> {
        return new Map<string, (item: OverridesResponse, node: ParseNode) => void>([
            ["@odata.nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("@odata.nextLink", this.nextLink);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("value", this.value);
        writer.writeAdditionalData(this.additionalData);
    };
}
