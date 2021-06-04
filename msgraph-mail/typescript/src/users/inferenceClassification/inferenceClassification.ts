import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from '../entity';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';

export class InferenceClassification extends Entity implements Parsable {
    /** A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.  */
    public overrides?: InferenceClassificationOverride[] | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["overrides", (o, n) => { (o as unknown as InferenceClassification).overrides = n.getCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("overrides", this.overrides);
    };
}
