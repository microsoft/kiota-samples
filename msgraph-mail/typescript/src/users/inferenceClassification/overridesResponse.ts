import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';

export class OverridesResponse implements Parsable<OverridesResponse> {
    public value?: InferenceClassificationOverride[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: OverridesResponse, node: ParseNode) => void> {
        return new Map<string, (item: OverridesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
