import {Message} from './message';
import {MessageCollectionResponse} from './messageCollectionResponse';
import {serializeMessage} from './serializeMessage';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMessageCollectionResponse(writer: SerializationWriter, messageCollectionResponse: MessageCollectionResponse | undefined = {} as MessageCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", messageCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Message>("value", messageCollectionResponse.value, serializeMessage);
        writer.writeAdditionalData(messageCollectionResponse.additionalData);
}
