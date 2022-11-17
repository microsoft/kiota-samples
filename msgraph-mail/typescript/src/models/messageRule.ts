<<<<<<< HEAD
import {Entity} from './entity';
import {MessageRuleActions} from './messageRuleActions';
import {MessageRulePredicates} from './messageRulePredicates';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface MessageRule extends Entity, Partial<Parsable> {
    /** The actions property */
    actions?: MessageRuleActions | undefined;
    /** The conditions property */
    conditions?: MessageRulePredicates | undefined;
    /** The display name of the rule. */
    displayName?: string | undefined;
    /** The exceptions property */
    exceptions?: MessageRulePredicates | undefined;
    /** Indicates whether the rule is in an error condition. Read-only. */
    hasError?: boolean | undefined;
    /** Indicates whether the rule is enabled to be applied to messages. */
    isEnabled?: boolean | undefined;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. */
    isReadOnly?: boolean | undefined;
    /** Indicates the order in which the rule is executed, among other rules. */
    sequence?: number | undefined;
=======
import {createMessageRuleActionsFromDiscriminatorValue} from './createMessageRuleActionsFromDiscriminatorValue';
import {createMessageRulePredicatesFromDiscriminatorValue} from './createMessageRulePredicatesFromDiscriminatorValue';
import {EntityImpl, MessageRuleActionsImpl, MessageRulePredicatesImpl} from './index';
import {MessageRule} from './messageRule';
import {MessageRuleActions} from './messageRuleActions';
import {MessageRulePredicates} from './messageRulePredicates';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleImpl extends EntityImpl implements MessageRule {
    /** The actions property */
    private _actions?: MessageRuleActions | undefined;
    /** The conditions property */
    private _conditions?: MessageRulePredicates | undefined;
    /** The display name of the rule. */
    private _displayName?: string | undefined;
    /** The exceptions property */
    private _exceptions?: MessageRulePredicates | undefined;
    /** Indicates whether the rule is in an error condition. Read-only. */
    private _hasError?: boolean | undefined;
    /** Indicates whether the rule is enabled to be applied to messages. */
    private _isEnabled?: boolean | undefined;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. */
    private _isReadOnly?: boolean | undefined;
    /** Indicates the order in which the rule is executed, among other rules. */
    private _sequence?: number | undefined;
    /**
     * Gets the actions property value. The actions property
     * @returns a MessageRuleActionsInterface
     */
    public get actions() {
        return this._actions;
    };
    /**
     * Sets the actions property value. The actions property
     * @param value Value to set for the actions property.
     */
    public set actions(value: MessageRuleActions | undefined) {
        if(value) {
            this._actions = value instanceof MessageRuleActionsImpl? value as MessageRuleActionsImpl: new MessageRuleActionsImpl(value);
        }
    };
    /**
     * Gets the conditions property value. The conditions property
     * @returns a MessageRulePredicatesInterface
     */
    public get conditions() {
        return this._conditions;
    };
    /**
     * Sets the conditions property value. The conditions property
     * @param value Value to set for the conditions property.
     */
    public set conditions(value: MessageRulePredicates | undefined) {
        if(value) {
            this._conditions = value instanceof MessageRulePredicatesImpl? value as MessageRulePredicatesImpl: new MessageRulePredicatesImpl(value);
        }
    };
    /**
     * Instantiates a new messageRule and sets the default values.
     * @param messageRuleParameterValue 
     */
    public constructor(messageRuleParameterValue?: MessageRule | undefined) {
        super(messageRuleParameterValue);
        this._actions = messageRuleParameterValue?.actions;
        this._conditions = messageRuleParameterValue?.conditions;
        this._displayName = messageRuleParameterValue?.displayName;
        this._exceptions = messageRuleParameterValue?.exceptions;
        this._hasError = messageRuleParameterValue?.hasError;
        this._isEnabled = messageRuleParameterValue?.isEnabled;
        this._isReadOnly = messageRuleParameterValue?.isReadOnly;
        this._sequence = messageRuleParameterValue?.sequence;
    };
    /**
     * Gets the displayName property value. The display name of the rule.
     * @returns a string
     */
    public get displayName() {
        return this._displayName;
    };
    /**
     * Sets the displayName property value. The display name of the rule.
     * @param value Value to set for the displayName property.
     */
    public set displayName(value: string | undefined) {
        if(value) {
            this._displayName = value;
        }
    };
    /**
     * Gets the exceptions property value. The exceptions property
     * @returns a MessageRulePredicatesInterface
     */
    public get exceptions() {
        return this._exceptions;
    };
    /**
     * Sets the exceptions property value. The exceptions property
     * @param value Value to set for the exceptions property.
     */
    public set exceptions(value: MessageRulePredicates | undefined) {
        if(value) {
            this._exceptions = value instanceof MessageRulePredicatesImpl? value as MessageRulePredicatesImpl: new MessageRulePredicatesImpl(value);
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "actions": n => { this.actions = n.getObjectValue<MessageRuleActionsImpl>(createMessageRuleActionsFromDiscriminatorValue); },
            "conditions": n => { this.conditions = n.getObjectValue<MessageRulePredicatesImpl>(createMessageRulePredicatesFromDiscriminatorValue); },
            "displayName": n => { this.displayName = n.getStringValue(); },
            "exceptions": n => { this.exceptions = n.getObjectValue<MessageRulePredicatesImpl>(createMessageRulePredicatesFromDiscriminatorValue); },
            "hasError": n => { this.hasError = n.getBooleanValue(); },
            "isEnabled": n => { this.isEnabled = n.getBooleanValue(); },
            "isReadOnly": n => { this.isReadOnly = n.getBooleanValue(); },
            "sequence": n => { this.sequence = n.getNumberValue(); },
        };
    };
    /**
     * Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @returns a boolean
     */
    public get hasError() {
        return this._hasError;
    };
    /**
     * Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @param value Value to set for the hasError property.
     */
    public set hasError(value: boolean | undefined) {
        if(value) {
            this._hasError = value;
        }
    };
    /**
     * Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @returns a boolean
     */
    public get isEnabled() {
        return this._isEnabled;
    };
    /**
     * Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @param value Value to set for the isEnabled property.
     */
    public set isEnabled(value: boolean | undefined) {
        if(value) {
            this._isEnabled = value;
        }
    };
    /**
     * Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @returns a boolean
     */
    public get isReadOnly() {
        return this._isReadOnly;
    };
    /**
     * Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @param value Value to set for the isReadOnly property.
     */
    public set isReadOnly(value: boolean | undefined) {
        if(value) {
            this._isReadOnly = value;
        }
    };
    /**
     * Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @returns a integer
     */
    public get sequence() {
        return this._sequence;
    };
    /**
     * Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @param value Value to set for the sequence property.
     */
    public set sequence(value: number | undefined) {
        if(value) {
            this._sequence = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.actions){
            writer.writeObjectValue<MessageRuleActionsImpl>("actions", (this.actions instanceof MessageRuleActionsImpl? this.actions as MessageRuleActionsImpl: new MessageRuleActionsImpl(this.actions)));
        }
        if(this.conditions){
            writer.writeObjectValue<MessageRulePredicatesImpl>("conditions", (this.conditions instanceof MessageRulePredicatesImpl? this.conditions as MessageRulePredicatesImpl: new MessageRulePredicatesImpl(this.conditions)));
        }
        if(this.displayName){
            writer.writeStringValue("displayName", this.displayName);
        }
        if(this.exceptions){
            writer.writeObjectValue<MessageRulePredicatesImpl>("exceptions", (this.exceptions instanceof MessageRulePredicatesImpl? this.exceptions as MessageRulePredicatesImpl: new MessageRulePredicatesImpl(this.exceptions)));
        }
        if(this.hasError){
            writer.writeBooleanValue("hasError", this.hasError);
        }
        if(this.isEnabled){
            writer.writeBooleanValue("isEnabled", this.isEnabled);
        }
        if(this.isReadOnly){
            writer.writeBooleanValue("isReadOnly", this.isReadOnly);
        }
        if(this.sequence){
            writer.writeNumberValue("sequence", this.sequence);
        }
    };
>>>>>>> main
}
