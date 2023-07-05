import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface Post extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The body property
     */
    body?: string | undefined;
    /**
     * The id property
     */
    id?: number | undefined;
    /**
     * The title property
     */
    title?: string | undefined;
    /**
     * The userId property
     */
    userId?: number | undefined;
}
