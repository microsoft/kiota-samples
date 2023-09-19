import { createExtensionFromDiscriminatorValue, serializeExtension, type Extension } from './extension';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createExtensionCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoExtensionCollectionResponse;
}
export function deserializeIntoExtensionCollectionResponse(extensionCollectionResponse: ExtensionCollectionResponse | undefined = {} as ExtensionCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { extensionCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { extensionCollectionResponse.value = n.getCollectionOfObjectValues<Extension>(createExtensionFromDiscriminatorValue); },
    }
}
export interface ExtensionCollectionResponse extends AdditionalDataHolder, Parsable {
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
    value?: Extension[] | undefined;
}
export function serializeExtensionCollectionResponse(writer: SerializationWriter, extensionCollectionResponse: ExtensionCollectionResponse | undefined = {} as ExtensionCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", extensionCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Extension>("value", extensionCollectionResponse.value, );
        writer.writeAdditionalData(extensionCollectionResponse.additionalData);
}
