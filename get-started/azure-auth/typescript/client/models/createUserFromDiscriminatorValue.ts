import {deserializeIntoUser} from './deserializeIntoUser';
import {User} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createUserFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoUser;
}
