import {InferenceClassificationImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInferenceClassificationFromDiscriminatorValue(parseNode: ParseNode | undefined) : InferenceClassificationImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new InferenceClassificationImpl();
}
