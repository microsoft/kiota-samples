import {createMessageFromDiscriminatorValue} from './createMessageFromDiscriminatorValue';
import {MessageImpl} from './index';
import {Message} from './message';
import {MessageCollectionResponse} from './messageCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageCollectionResponseImpl implements MessageCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: Message[] | undefined;
    /**
     * Instantiates a new MessageCollectionResponse and sets the default values.
     * @param messageCollectionResponseParameterValue 
     */
    public constructor(messageCollectionResponseParameterValue?: MessageCollectionResponse | undefined) {
        this.additionalData = messageCollectionResponseParameterValue?.additionalData ? messageCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = messageCollectionResponseParameterValue?.nextLink;
        this.value = messageCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MessageImpl>(createMessageFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: MessageImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new MessageImpl(element));});
            writer.writeCollectionOfObjectValues<MessageImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
