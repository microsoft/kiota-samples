<<<<<<< HEAD
import {ExtensionCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createExtensionCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ExtensionCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ExtensionCollectionResponseImpl();
=======
import {ExtensionCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createExtensionCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ExtensionCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ExtensionCollectionResponse();
>>>>>>> main
}
