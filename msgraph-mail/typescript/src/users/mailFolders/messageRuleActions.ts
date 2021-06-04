import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Importance} from '../importance';
import {Recipient} from '../recipient';

export class MessageRuleActions implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** A list of categories to be assigned to a message.  */
    public assignCategories?: string[] | undefined;
    /** The ID of a folder that a message is to be copied to.  */
    public copyToFolder?: string | undefined;
    /** Indicates whether a message should be moved to the Deleted Items folder.  */
    public delete?: boolean | undefined;
    /** The email addresses of the recipients to which a message should be forwarded as an attachment.  */
    public forwardAsAttachmentTo?: Recipient[] | undefined;
    /** The email addresses of the recipients to which a message should be forwarded.  */
    public forwardTo?: Recipient[] | undefined;
    /** Indicates whether a message should be marked as read.  */
    public markAsRead?: boolean | undefined;
    public markImportance?: Importance | undefined;
    /** The ID of the folder that a message will be moved to.  */
    public moveToFolder?: string | undefined;
    /** Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.  */
    public permanentDelete?: boolean | undefined;
    /** The email address to which a message should be redirected.  */
    public redirectTo?: Recipient[] | undefined;
    /** Indicates whether subsequent rules should be evaluated.  */
    public stopProcessingRules?: boolean | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["assignCategories", (o, n) => { (o as unknown as MessageRuleActions).assignCategories = n.getCollectionOfPrimitiveValues<string>(); }],
            ["copyToFolder", (o, n) => { (o as unknown as MessageRuleActions).copyToFolder = n.getStringValue(); }],
            ["delete", (o, n) => { (o as unknown as MessageRuleActions).delete = n.getBooleanValue(); }],
            ["forwardAsAttachmentTo", (o, n) => { (o as unknown as MessageRuleActions).forwardAsAttachmentTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["forwardTo", (o, n) => { (o as unknown as MessageRuleActions).forwardTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["markAsRead", (o, n) => { (o as unknown as MessageRuleActions).markAsRead = n.getBooleanValue(); }],
            ["markImportance", (o, n) => { (o as unknown as MessageRuleActions).markImportance = n.getEnumValue<Importance>(Importance); }],
            ["moveToFolder", (o, n) => { (o as unknown as MessageRuleActions).moveToFolder = n.getStringValue(); }],
            ["permanentDelete", (o, n) => { (o as unknown as MessageRuleActions).permanentDelete = n.getBooleanValue(); }],
            ["redirectTo", (o, n) => { (o as unknown as MessageRuleActions).redirectTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["stopProcessingRules", (o, n) => { (o as unknown as MessageRuleActions).stopProcessingRules = n.getBooleanValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
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
        writer.writeAdditionalData(this.additionalData);
    };
}
