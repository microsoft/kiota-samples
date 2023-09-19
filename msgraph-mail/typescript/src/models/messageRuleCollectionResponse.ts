import { createMessageRuleFromDiscriminatorValue, serializeMessageRule, type MessageRule } from './messageRule';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageRuleCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessageRuleCollectionResponse;
}
export function deserializeIntoMessageRuleCollectionResponse(messageRuleCollectionResponse: MessageRuleCollectionResponse | undefined = {} as MessageRuleCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { messageRuleCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { messageRuleCollectionResponse.value = n.getCollectionOfObjectValues<MessageRule>(createMessageRuleFromDiscriminatorValue); },
    }
}
export interface MessageRuleCollectionResponse extends AdditionalDataHolder, Parsable {
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
    value?: MessageRule[] | undefined;
}
export function serializeMessageRuleCollectionResponse(writer: SerializationWriter, messageRuleCollectionResponse: MessageRuleCollectionResponse | undefined = {} as MessageRuleCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", messageRuleCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<MessageRule>("value", messageRuleCollectionResponse.value, );
        writer.writeAdditionalData(messageRuleCollectionResponse.additionalData);
}
