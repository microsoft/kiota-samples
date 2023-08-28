import {createEmailAddressFromDiscriminatorValue} from './createEmailAddressFromDiscriminatorValue';
import {deserializeIntoEntity} from './deserializeIntoEntity';
import type {EmailAddress} from './emailAddress';
import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {InferenceClassificationType} from './inferenceClassificationType';
import {serializeEmailAddress} from './serializeEmailAddress';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoInferenceClassificationOverride(inferenceClassificationOverride: InferenceClassificationOverride | undefined = {} as InferenceClassificationOverride) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(inferenceClassificationOverride),
        "classifyAs": n => { inferenceClassificationOverride.classifyAs = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); },
        "senderEmailAddress": n => { inferenceClassificationOverride.senderEmailAddress = n.getObjectValue<EmailAddress>(createEmailAddressFromDiscriminatorValue); },
    }
}
