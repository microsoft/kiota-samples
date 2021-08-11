require 'microsoft_kiota_abstractions'
require_relative './users'

module Graphrubyv4::Users
    class InternetMessageHeader
        include MicrosoftKiotaAbstractions::Parsable
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
        ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Gets the name property value. Represents the key in a key-value pair.
        ## @return a string
        ## 
        def  name
            return @name
        end
        ## 
        ## Gets the value property value. The value in a key-value pair.
        ## @return a string
        ## 
        def  value
            return @value
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
        ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = additionalData
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
        ## Sets the value property value. The value in a key-value pair.
        ## @param value Value to set for the value property.
        ## @return a void
        ## 
        def  value=(value)
            @value = value
        end
    end
end
