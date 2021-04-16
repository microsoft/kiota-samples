import {Entity} from './entity';
import {InferenceClassificationOverride} from './inferenceClassification/inferenceClassificationOverride';

export class InferenceClassification extends Entity {
    public overrides?: InferenceClassificationOverride[] | undefined;
}
