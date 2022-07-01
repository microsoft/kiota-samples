import {createEmailAddressFromDiscriminatorValue} from './createEmailAddressFromDiscriminatorValue';
import {EmailAddress} from './emailAddress';
import {EmailAddressImpl} from './index';
import {Recipient} from './recipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class RecipientImpl implements Recipient {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The emailAddress property */
    private _emailAddress?: EmailAddress | undefined;
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
     * Instantiates a new recipient and sets the default values.
     * @param recipientParameterValue 
     */
    public constructor(recipientParameterValue?: Recipient | undefined) {
        this._additionalData = recipientParameterValue?.additionalData ? recipientParameterValue?.additionalData! : {};
        this._emailAddress = recipientParameterValue?.emailAddress;
    };
    /**
     * Gets the emailAddress property value. The emailAddress property
     * @returns a EmailAddressInterface
     */
    public get emailAddress() {
        return this._emailAddress;
    };
    /**
     * Sets the emailAddress property value. The emailAddress property
     * @param value Value to set for the emailAddress property.
     */
    public set emailAddress(value: EmailAddress | undefined) {
        if(value) {
            this._emailAddress = value instanceof EmailAddressImpl? value as EmailAddressImpl: new EmailAddressImpl(value);
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "emailAddress": n => { this.emailAddress = n.getObjectValue<EmailAddressImpl>(createEmailAddressFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.emailAddress){
            writer.writeObjectValue<EmailAddressImpl>("emailAddress", (this.emailAddress instanceof EmailAddressImpl? this.emailAddress as EmailAddressImpl: new EmailAddressImpl(this.emailAddress)));
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
