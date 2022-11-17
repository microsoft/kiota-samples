require 'microsoft_kiota_abstractions'
require_relative './models'

module Graph::Models
    class User
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The displayName property
        @display_name
        ## 
        # The id property
        @id
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
        ## @return a user
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## Gets the displayName property value. The displayName property
        ## @return a string
        ## 
        def  display_name
            return @display_name
        end
        ## 
        ## Sets the displayName property value. The displayName property
        ## @param value Value to set for the displayName property.
        ## @return a void
        ## 
        def  display_name=(displayName)
            @display_name = displayName
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "displayName" => lambda {|n| @display_name = n.get_string_value() },
                "id" => lambda {|n| @id = n.get_string_value() },
            }
        end
        ## 
        ## Gets the id property value. The id property
        ## @return a string
        ## 
        def  id
            return @id
        end
        ## 
        ## Sets the id property value. The id property
        ## @param value Value to set for the id property.
        ## @return a void
        ## 
        def  id=(id)
            @id = id
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_string_value("displayName", @display_name)
            writer.write_string_value("id", @id)
            writer.write_additional_data(@additional_data)
        end
    end
end
