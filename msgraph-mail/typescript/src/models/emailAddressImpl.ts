import {EmailAddress} from './emailAddress';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class EmailAddressImpl implements EmailAddress {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The email address of an entity instance. */
    public address?: string | undefined;
    /** The display name of an entity instance. */
    public name?: string | undefined;
    /**
     * Instantiates a new emailAddress and sets the default values.
     * @param emailAddressParameterValue 
     */
    public constructor(emailAddressParameterValue?: EmailAddress | undefined) {
        this.additionalData = emailAddressParameterValue?.additionalData ? emailAddressParameterValue?.additionalData! : {};
        this.address = emailAddressParameterValue?.address;
        this.name = emailAddressParameterValue?.name;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "address": n => { this.address = n.getStringValue(); },
            "name": n => { this.name = n.getStringValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.address){
            writer.writeStringValue("address", this.address);
        }
        if(this.name){
            writer.writeStringValue("name", this.name);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
