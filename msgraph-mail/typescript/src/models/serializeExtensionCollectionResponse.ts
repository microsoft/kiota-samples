import type {Extension} from './extension';
import type {ExtensionCollectionResponse} from './extensionCollectionResponse';
import {serializeExtension} from './serializeExtension';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeExtensionCollectionResponse(writer: SerializationWriter, extensionCollectionResponse: ExtensionCollectionResponse | undefined = {} as ExtensionCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", extensionCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Extension>("value", extensionCollectionResponse.value, serializeExtension);
        writer.writeAdditionalData(extensionCollectionResponse.additionalData);
}
