import {EmailAddress} from '../emailAddress';
import {Entity} from '../entity';
import {InferenceClassificationType} from '../inferenceClassificationType';

export class InferenceClassificationOverride extends Entity {
    public classifyAs?: InferenceClassificationType | undefined;
    public senderEmailAddress?: EmailAddress | undefined;
}
