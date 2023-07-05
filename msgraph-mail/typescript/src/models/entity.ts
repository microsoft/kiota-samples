import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface Entity extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The unique idenfier for an entity. Read-only.
     */
    id?: string | undefined;
}
