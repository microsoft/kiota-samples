require './models'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Models
    class ItemBody
        include IAdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The content of the item.
        @content
        ## 
        # The contentType property
        @content_type
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
        ## Gets the content property value. The content of the item.
        ## @return a string
        ## 
        def  content
            return @content
        end
        ## 
        ## Sets the content property value. The content of the item.
        ## @param value Value to set for the content property.
        ## @return a void
        ## 
        def  content=(content)
            @content = content
        end
        ## 
        ## Gets the contentType property value. The contentType property
        ## @return a body_type
        ## 
        def  content_type
            return @content_type
        end
        ## 
        ## Sets the contentType property value. The contentType property
        ## @param value Value to set for the contentType property.
        ## @return a void
        ## 
        def  content_type=(contentType)
            @content_type = contentType
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a item_body
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
                "content" => lambda {|o, n| o.content = n.get_string_value() },
                "contentType" => lambda {|o, n| o.content_type = n.get_enum_value(Graphrubyv4::Models::BodyType) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_string_value("content", @content)
            writer.write_enum_value("contentType", @content_type)
            writer.write_additional_data(@additional_data)
        end
    end
end
