import {EmailAddressImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEmailAddressFromDiscriminatorValue(parseNode: ParseNode | undefined) : EmailAddressImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new EmailAddressImpl();
}
