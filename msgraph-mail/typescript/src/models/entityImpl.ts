import {Entity} from './entity';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class EntityImpl implements Entity {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
<<<<<<< HEAD:msgraph-mail/typescript/src/models/entityImpl.ts
    private _additionalData?: Record<string, unknown> | undefined;
    /** The id property */
=======
    private _additionalData: Record<string, unknown>;
    /** The unique idenfier for an entity. Read-only. */
>>>>>>> main:msgraph-mail/typescript/src/models/entity.ts
    private _id?: string | undefined;
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
     * Instantiates a new entity and sets the default values.
     * @param entityParameterValue 
     */
    public constructor(entityParameterValue?: Entity | undefined) {
        this._additionalData = entityParameterValue?.additionalData ? entityParameterValue?.additionalData! : {};
        this._id = entityParameterValue?.id;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "id": n => { this.id = n.getStringValue(); },
        };
    };
    /**
<<<<<<< HEAD:msgraph-mail/typescript/src/models/entityImpl.ts
     * Gets the id property value. The id property
=======
     * Gets the id property value. The unique idenfier for an entity. Read-only.
>>>>>>> main:msgraph-mail/typescript/src/models/entity.ts
     * @returns a string
     */
    public get id() {
        return this._id;
    };
    /**
<<<<<<< HEAD:msgraph-mail/typescript/src/models/entityImpl.ts
     * Sets the id property value. The id property
=======
     * Sets the id property value. The unique idenfier for an entity. Read-only.
>>>>>>> main:msgraph-mail/typescript/src/models/entity.ts
     * @param value Value to set for the id property.
     */
    public set id(value: string | undefined) {
        if(value) {
            this._id = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.id){
            writer.writeStringValue("id", this.id);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
