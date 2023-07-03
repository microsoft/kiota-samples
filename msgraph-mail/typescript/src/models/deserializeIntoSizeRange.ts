import {SizeRange} from './sizeRange';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoSizeRange(sizeRange: SizeRange | undefined = {} as SizeRange) : Record<string, (node: ParseNode) => void> {
    return {
        "maximumSize": n => { sizeRange.maximumSize = n.getNumberValue(); },
        "minimumSize": n => { sizeRange.minimumSize = n.getNumberValue(); },
    }
}
