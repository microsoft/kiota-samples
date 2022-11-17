require 'microsoft_kiota_abstractions'
require_relative './models'

module Graphrubyv4::Models
    class Recipient
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The emailAddress property
        @email_address
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
        ## @return a recipient
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## Gets the emailAddress property value. The emailAddress property
        ## @return a email_address
        ## 
        def  email_address
            return @email_address
        end
        ## 
        ## Sets the emailAddress property value. The emailAddress property
        ## @param value Value to set for the emailAddress property.
        ## @return a void
        ## 
        def  email_address=(emailAddress)
            @email_address = emailAddress
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "emailAddress" => lambda {|n| @email_address = n.get_object_value(Graphrubyv4::Models::EmailAddress) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_object_value("emailAddress", @email_address)
            writer.write_additional_data(@additional_data)
        end
    end
end
