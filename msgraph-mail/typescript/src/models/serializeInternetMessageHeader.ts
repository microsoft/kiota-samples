import type {InternetMessageHeader} from './internetMessageHeader';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeInternetMessageHeader(writer: SerializationWriter, internetMessageHeader: InternetMessageHeader | undefined = {} as InternetMessageHeader) : void {
        writer.writeStringValue("name", internetMessageHeader.name);
        writer.writeStringValue("value", internetMessageHeader.value);
        writer.writeAdditionalData(internetMessageHeader.additionalData);
}
