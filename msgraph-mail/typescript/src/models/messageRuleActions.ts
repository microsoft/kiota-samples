import {Importance} from './importance';
import {Recipient} from './recipient';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface MessageRuleActions extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * A list of categories to be assigned to a message.
     */
    assignCategories?: string[] | undefined;
    /**
     * The ID of a folder that a message is to be copied to.
     */
    copyToFolder?: string | undefined;
    /**
     * Indicates whether a message should be moved to the Deleted Items folder.
     */
    delete?: boolean | undefined;
    /**
     * The email addresses of the recipients to which a message should be forwarded as an attachment.
     */
    forwardAsAttachmentTo?: Recipient[] | undefined;
    /**
     * The email addresses of the recipients to which a message should be forwarded.
     */
    forwardTo?: Recipient[] | undefined;
    /**
     * Indicates whether a message should be marked as read.
     */
    markAsRead?: boolean | undefined;
    /**
     * The markImportance property
     */
    markImportance?: Importance | undefined;
    /**
     * The ID of the folder that a message will be moved to.
     */
    moveToFolder?: string | undefined;
    /**
     * Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     */
    permanentDelete?: boolean | undefined;
    /**
     * The email addresses to which a message should be redirected.
     */
    redirectTo?: Recipient[] | undefined;
    /**
     * Indicates whether subsequent rules should be evaluated.
     */
    stopProcessingRules?: boolean | undefined;
}
