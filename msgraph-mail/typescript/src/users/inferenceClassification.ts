import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';
import {InferenceClassificationOverride} from './inferenceClassification/inferenceClassificationOverride';

export class InferenceClassification extends Entity implements Parsable<InferenceClassification> {
    public overrides?: InferenceClassificationOverride[] | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("overrides", this.overrides);
    };
    public deserializeFields () : Map<string, (item: InferenceClassification, node: ParseNode) => void> {
        return new Map<string, (item: InferenceClassification, node: ParseNode) => void>([...super.DeserializeFields(),
            ["overrides", (o, n) => { o.overrides = n.getCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride); }],
        ]);
    };
}
