import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface EmailAddress extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** The email address of the person or entity. */
    address?: string | undefined;
    /** The display name of the person or entity. */
    name?: string | undefined;
}
