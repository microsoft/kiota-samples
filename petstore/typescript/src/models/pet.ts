import {createCategoryFromDiscriminatorValue} from './createCategoryFromDiscriminatorValue';
import {createTagFromDiscriminatorValue} from './createTagFromDiscriminatorValue';
import {Category, Tag} from './index';
import {Pet_status} from './pet_status';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Pet implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private _additionalData: Record<string, unknown>;
    /** The category property */
    private _category?: Category | undefined;
    /** The id property */
    private _id?: number | undefined;
    /** The name property */
    private _name?: string | undefined;
    /** The photoUrls property */
    private _photoUrls?: string[] | undefined;
    /** pet status in the store */
    private _status?: Pet_status | undefined;
    /** The tags property */
    private _tags?: Tag[] | undefined;
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
    public set additionalData(value: Record<string, unknown>) {
        this._additionalData = value;
    };
    /**
     * Gets the category property value. The category property
     * @returns a Category
     */
    public get category() {
        return this._category;
    };
    /**
     * Sets the category property value. The category property
     * @param value Value to set for the category property.
     */
    public set category(value: Category | undefined) {
        this._category = value;
    };
    /**
     * Instantiates a new Pet and sets the default values.
     */
    public constructor() {
        this._additionalData = {};
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "category": n => { this.category = n.getObjectValue<Category>(createCategoryFromDiscriminatorValue); },
            "id": n => { this.id = n.getNumberValue(); },
            "name": n => { this.name = n.getStringValue(); },
            "photoUrls": n => { this.photoUrls = n.getCollectionOfPrimitiveValues<string>(); },
            "status": n => { this.status = n.getEnumValue<Pet_status>(Pet_status); },
            "tags": n => { this.tags = n.getCollectionOfObjectValues<Tag>(createTagFromDiscriminatorValue); },
        };
    };
    /**
     * Gets the id property value. The id property
     * @returns a int64
     */
    public get id() {
        return this._id;
    };
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     */
    public set id(value: number | undefined) {
        this._id = value;
    };
    /**
     * Gets the name property value. The name property
     * @returns a string
     */
    public get name() {
        return this._name;
    };
    /**
     * Sets the name property value. The name property
     * @param value Value to set for the name property.
     */
    public set name(value: string | undefined) {
        this._name = value;
    };
    /**
     * Gets the photoUrls property value. The photoUrls property
     * @returns a string
     */
    public get photoUrls() {
        return this._photoUrls;
    };
    /**
     * Sets the photoUrls property value. The photoUrls property
     * @param value Value to set for the photoUrls property.
     */
    public set photoUrls(value: string[] | undefined) {
        this._photoUrls = value;
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        writer.writeObjectValue<Category>("category", this.category);
        writer.writeNumberValue("id", this.id);
        writer.writeStringValue("name", this.name);
        writer.writeCollectionOfPrimitiveValues<string>("photoUrls", this.photoUrls);
        writer.writeEnumValue<Pet_status>("status", this.status);
        writer.writeCollectionOfObjectValues<Tag>("tags", this.tags);
        writer.writeAdditionalData(this.additionalData);
    };
    /**
     * Gets the status property value. pet status in the store
     * @returns a Pet_status
     */
    public get status() {
        return this._status;
    };
    /**
     * Sets the status property value. pet status in the store
     * @param value Value to set for the status property.
     */
    public set status(value: Pet_status | undefined) {
        this._status = value;
    };
    /**
     * Gets the tags property value. The tags property
     * @returns a Tag
     */
    public get tags() {
        return this._tags;
    };
    /**
     * Sets the tags property value. The tags property
     * @param value Value to set for the tags property.
     */
    public set tags(value: Tag[] | undefined) {
        this._tags = value;
    };
}
