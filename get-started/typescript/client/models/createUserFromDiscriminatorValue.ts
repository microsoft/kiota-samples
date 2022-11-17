import {User} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createUserFromDiscriminatorValue(parseNode: ParseNode | undefined) : User {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new User();
}
