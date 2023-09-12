import { deserializeIntoUser } from './deserializeIntoUser';
import { type User } from './index';
import { type ParseNode } from '@microsoft/kiota-abstractions';

export function createUserFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoUser;
}
