import type {EmailAddress} from './emailAddress';
import type {Entity} from './entity';
import {InferenceClassificationType} from './inferenceClassificationType';
import type {Parsable} from '@microsoft/kiota-abstractions';

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
