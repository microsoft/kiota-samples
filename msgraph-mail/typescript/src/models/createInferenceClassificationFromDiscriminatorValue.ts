import {deserializeIntoInferenceClassification} from './deserializeIntoInferenceClassification';
import {InferenceClassification} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInferenceClassificationFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInferenceClassification;
}
