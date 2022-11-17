import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {MailFolderImpl} from './index';
import {MailFolder} from './mailFolder';
import {MailFolderCollectionResponse} from './mailFolderCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MailFolderCollectionResponseImpl implements MailFolderCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData?: Record<string, unknown> | undefined;
    /** The OdataNextLink property */
    private _odataNextLink?: string | undefined;
    /** The value property */
    private _value?: MailFolder[] | undefined;
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
     * Instantiates a new MailFolderCollectionResponse and sets the default values.
     * @param mailFolderCollectionResponseParameterValue 
     */
    public constructor(mailFolderCollectionResponseParameterValue?: MailFolderCollectionResponse | undefined) {
        this._additionalData = {};
        this._additionalData = mailFolderCollectionResponseParameterValue?.additionalData ? mailFolderCollectionResponseParameterValue?.additionalData! : {};
        this._odataNextLink = mailFolderCollectionResponseParameterValue?.odataNextLink;
        this._value = mailFolderCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.odataNextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MailFolderImpl>(createMailFolderFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: MailFolderImpl[] = [];
        this.value?.forEach(element => {
            valueArrValue.push((element instanceof MailFolderImpl? element as MailFolderImpl:new MailFolderImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MailFolderImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData!);
    };
    /**
     * Gets the value property value. The value property
     * @returns a MailFolderInterface
     */
    public get value() {
        return this._value;
    };
    /**
     * Sets the value property value. The value property
     * @param value Value to set for the value property.
     */
    public set value(value: MailFolder[] | undefined) {
        if(value) {
            const valueArrValue: MailFolderImpl[] = [];
            this.value?.forEach(element => {
                valueArrValue.push((element instanceof MailFolderImpl? element as MailFolderImpl:new MailFolderImpl(element)));
            });
            this._value = valueArrValue;
        }
    };
}
