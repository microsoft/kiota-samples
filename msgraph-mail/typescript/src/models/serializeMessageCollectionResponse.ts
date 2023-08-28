import type {Message} from './message';
import type {MessageCollectionResponse} from './messageCollectionResponse';
import {serializeMessage} from './serializeMessage';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMessageCollectionResponse(writer: SerializationWriter, messageCollectionResponse: MessageCollectionResponse | undefined = {} as MessageCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", messageCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Message>("value", messageCollectionResponse.value, serializeMessage);
        writer.writeAdditionalData(messageCollectionResponse.additionalData);
}
