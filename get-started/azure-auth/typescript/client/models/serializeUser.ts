import { type User } from './user';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function serializeUser(writer: SerializationWriter, user: User | undefined = {} as User) : void {
        writer.writeStringValue("displayName", user.displayName);
        writer.writeStringValue("id", user.id);
        writer.writeAdditionalData(user.additionalData);
}
