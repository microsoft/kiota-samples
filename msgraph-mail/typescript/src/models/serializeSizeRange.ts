import type {SizeRange} from './sizeRange';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeSizeRange(writer: SerializationWriter, sizeRange: SizeRange | undefined = {} as SizeRange) : void {
        writer.writeNumberValue("maximumSize", sizeRange.maximumSize);
        writer.writeNumberValue("minimumSize", sizeRange.minimumSize);
        writer.writeAdditionalData(sizeRange.additionalData);
}
