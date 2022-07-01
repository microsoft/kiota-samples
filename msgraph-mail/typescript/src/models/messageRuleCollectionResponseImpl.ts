import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import {MessageRuleImpl} from './index';
import {MessageRule} from './messageRule';
import {MessageRuleCollectionResponse} from './messageRuleCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleCollectionResponseImpl implements MessageRuleCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The nextLink property */
    private _nextLink?: string | undefined;
    /** The value property */
    private _value?: MessageRule[] | undefined;
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
     * Instantiates a new MessageRuleCollectionResponse and sets the default values.
     * @param messageRuleCollectionResponseParameterValue 
     */
    public constructor(messageRuleCollectionResponseParameterValue?: MessageRuleCollectionResponse | undefined) {
        this._additionalData = messageRuleCollectionResponseParameterValue?.additionalData ? messageRuleCollectionResponseParameterValue?.additionalData! : {};
        this._nextLink = messageRuleCollectionResponseParameterValue?.nextLink;
        this._value = messageRuleCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MessageRuleImpl>(createMessageRuleFromDiscriminatorValue); },
        };
    };
    /**
     * Gets the @odata.nextLink property value. The nextLink property
     * @returns a string
     */
    public get nextLink() {
        return this._nextLink;
    };
    /**
     * Sets the @odata.nextLink property value. The nextLink property
     * @param value Value to set for the nextLink property.
     */
    public set nextLink(value: string | undefined) {
        if(value) {
            this._nextLink = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.nextLink){
            writer.writeStringValue("@odata.nextLink", this.nextLink);
        }
        if(this.value && this.value.length != 0){        const valueArrValue: MessageRuleImpl[] = [];
        this.value?.forEach(element => {
            valueArrValue.push((element instanceof MessageRuleImpl? element as MessageRuleImpl:new MessageRuleImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MessageRuleImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the value property value. The value property
     * @returns a MessageRuleInterface
     */
    public get value() {
        return this._value;
    };
    /**
     * Sets the value property value. The value property
     * @param value Value to set for the value property.
     */
    public set value(value: MessageRule[] | undefined) {
        if(value) {
            const valueArrValue: MessageRuleImpl[] = [];
            this.value?.forEach(element => {
                valueArrValue.push((element instanceof MessageRuleImpl? element as MessageRuleImpl:new MessageRuleImpl(element)));
            });
            this._value = valueArrValue;
        }
    };
}
