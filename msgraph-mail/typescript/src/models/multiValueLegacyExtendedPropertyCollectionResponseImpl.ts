import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {MultiValueLegacyExtendedPropertyImpl} from './index';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {MultiValueLegacyExtendedPropertyCollectionResponse} from './multiValueLegacyExtendedPropertyCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MultiValueLegacyExtendedPropertyCollectionResponseImpl implements MultiValueLegacyExtendedPropertyCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The OdataNextLink property */
    private _odataNextLink?: string | undefined;
    /** The value property */
    private _value?: MultiValueLegacyExtendedProperty[] | undefined;
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
     * Instantiates a new MultiValueLegacyExtendedPropertyCollectionResponse and sets the default values.
     * @param multiValueLegacyExtendedPropertyCollectionResponseParameterValue 
     */
    public constructor(multiValueLegacyExtendedPropertyCollectionResponseParameterValue?: MultiValueLegacyExtendedPropertyCollectionResponse | undefined) {
        this._additionalData = {};
        this._additionalData = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData ? multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData! : {};
        this._odataNextLink = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.odataNextLink;
        this._value = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.odataNextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
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
    public set odataNextLink(value: string | undefined) {
        if(value) {
            this._odataNextLink = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.odataNextLink){
            writer.writeStringValue("@odata.nextLink", this.odataNextLink);
        }
        if(this.value && this.value.length != 0){        const valueArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
        this.value?.forEach(element => {
            valueArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData!);
    };
    /**
     * Gets the value property value. The value property
     * @returns a MultiValueLegacyExtendedPropertyInterface
     */
    public get value() {
        return this._value;
    };
    /**
     * Sets the value property value. The value property
     * @param value Value to set for the value property.
     */
    public set value(value: MultiValueLegacyExtendedProperty[] | undefined) {
        if(value) {
            const valueArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
            this.value?.forEach(element => {
                valueArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
            });
            this._value = valueArrValue;
        }
    };
}
