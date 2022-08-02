require './models'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Models
    class ExtensionCollectionResponse
        include IAdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The OdataNextLink property
        @odata_next_link
        ## 
        # The value property
        @value
        ## 
        ## Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = additionalData
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a extension_collection_response
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "@odata.nextLink" => lambda {|o, n| o.odata_next_link = n.get_string_value() },
                "value" => lambda {|o, n| o.value = n.get_collection_of_object_values(Graphrubyv4::Models::Extension) },
            }
        end
        ## 
        ## Gets the @odata.nextLink property value. The OdataNextLink property
        ## @return a string
        ## 
        def  odata_next_link
            return @odata_next_link
        end
        ## 
        ## Sets the @odata.nextLink property value. The OdataNextLink property
        ## @param value Value to set for the OdataNextLink property.
        ## @return a void
        ## 
        def  odata_next_link=(odataNextLink)
            @odata_next_link = odataNextLink
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_string_value("@odata.nextLink", @odata_next_link)
            writer.write_collection_of_object_values("value", @value)
            writer.write_additional_data(@additional_data)
        end
        ## 
        ## Gets the value property value. The value property
        ## @return a extension
        ## 
        def  value
            return @value
        end
        ## 
        ## Sets the value property value. The value property
        ## @param value Value to set for the value property.
        ## @return a void
        ## 
        def  value=(value)
            @value = value
        end
    end
end
