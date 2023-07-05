import {DateTimeTimeZone} from './dateTimeTimeZone';
import {FollowupFlagStatus} from './followupFlagStatus';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface FollowupFlag extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The completedDateTime property
     */
    completedDateTime?: DateTimeTimeZone | undefined;
    /**
     * The dueDateTime property
     */
    dueDateTime?: DateTimeTimeZone | undefined;
    /**
     * The flagStatus property
     */
    flagStatus?: FollowupFlagStatus | undefined;
    /**
     * The startDateTime property
     */
    startDateTime?: DateTimeTimeZone | undefined;
}
