import {deserializeIntoInternetMessageHeader} from './deserializeIntoInternetMessageHeader';
import {InternetMessageHeader} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInternetMessageHeaderFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInternetMessageHeader;
}
