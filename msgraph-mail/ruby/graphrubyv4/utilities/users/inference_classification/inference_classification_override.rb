require 'microsoft_kiota_abstractions'
require_relative '../entity'

module Graphrubyv4::Utilities::Users::InferenceClassification
    class InferenceClassificationOverride < Graphrubyv4::Utilities::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
        @classify_as
        @sender_email_address
        ## 
        ## Gets the classifyAs property value. 
        ## @return a inference_classification_type
        ## 
        def  classify_as
            return @classify_as
        end
        ## 
        ## Gets the senderEmailAddress property value. 
        ## @return a email_address
        ## 
        def  sender_email_address
            return @sender_email_address
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "classify_as" => lambda {|o, n| o.classify_as = n.get_enum_value(inference_classification_type) },
                "sender_email_address" => lambda {|o, n| o.sender_email_address = n.get_object_value(email_address) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super.serialize(writer)
            writer.write_enum_value("classify_as", self.classify_as)
            writer.write_object_value("sender_email_address", self.sender_email_address)
        end
        ## 
        ## Sets the classifyAs property value. 
        ## @param value Value to set for the classifyAs property.
        ## @return a void
        ## 
        def  classify_as=(classifyAs)
            @classify_as = (classifyAs)
        end
        ## 
        ## Sets the senderEmailAddress property value. 
        ## @param value Value to set for the senderEmailAddress property.
        ## @return a void
        ## 
        def  sender_email_address=(senderEmailAddress)
            @sender_email_address = (senderEmailAddress)
        end
    end
end
