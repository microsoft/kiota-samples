import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface InternetMessageHeader extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** Represents the key in a key-value pair. */
    name?: string | undefined;
    /** The value in a key-value pair. */
    value?: string | undefined;
}
