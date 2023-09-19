import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createSizeRangeFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoSizeRange;
}
export function deserializeIntoSizeRange(sizeRange: SizeRange | undefined = {} as SizeRange) : Record<string, (node: ParseNode) => void> {
    return {
        "maximumSize": n => { sizeRange.maximumSize = n.getNumberValue(); },
        "minimumSize": n => { sizeRange.minimumSize = n.getNumberValue(); },
    }
}
export function serializeSizeRange(writer: SerializationWriter, sizeRange: SizeRange | undefined = {} as SizeRange) : void {
        writer.writeNumberValue("maximumSize", sizeRange.maximumSize);
        writer.writeNumberValue("minimumSize", sizeRange.minimumSize);
        writer.writeAdditionalData(sizeRange.additionalData);
}
export interface SizeRange extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     */
    maximumSize?: number | undefined;
    /**
     * The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     */
    minimumSize?: number | undefined;
}
