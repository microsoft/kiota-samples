import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from '../entity';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';

export class InferenceClassification extends Entity implements Parsable<InferenceClassification> {
    /** A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.  */
    public overrides?: InferenceClassificationOverride[] | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("overrides", this.overrides);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: InferenceClassification, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: InferenceClassification, node: ParseNode) => void> {
        return new Map<string, (item: InferenceClassification, node: ParseNode) => void>([...super.deserializeFields(),
            ["overrides", (o, n) => { o.overrides = n.getCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride); }],
        ]);
    };
}
