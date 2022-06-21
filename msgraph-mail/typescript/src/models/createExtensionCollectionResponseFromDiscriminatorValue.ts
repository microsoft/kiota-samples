import {ExtensionCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createExtensionCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ExtensionCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ExtensionCollectionResponseImpl();
}
