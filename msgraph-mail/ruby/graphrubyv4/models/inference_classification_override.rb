require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './models'

module Graphrubyv4::Models
    class InferenceClassificationOverride < Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # The classifyAs property
        @classify_as
        ## 
        # The senderEmailAddress property
        @sender_email_address
        ## 
        ## Gets the classifyAs property value. The classifyAs property
        ## @return a inference_classification_type
        ## 
        def classify_as
            return @classify_as
        end
        ## 
        ## Sets the classifyAs property value. The classifyAs property
        ## @param value Value to set for the classifyAs property.
        ## @return a void
        ## 
        def classify_as=(classifyAs)
            @classify_as = classifyAs
        end
        ## 
        ## Instantiates a new InferenceClassificationOverride and sets the default values.
        ## @return a void
        ## 
        def initialize() 
            super
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a inference_classification_override
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "classifyAs" => lambda {|n| @classify_as = n.get_enum_value(Graphrubyv4::Models::InferenceClassificationType) },
                "senderEmailAddress" => lambda {|n| @sender_email_address = n.get_object_value(Graphrubyv4::Models::EmailAddress) },
            })
        end
        ## 
        ## Gets the senderEmailAddress property value. The senderEmailAddress property
        ## @return a email_address
        ## 
        def sender_email_address
            return @sender_email_address
        end
        ## 
        ## Sets the senderEmailAddress property value. The senderEmailAddress property
        ## @param value Value to set for the senderEmailAddress property.
        ## @return a void
        ## 
        def sender_email_address=(senderEmailAddress)
            @sender_email_address = senderEmailAddress
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
            writer.write_enum_value("classifyAs", @classify_as)
            writer.write_object_value("senderEmailAddress", @sender_email_address)
        end
    end
end
