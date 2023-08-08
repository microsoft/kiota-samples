import type {User} from './user';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoUser(user: User | undefined = {} as User) : Record<string, (node: ParseNode) => void> {
    return {
        "displayName": n => { user.displayName = n.getStringValue(); },
        "id": n => { user.id = n.getStringValue(); },
    }
}
