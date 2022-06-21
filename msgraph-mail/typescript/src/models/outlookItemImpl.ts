import {EntityImpl} from './index';
import {OutlookItem} from './outlookItem';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class OutlookItemImpl extends EntityImpl implements OutlookItem {
    /** The categories associated with the item */
    public categories?: string[] | undefined;
    /** Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. */
    public changeKey?: string | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    public createdDateTime?: Date | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    public lastModifiedDateTime?: Date | undefined;
    /**
     * Instantiates a new outlookItem and sets the default values.
     * @param outlookItemParameterValue 
     */
    public constructor(outlookItemParameterValue?: OutlookItem | undefined) {
        super(outlookItemParameterValue);
        this.categories = outlookItemParameterValue?.categories;
        this.changeKey = outlookItemParameterValue?.changeKey;
        this.createdDateTime = outlookItemParameterValue?.createdDateTime;
        this.lastModifiedDateTime = outlookItemParameterValue?.lastModifiedDateTime;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "categories": n => { this.categories = n.getCollectionOfPrimitiveValues<string>(); },
            "changeKey": n => { this.changeKey = n.getStringValue(); },
            "createdDateTime": n => { this.createdDateTime = n.getDateValue(); },
            "lastModifiedDateTime": n => { this.lastModifiedDateTime = n.getDateValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.categories){
            writer.writeCollectionOfPrimitiveValues<string>("categories", this.categories);
        }
        if(this.changeKey){
            writer.writeStringValue("changeKey", this.changeKey);
        }
        if(this.createdDateTime){
            writer.writeDateValue("createdDateTime", this.createdDateTime);
        }
        if(this.lastModifiedDateTime){
            writer.writeDateValue("lastModifiedDateTime", this.lastModifiedDateTime);
        }
    };
}
