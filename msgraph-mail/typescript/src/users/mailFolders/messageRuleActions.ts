import {Importance} from '../importance';
import {Recipient} from '../recipient';

export class MessageRuleActions {
    public assignCategories?: object[] | undefined;
    public copyToFolder?: string | undefined;
    public delete?: boolean | undefined;
    public forwardAsAttachmentTo?: Recipient[] | undefined;
    public forwardTo?: Recipient[] | undefined;
    public markAsRead?: boolean | undefined;
    public markImportance?: Importance | undefined;
    public moveToFolder?: string | undefined;
    public permanentDelete?: boolean | undefined;
    public redirectTo?: Recipient[] | undefined;
    public stopProcessingRules?: boolean | undefined;
}
