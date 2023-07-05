import {InternetMessageHeader} from './internetMessageHeader';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoInternetMessageHeader(internetMessageHeader: InternetMessageHeader | undefined = {} as InternetMessageHeader) : Record<string, (node: ParseNode) => void> {
    return {
        "name": n => { internetMessageHeader.name = n.getStringValue(); },
        "value": n => { internetMessageHeader.value = n.getStringValue(); },
    }
}
