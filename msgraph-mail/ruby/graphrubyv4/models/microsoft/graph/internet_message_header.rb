require './models'
require 'microsoft_kiota_abstractions'

module GraphC::Models
    class InternetMessageHeader
        include IAdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # Represents the key in a key-value pair.
        @name
        ## 
        # The value in a key-value pair.
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
        ## @return a internet_message_header
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
                "name" => lambda {|o, n| o.name = n.get_string_value() },
                "value" => lambda {|o, n| o.value = n.get_string_value() },
            }
        end
        ## 
        ## Gets the name property value. Represents the key in a key-value pair.
        ## @return a string
        ## 
        def  name
            return @name
        end
        ## 
        ## Sets the name property value. Represents the key in a key-value pair.
        ## @param value Value to set for the name property.
        ## @return a void
        ## 
        def  name=(name)
            @name = name
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_string_value("name", @name)
            writer.write_string_value("value", @value)
            writer.write_additional_data(@additional_data)
        end
        ## 
        ## Gets the value property value. The value in a key-value pair.
        ## @return a string
        ## 
        def  value
            return @value
        end
        ## 
        ## Sets the value property value. The value in a key-value pair.
        ## @param value Value to set for the value property.
        ## @return a void
        ## 
        def  value=(value)
            @value = value
        end
    end
end
