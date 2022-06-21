import {createMessageRuleActionsFromDiscriminatorValue} from './createMessageRuleActionsFromDiscriminatorValue';
import {createMessageRulePredicatesFromDiscriminatorValue} from './createMessageRulePredicatesFromDiscriminatorValue';
import {EntityImpl, MessageRuleActionsImpl, MessageRulePredicatesImpl} from './index';
import {MessageRule} from './messageRule';
import {MessageRuleActions} from './messageRuleActions';
import {MessageRulePredicates} from './messageRulePredicates';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleImpl extends EntityImpl implements MessageRule {
    /** The actions property */
    public actions?: MessageRuleActions | undefined;
    /** The conditions property */
    public conditions?: MessageRulePredicates | undefined;
    /** The display name of the rule. */
    public displayName?: string | undefined;
    /** The exceptions property */
    public exceptions?: MessageRulePredicates | undefined;
    /** Indicates whether the rule is in an error condition. Read-only. */
    public hasError?: boolean | undefined;
    /** Indicates whether the rule is enabled to be applied to messages. */
    public isEnabled?: boolean | undefined;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. */
    public isReadOnly?: boolean | undefined;
    /** Indicates the order in which the rule is executed, among other rules. */
    public sequence?: number | undefined;
    /**
     * Instantiates a new messageRule and sets the default values.
     * @param messageRuleParameterValue 
     */
    public constructor(messageRuleParameterValue?: MessageRule | undefined) {
        super(messageRuleParameterValue);
        this.actions = messageRuleParameterValue?.actions;
        this.conditions = messageRuleParameterValue?.conditions;
        this.displayName = messageRuleParameterValue?.displayName;
        this.exceptions = messageRuleParameterValue?.exceptions;
        this.hasError = messageRuleParameterValue?.hasError;
        this.isEnabled = messageRuleParameterValue?.isEnabled;
        this.isReadOnly = messageRuleParameterValue?.isReadOnly;
        this.sequence = messageRuleParameterValue?.sequence;
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
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.actions){
            writer.writeObjectValue<MessageRuleActionsImpl>("actions", new MessageRuleActionsImpl(this.actions));
        }
        if(this.conditions){
            writer.writeObjectValue<MessageRulePredicatesImpl>("conditions", new MessageRulePredicatesImpl(this.conditions));
        }
        if(this.displayName){
            writer.writeStringValue("displayName", this.displayName);
        }
        if(this.exceptions){
            writer.writeObjectValue<MessageRulePredicatesImpl>("exceptions", new MessageRulePredicatesImpl(this.exceptions));
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
}
