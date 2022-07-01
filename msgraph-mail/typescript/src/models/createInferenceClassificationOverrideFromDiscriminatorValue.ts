import {InferenceClassificationOverrideImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInferenceClassificationOverrideFromDiscriminatorValue(parseNode: ParseNode | undefined) : InferenceClassificationOverrideImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new InferenceClassificationOverrideImpl();
}
