import { BodyType } from './bodyType';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createItemBodyFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoItemBody;
}
export function deserializeIntoItemBody(itemBody: ItemBody | undefined = {} as ItemBody) : Record<string, (node: ParseNode) => void> {
    return {
        "content": n => { itemBody.content = n.getStringValue(); },
        "contentType": n => { itemBody.contentType = n.getEnumValue<BodyType>(BodyType); },
    }
}
export interface ItemBody extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The content of the item.
     */
    content?: string | undefined;
    /**
     * The contentType property
     */
    contentType?: BodyType | undefined;
}
export function serializeItemBody(writer: SerializationWriter, itemBody: ItemBody | undefined = {} as ItemBody) : void {
        writer.writeStringValue("content", itemBody.content);
        writer.writeEnumValue<BodyType>("contentType", itemBody.contentType);
        writer.writeAdditionalData(itemBody.additionalData);
}
