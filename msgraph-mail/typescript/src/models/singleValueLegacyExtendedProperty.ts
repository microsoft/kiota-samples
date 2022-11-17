<<<<<<< HEAD
import {Entity} from './entity';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface SingleValueLegacyExtendedProperty extends Entity, Partial<Parsable> {
    /** A property value. */
    value?: string | undefined;
=======
import {EntityImpl} from './index';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class SingleValueLegacyExtendedPropertyImpl extends EntityImpl implements SingleValueLegacyExtendedProperty {
    /** A property value. */
    private _value?: string | undefined;
    /**
<<<<<<< HEAD:msgraph-mail/typescript/src/models/singleValueLegacyExtendedPropertyImpl.ts
     * Instantiates a new singleValueLegacyExtendedProperty and sets the default values.
     * @param singleValueLegacyExtendedPropertyParameterValue 
=======
     * Instantiates a new SingleValueLegacyExtendedProperty and sets the default values.
>>>>>>> main:msgraph-mail/typescript/src/models/singleValueLegacyExtendedProperty.ts
     */
    public constructor(singleValueLegacyExtendedPropertyParameterValue?: SingleValueLegacyExtendedProperty | undefined) {
        super(singleValueLegacyExtendedPropertyParameterValue);
        this._value = singleValueLegacyExtendedPropertyParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "value": n => { this.value = n.getStringValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.value){
            writer.writeStringValue("value", this.value);
        }
    };
    /**
     * Gets the value property value. A property value.
     * @returns a string
     */
    public get value() {
        return this._value;
    };
    /**
     * Sets the value property value. A property value.
     * @param value Value to set for the value property.
     */
    public set value(value: string | undefined) {
        if(value) {
            this._value = value;
        }
    };
>>>>>>> main
}
