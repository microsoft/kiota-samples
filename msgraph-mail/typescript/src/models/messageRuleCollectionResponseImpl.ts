import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
<<<<<<< HEAD:msgraph-mail/typescript/src/models/messageRuleCollectionResponseImpl.ts
import {MessageRuleImpl} from './index';
import {MessageRule} from './messageRule';
import {MessageRuleCollectionResponse} from './messageRuleCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleCollectionResponseImpl implements MessageRuleCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The nextLink property */
    private _nextLink?: string | undefined;
=======
import {MessageRule} from './index';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleCollectionResponse implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The OdataNextLink property */
    private _odataNextLink?: string | undefined;
>>>>>>> main:msgraph-mail/typescript/src/models/messageRuleCollectionResponse.ts
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
<<<<<<< HEAD:msgraph-mail/typescript/src/models/messageRuleCollectionResponseImpl.ts
     * @param messageRuleCollectionResponseParameterValue 
=======
>>>>>>> main:msgraph-mail/typescript/src/models/messageRuleCollectionResponse.ts
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
<<<<<<< HEAD:msgraph-mail/typescript/src/models/messageRuleCollectionResponseImpl.ts
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MessageRuleImpl>(createMessageRuleFromDiscriminatorValue); },
=======
            "@odata.nextLink": n => { this.odataNextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MessageRule>(createMessageRuleFromDiscriminatorValue); },
>>>>>>> main:msgraph-mail/typescript/src/models/messageRuleCollectionResponse.ts
        };
    };
    /**
     * Gets the @odata.nextLink property value. The OdataNextLink property
     * @returns a string
     */
    public get odataNextLink() {
        return this._odataNextLink;
    };
    /**
     * Sets the @odata.nextLink property value. The OdataNextLink property
     * @param value Value to set for the OdataNextLink property.
     */
<<<<<<< HEAD:msgraph-mail/typescript/src/models/messageRuleCollectionResponseImpl.ts
    public set nextLink(value: string | undefined) {
        if(value) {
            this._nextLink = value;
        }
=======
    public set odataNextLink(value: string | undefined) {
        this._odataNextLink = value;
>>>>>>> main:msgraph-mail/typescript/src/models/messageRuleCollectionResponse.ts
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
<<<<<<< HEAD:msgraph-mail/typescript/src/models/messageRuleCollectionResponseImpl.ts
        if(this.nextLink){
            writer.writeStringValue("@odata.nextLink", this.nextLink);
        }
        if(this.value && this.value.length != 0){        const valueArrValue: MessageRuleImpl[] = [];
        this.value?.forEach(element => {
            valueArrValue.push((element instanceof MessageRuleImpl? element as MessageRuleImpl:new MessageRuleImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MessageRuleImpl>("value", valueArrValue);
        }
=======
        writer.writeStringValue("@odata.nextLink", this.odataNextLink);
        writer.writeCollectionOfObjectValues<MessageRule>("value", this.value);
>>>>>>> main:msgraph-mail/typescript/src/models/messageRuleCollectionResponse.ts
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
