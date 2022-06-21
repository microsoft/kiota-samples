import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {MailFolderImpl} from './index';
import {MailFolder} from './mailFolder';
import {MailFolderCollectionResponse} from './mailFolderCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MailFolderCollectionResponseImpl implements MailFolderCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: MailFolder[] | undefined;
    /**
     * Instantiates a new MailFolderCollectionResponse and sets the default values.
     * @param mailFolderCollectionResponseParameterValue 
     */
    public constructor(mailFolderCollectionResponseParameterValue?: MailFolderCollectionResponse | undefined) {
        this.additionalData = mailFolderCollectionResponseParameterValue?.additionalData ? mailFolderCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = mailFolderCollectionResponseParameterValue?.nextLink;
        this.value = mailFolderCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MailFolderImpl>(createMailFolderFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: MailFolderImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new MailFolderImpl(element));});
            writer.writeCollectionOfObjectValues<MailFolderImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
