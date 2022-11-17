<<<<<<< HEAD
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface EmailAddress extends Partial<AdditionalDataHolder>, Partial<Parsable> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown> | undefined;
    /** The email address of an entity instance. */
    address?: string | undefined;
    /** The display name of an entity instance. */
    name?: string | undefined;
=======
import {EmailAddress} from './emailAddress';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class EmailAddressImpl implements EmailAddress {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The email address of an entity instance. */
    private _address?: string | undefined;
    /** The display name of an entity instance. */
    private _name?: string | undefined;
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
     * Gets the address property value. The email address of an entity instance.
     * @returns a string
     */
    public get address() {
        return this._address;
    };
    /**
     * Sets the address property value. The email address of an entity instance.
     * @param value Value to set for the address property.
     */
    public set address(value: string | undefined) {
        if(value) {
            this._address = value;
        }
    };
    /**
     * Instantiates a new emailAddress and sets the default values.
     * @param emailAddressParameterValue 
     */
    public constructor(emailAddressParameterValue?: EmailAddress | undefined) {
        this._additionalData = emailAddressParameterValue?.additionalData ? emailAddressParameterValue?.additionalData! : {};
        this._address = emailAddressParameterValue?.address;
        this._name = emailAddressParameterValue?.name;
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
     * Gets the name property value. The display name of an entity instance.
     * @returns a string
     */
    public get name() {
        return this._name;
    };
    /**
     * Sets the name property value. The display name of an entity instance.
     * @param value Value to set for the name property.
     */
    public set name(value: string | undefined) {
        if(value) {
            this._name = value;
        }
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
>>>>>>> main
}
