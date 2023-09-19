import { createMessageFromDiscriminatorValue, serializeMessage, type Message } from './message';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessageCollectionResponse;
}
export function deserializeIntoMessageCollectionResponse(messageCollectionResponse: MessageCollectionResponse | undefined = {} as MessageCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { messageCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { messageCollectionResponse.value = n.getCollectionOfObjectValues<Message>(createMessageFromDiscriminatorValue); },
    }
}
export interface MessageCollectionResponse extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The OdataNextLink property
     */
    odataNextLink?: string | undefined;
    /**
     * The value property
     */
    value?: Message[] | undefined;
}
export function serializeMessageCollectionResponse(writer: SerializationWriter, messageCollectionResponse: MessageCollectionResponse | undefined = {} as MessageCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", messageCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Message>("value", messageCollectionResponse.value, );
        writer.writeAdditionalData(messageCollectionResponse.additionalData);
}
