require 'microsoft_kiota_abstractions'
require_relative './models'

module Graphrubyv4::Models
    class EmailAddress
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The email address of the person or entity.
        @address
        ## 
        # The display name of the person or entity.
        @name
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
        ## Gets the address property value. The email address of the person or entity.
        ## @return a string
        ## 
        def  address
            return @address
        end
        ## 
        ## Sets the address property value. The email address of the person or entity.
        ## @param value Value to set for the address property.
        ## @return a void
        ## 
        def  address=(address)
            @address = address
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a email_address
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
                "address" => lambda {|n| @address = n.get_string_value() },
                "name" => lambda {|n| @name = n.get_string_value() },
            }
        end
        ## 
        ## Gets the name property value. The display name of the person or entity.
        ## @return a string
        ## 
        def  name
            return @name
        end
        ## 
        ## Sets the name property value. The display name of the person or entity.
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
            writer.write_string_value("address", @address)
            writer.write_string_value("name", @name)
            writer.write_additional_data(@additional_data)
        end
    end
end
