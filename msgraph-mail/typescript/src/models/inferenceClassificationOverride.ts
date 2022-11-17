import {EmailAddress} from './emailAddress';
import {Entity} from './entity';
import {InferenceClassificationType} from './inferenceClassificationType';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface InferenceClassificationOverride extends Entity, Partial<Parsable> {
    /** The classifyAs property */
    classifyAs?: InferenceClassificationType | undefined;
    /** The senderEmailAddress property */
    senderEmailAddress?: EmailAddress | undefined;
}
