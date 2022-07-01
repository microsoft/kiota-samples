import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface SizeRange extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    maximumSize?: number | undefined;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    minimumSize?: number | undefined;
}
