import {deserializeIntoExtensionCollectionResponse} from './deserializeIntoExtensionCollectionResponse';
import {ExtensionCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createExtensionCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoExtensionCollectionResponse;
}
