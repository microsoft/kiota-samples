import {User} from './user';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeUser(writer: SerializationWriter, user: User | undefined = {} as User) : void {
        writer.writeStringValue("displayName", user.displayName);
        writer.writeStringValue("id", user.id);
        writer.writeAdditionalData(user.additionalData);
}
