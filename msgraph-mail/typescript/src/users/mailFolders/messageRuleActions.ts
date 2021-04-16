import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Importance} from '../importance';
import {Recipient} from '../recipient';

export class MessageRuleActions implements Parsable<MessageRuleActions> {
    public assignCategories?: string[] | undefined;
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
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfPrimitiveValues<string>("assignCategories", this.assignCategories);
        writer.writeStringValue("copyToFolder", this.copyToFolder);
        writer.writeBooleanValue("delete", this.delete);
        writer.writeCollectionOfObjectValues<Recipient>("forwardAsAttachmentTo", this.forwardAsAttachmentTo);
        writer.writeCollectionOfObjectValues<Recipient>("forwardTo", this.forwardTo);
        writer.writeBooleanValue("markAsRead", this.markAsRead);
        writer.writeEnumValue<Importance>("markImportance", this.markImportance);
        writer.writeStringValue("moveToFolder", this.moveToFolder);
        writer.writeBooleanValue("permanentDelete", this.permanentDelete);
        writer.writeCollectionOfObjectValues<Recipient>("redirectTo", this.redirectTo);
        writer.writeBooleanValue("stopProcessingRules", this.stopProcessingRules);
    };
    public deserializeFields () : Map<string, (item: MessageRuleActions, node: ParseNode) => void> {
        return new Map<string, (item: MessageRuleActions, node: ParseNode) => void>([
            ["assignCategories", (o, n) => { o.assignCategories = n.getCollectionOfPrimitiveValues<string>(); }],
            ["copyToFolder", (o, n) => { o.copyToFolder = n.getStringValue(); }],
            ["delete", (o, n) => { o.delete = n.getBooleanValue(); }],
            ["forwardAsAttachmentTo", (o, n) => { o.forwardAsAttachmentTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["forwardTo", (o, n) => { o.forwardTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["markAsRead", (o, n) => { o.markAsRead = n.getBooleanValue(); }],
            ["markImportance", (o, n) => { o.markImportance = n.getEnumValue<Importance>(Importance); }],
            ["moveToFolder", (o, n) => { o.moveToFolder = n.getStringValue(); }],
            ["permanentDelete", (o, n) => { o.permanentDelete = n.getBooleanValue(); }],
            ["redirectTo", (o, n) => { o.redirectTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["stopProcessingRules", (o, n) => { o.stopProcessingRules = n.getBooleanValue(); }],
        ]);
    };
}
