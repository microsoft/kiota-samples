import {createExtensionFromDiscriminatorValue} from './createExtensionFromDiscriminatorValue';
import {Extension} from './extension';
import {ExtensionCollectionResponse} from './extensionCollectionResponse';
import {ExtensionImpl} from './index';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class ExtensionCollectionResponseImpl implements ExtensionCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: Extension[] | undefined;
    /**
     * Instantiates a new ExtensionCollectionResponse and sets the default values.
     * @param extensionCollectionResponseParameterValue 
     */
    public constructor(extensionCollectionResponseParameterValue?: ExtensionCollectionResponse | undefined) {
        this.additionalData = extensionCollectionResponseParameterValue?.additionalData ? extensionCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = extensionCollectionResponseParameterValue?.nextLink;
        this.value = extensionCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<ExtensionImpl>(createExtensionFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: ExtensionImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new ExtensionImpl(element));});
            writer.writeCollectionOfObjectValues<ExtensionImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
