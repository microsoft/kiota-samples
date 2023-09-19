import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createInternetMessageHeaderFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInternetMessageHeader;
}
export function deserializeIntoInternetMessageHeader(internetMessageHeader: InternetMessageHeader | undefined = {} as InternetMessageHeader) : Record<string, (node: ParseNode) => void> {
    return {
        "name": n => { internetMessageHeader.name = n.getStringValue(); },
        "value": n => { internetMessageHeader.value = n.getStringValue(); },
    }
}
export interface InternetMessageHeader extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * Represents the key in a key-value pair.
     */
    name?: string | undefined;
    /**
     * The value in a key-value pair.
     */
    value?: string | undefined;
}
export function serializeInternetMessageHeader(writer: SerializationWriter, internetMessageHeader: InternetMessageHeader | undefined = {} as InternetMessageHeader) : void {
        writer.writeStringValue("name", internetMessageHeader.name);
        writer.writeStringValue("value", internetMessageHeader.value);
        writer.writeAdditionalData(internetMessageHeader.additionalData);
}
