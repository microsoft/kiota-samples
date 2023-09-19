import { createEmailAddressFromDiscriminatorValue, serializeEmailAddress, type EmailAddress } from './emailAddress';
import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { InferenceClassificationType } from './inferenceClassificationType';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createInferenceClassificationOverrideFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInferenceClassificationOverride;
}
export function deserializeIntoInferenceClassificationOverride(inferenceClassificationOverride: InferenceClassificationOverride | undefined = {} as InferenceClassificationOverride) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(inferenceClassificationOverride),
        "classifyAs": n => { inferenceClassificationOverride.classifyAs = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); },
        "senderEmailAddress": n => { inferenceClassificationOverride.senderEmailAddress = n.getObjectValue<EmailAddress>(createEmailAddressFromDiscriminatorValue); },
    }
}
export interface InferenceClassificationOverride extends Entity, Parsable {
    /**
     * The classifyAs property
     */
    classifyAs?: InferenceClassificationType | undefined;
    /**
     * The senderEmailAddress property
     */
    senderEmailAddress?: EmailAddress | undefined;
}
export function serializeInferenceClassificationOverride(writer: SerializationWriter, inferenceClassificationOverride: InferenceClassificationOverride | undefined = {} as InferenceClassificationOverride) : void {
        serializeEntity(writer, inferenceClassificationOverride)
        writer.writeEnumValue<InferenceClassificationType>("classifyAs", inferenceClassificationOverride.classifyAs);
        writer.writeObjectValue<EmailAddress>("senderEmailAddress", inferenceClassificationOverride.senderEmailAddress, );
}
