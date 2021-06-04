import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {EmailAddress} from '../emailAddress';
import {Entity} from '../entity';
import {InferenceClassificationType} from '../inferenceClassificationType';

export class InferenceClassificationOverride extends Entity implements Parsable {
    public classifyAs?: InferenceClassificationType | undefined;
    public senderEmailAddress?: EmailAddress | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["classifyAs", (o, n) => { (o as unknown as InferenceClassificationOverride).classifyAs = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); }],
            ["senderEmailAddress", (o, n) => { (o as unknown as InferenceClassificationOverride).senderEmailAddress = n.getObjectValue<EmailAddress>(EmailAddress); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeEnumValue<InferenceClassificationType>("classifyAs", this.classifyAs);
        writer.writeObjectValue<EmailAddress>("senderEmailAddress", this.senderEmailAddress);
    };
}
