import {Extension} from './extension';
import {EntityImpl} from './index';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class ExtensionImpl extends EntityImpl implements Extension {
    /**
<<<<<<< HEAD:msgraph-mail/typescript/src/models/extensionImpl.ts
     * Instantiates a new extension and sets the default values.
     * @param extensionParameterValue 
=======
     * Instantiates a new Extension and sets the default values.
>>>>>>> main:msgraph-mail/typescript/src/models/extension.ts
     */
    public constructor(extensionParameterValue?: Extension | undefined) {
        super(extensionParameterValue);
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
    };
}
