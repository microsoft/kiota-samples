import {InferenceClassificationOverrideCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : InferenceClassificationOverrideCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new InferenceClassificationOverrideCollectionResponseImpl();
}
