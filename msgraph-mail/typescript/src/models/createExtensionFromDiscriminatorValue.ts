import {ExtensionImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createExtensionFromDiscriminatorValue(parseNode: ParseNode | undefined) : ExtensionImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ExtensionImpl();
}
