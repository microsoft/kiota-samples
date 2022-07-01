import {InternetMessageHeaderImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInternetMessageHeaderFromDiscriminatorValue(parseNode: ParseNode | undefined) : InternetMessageHeaderImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new InternetMessageHeaderImpl();
}
