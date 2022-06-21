import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import {MessageRuleImpl} from './index';
import {MessageRule} from './messageRule';
import {MessageRuleCollectionResponse} from './messageRuleCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleCollectionResponseImpl implements MessageRuleCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: MessageRule[] | undefined;
    /**
     * Instantiates a new MessageRuleCollectionResponse and sets the default values.
     * @param messageRuleCollectionResponseParameterValue 
     */
    public constructor(messageRuleCollectionResponseParameterValue?: MessageRuleCollectionResponse | undefined) {
        this.additionalData = messageRuleCollectionResponseParameterValue?.additionalData ? messageRuleCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = messageRuleCollectionResponseParameterValue?.nextLink;
        this.value = messageRuleCollectionResponseParameterValue?.value;
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
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.nextLink){
            writer.writeStringValue("@odata.nextLink", this.nextLink);
        }
        if(this.value && this.value.length != 0){        const valueArrValue: MessageRuleImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new MessageRuleImpl(element));});
            writer.writeCollectionOfObjectValues<MessageRuleImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
