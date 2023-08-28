import type {EmailAddress} from './emailAddress';
import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {InferenceClassificationType} from './inferenceClassificationType';
import {serializeEmailAddress} from './serializeEmailAddress';
import {serializeEntity} from './serializeEntity';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeInferenceClassificationOverride(writer: SerializationWriter, inferenceClassificationOverride: InferenceClassificationOverride | undefined = {} as InferenceClassificationOverride) : void {
        serializeEntity(writer, inferenceClassificationOverride)
        writer.writeEnumValue<InferenceClassificationType>("classifyAs", inferenceClassificationOverride.classifyAs);
        writer.writeObjectValue<EmailAddress>("senderEmailAddress", inferenceClassificationOverride.senderEmailAddress, serializeEmailAddress);
}
