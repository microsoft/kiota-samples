import {createEmailAddressFromDiscriminatorValue} from './createEmailAddressFromDiscriminatorValue';
import {EmailAddress} from './emailAddress';
import {EmailAddressImpl} from './index';
import {Recipient} from './recipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class RecipientImpl implements Recipient {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The emailAddress property */
    public emailAddress?: EmailAddress | undefined;
    /**
     * Instantiates a new recipient and sets the default values.
     * @param recipientParameterValue 
     */
    public constructor(recipientParameterValue?: Recipient | undefined) {
        this.additionalData = recipientParameterValue?.additionalData ? recipientParameterValue?.additionalData! : {};
        this.emailAddress = recipientParameterValue?.emailAddress;
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
            writer.writeObjectValue<EmailAddressImpl>("emailAddress", new EmailAddressImpl(this.emailAddress));
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
