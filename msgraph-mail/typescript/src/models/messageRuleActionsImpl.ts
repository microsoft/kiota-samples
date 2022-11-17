import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {Importance} from './importance';
import {RecipientImpl} from './index';
import {MessageRuleActions} from './messageRuleActions';
import {Recipient} from './recipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleActionsImpl implements MessageRuleActions {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** A list of categories to be assigned to a message. */
    private _assignCategories?: string[] | undefined;
    /** The ID of a folder that a message is to be copied to. */
    private _copyToFolder?: string | undefined;
    /** Indicates whether a message should be moved to the Deleted Items folder. */
    private _delete?: boolean | undefined;
    /** The email addresses of the recipients to which a message should be forwarded as an attachment. */
    private _forwardAsAttachmentTo?: Recipient[] | undefined;
    /** The email addresses of the recipients to which a message should be forwarded. */
    private _forwardTo?: Recipient[] | undefined;
    /** Indicates whether a message should be marked as read. */
    private _markAsRead?: boolean | undefined;
    /** The markImportance property */
    private _markImportance?: Importance | undefined;
    /** The ID of the folder that a message will be moved to. */
    private _moveToFolder?: string | undefined;
    /** Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder. */
    private _permanentDelete?: boolean | undefined;
    /** The email addresses to which a message should be redirected. */
    private _redirectTo?: Recipient[] | undefined;
    /** Indicates whether subsequent rules should be evaluated. */
    private _stopProcessingRules?: boolean | undefined;
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @returns a Record<string, unknown>
     */
    public get additionalData() {
        return this._additionalData;
    };
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public set additionalData(value: Record<string, unknown> | undefined) {
        if(value) {
            this._additionalData = value;
        }
    };
    /**
     * Gets the assignCategories property value. A list of categories to be assigned to a message.
     * @returns a string
     */
    public get assignCategories() {
        return this._assignCategories;
    };
    /**
     * Sets the assignCategories property value. A list of categories to be assigned to a message.
     * @param value Value to set for the assignCategories property.
     */
    public set assignCategories(value: string[] | undefined) {
        if(value) {
            this._assignCategories = value;
        }
    };
    /**
     * Instantiates a new messageRuleActions and sets the default values.
     * @param messageRuleActionsParameterValue 
     */
    public constructor(messageRuleActionsParameterValue?: MessageRuleActions | undefined) {
        this._additionalData = {};
        this._additionalData = messageRuleActionsParameterValue?.additionalData ? messageRuleActionsParameterValue?.additionalData! : {};
        this._assignCategories = messageRuleActionsParameterValue?.assignCategories;
        this._copyToFolder = messageRuleActionsParameterValue?.copyToFolder;
        this._delete = messageRuleActionsParameterValue?.delete;
        this._forwardAsAttachmentTo = messageRuleActionsParameterValue?.forwardAsAttachmentTo;
        this._forwardTo = messageRuleActionsParameterValue?.forwardTo;
        this._markAsRead = messageRuleActionsParameterValue?.markAsRead;
        this._markImportance = messageRuleActionsParameterValue?.markImportance;
        this._moveToFolder = messageRuleActionsParameterValue?.moveToFolder;
        this._permanentDelete = messageRuleActionsParameterValue?.permanentDelete;
        this._redirectTo = messageRuleActionsParameterValue?.redirectTo;
        this._stopProcessingRules = messageRuleActionsParameterValue?.stopProcessingRules;
    };
    /**
     * Gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     * @returns a string
     */
    public get copyToFolder() {
        return this._copyToFolder;
    };
    /**
     * Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     * @param value Value to set for the copyToFolder property.
     */
    public set copyToFolder(value: string | undefined) {
        if(value) {
            this._copyToFolder = value;
        }
    };
    /**
     * Gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     * @returns a boolean
     */
    public get delete() {
        return this._delete;
    };
    /**
     * Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     * @param value Value to set for the delete property.
     */
    public set delete(value: boolean | undefined) {
        if(value) {
            this._delete = value;
        }
    };
    /**
     * Gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     * @returns a RecipientInterface
     */
    public get forwardAsAttachmentTo() {
        return this._forwardAsAttachmentTo;
    };
    /**
     * Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     * @param value Value to set for the forwardAsAttachmentTo property.
     */
    public set forwardAsAttachmentTo(value: Recipient[] | undefined) {
        if(value) {
            const forwardAsAttachmentToArrValue: RecipientImpl[] = [];
            this.forwardAsAttachmentTo?.forEach(element => {
                forwardAsAttachmentToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._forwardAsAttachmentTo = forwardAsAttachmentToArrValue;
        }
    };
    /**
     * Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @returns a RecipientInterface
     */
    public get forwardTo() {
        return this._forwardTo;
    };
    /**
     * Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @param value Value to set for the forwardTo property.
     */
    public set forwardTo(value: Recipient[] | undefined) {
        if(value) {
            const forwardToArrValue: RecipientImpl[] = [];
            this.forwardTo?.forEach(element => {
                forwardToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._forwardTo = forwardToArrValue;
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "assignCategories": n => { this.assignCategories = n.getCollectionOfPrimitiveValues<string>(); },
            "copyToFolder": n => { this.copyToFolder = n.getStringValue(); },
            "delete": n => { this.delete = n.getBooleanValue(); },
            "forwardAsAttachmentTo": n => { this.forwardAsAttachmentTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "forwardTo": n => { this.forwardTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "markAsRead": n => { this.markAsRead = n.getBooleanValue(); },
            "markImportance": n => { this.markImportance = n.getEnumValue<Importance>(Importance); },
            "moveToFolder": n => { this.moveToFolder = n.getStringValue(); },
            "permanentDelete": n => { this.permanentDelete = n.getBooleanValue(); },
            "redirectTo": n => { this.redirectTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "stopProcessingRules": n => { this.stopProcessingRules = n.getBooleanValue(); },
        };
    };
    /**
     * Gets the markAsRead property value. Indicates whether a message should be marked as read.
     * @returns a boolean
     */
    public get markAsRead() {
        return this._markAsRead;
    };
    /**
     * Sets the markAsRead property value. Indicates whether a message should be marked as read.
     * @param value Value to set for the markAsRead property.
     */
    public set markAsRead(value: boolean | undefined) {
        if(value) {
            this._markAsRead = value;
        }
    };
    /**
     * Gets the markImportance property value. The markImportance property
     * @returns a importance
     */
    public get markImportance() {
        return this._markImportance;
    };
    /**
     * Sets the markImportance property value. The markImportance property
     * @param value Value to set for the markImportance property.
     */
    public set markImportance(value: Importance | undefined) {
        if(value) {
            this._markImportance = value;
        }
    };
    /**
     * Gets the moveToFolder property value. The ID of the folder that a message will be moved to.
     * @returns a string
     */
    public get moveToFolder() {
        return this._moveToFolder;
    };
    /**
     * Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
     * @param value Value to set for the moveToFolder property.
     */
    public set moveToFolder(value: string | undefined) {
        if(value) {
            this._moveToFolder = value;
        }
    };
    /**
     * Gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     * @returns a boolean
     */
    public get permanentDelete() {
        return this._permanentDelete;
    };
    /**
     * Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     * @param value Value to set for the permanentDelete property.
     */
    public set permanentDelete(value: boolean | undefined) {
        if(value) {
            this._permanentDelete = value;
        }
    };
    /**
     * Gets the redirectTo property value. The email addresses to which a message should be redirected.
     * @returns a RecipientInterface
     */
    public get redirectTo() {
        return this._redirectTo;
    };
    /**
     * Sets the redirectTo property value. The email addresses to which a message should be redirected.
     * @param value Value to set for the redirectTo property.
     */
    public set redirectTo(value: Recipient[] | undefined) {
        if(value) {
            const redirectToArrValue: RecipientImpl[] = [];
            this.redirectTo?.forEach(element => {
                redirectToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._redirectTo = redirectToArrValue;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.assignCategories){
            writer.writeCollectionOfPrimitiveValues<string>("assignCategories", this.assignCategories);
        }
        if(this.copyToFolder){
            writer.writeStringValue("copyToFolder", this.copyToFolder);
        }
        if(this.delete){
            writer.writeBooleanValue("delete", this.delete);
        }
        if(this.forwardAsAttachmentTo && this.forwardAsAttachmentTo.length != 0){        const forwardAsAttachmentToArrValue: RecipientImpl[] = [];
        this.forwardAsAttachmentTo?.forEach(element => {
            forwardAsAttachmentToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("forwardAsAttachmentTo", forwardAsAttachmentToArrValue);
        }
        if(this.forwardTo && this.forwardTo.length != 0){        const forwardToArrValue: RecipientImpl[] = [];
        this.forwardTo?.forEach(element => {
            forwardToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("forwardTo", forwardToArrValue);
        }
        if(this.markAsRead){
            writer.writeBooleanValue("markAsRead", this.markAsRead);
        }
        if(this.markImportance){
            writer.writeEnumValue<Importance>("markImportance", this.markImportance);
        }
        if(this.moveToFolder){
            writer.writeStringValue("moveToFolder", this.moveToFolder);
        }
        if(this.permanentDelete){
            writer.writeBooleanValue("permanentDelete", this.permanentDelete);
        }
        if(this.redirectTo && this.redirectTo.length != 0){        const redirectToArrValue: RecipientImpl[] = [];
        this.redirectTo?.forEach(element => {
            redirectToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("redirectTo", redirectToArrValue);
        }
        if(this.stopProcessingRules){
            writer.writeBooleanValue("stopProcessingRules", this.stopProcessingRules);
        }
        writer.writeAdditionalData(this.additionalData!);
    };
    /**
     * Gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     * @returns a boolean
     */
    public get stopProcessingRules() {
        return this._stopProcessingRules;
    };
    /**
     * Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     * @param value Value to set for the stopProcessingRules property.
     */
    public set stopProcessingRules(value: boolean | undefined) {
        if(value) {
            this._stopProcessingRules = value;
        }
    };
}
