require 'date'
require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './models'

module Graphrubyv4::Models
    class Attachment < Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # The MIME type.
        @content_type
        ## 
        # true if the attachment is an inline attachment; otherwise, false.
        @is_inline
        ## 
        # The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
        @last_modified_date_time
        ## 
        # The attachment's file name.
        @name
        ## 
        # The length of the attachment in bytes.
        @size
        ## 
        ## Gets the contentType property value. The MIME type.
        ## @return a string
        ## 
        def  content_type
            return @content_type
        end
        ## 
        ## Sets the contentType property value. The MIME type.
        ## @param value Value to set for the contentType property.
        ## @return a void
        ## 
        def  content_type=(contentType)
            @content_type = contentType
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a attachment
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
                "contentType" => lambda {|n| @content_type = n.get_string_value() },
                "isInline" => lambda {|n| @is_inline = n.get_boolean_value() },
                "lastModifiedDateTime" => lambda {|n| @last_modified_date_time = n.get_date_time_value() },
                "name" => lambda {|n| @name = n.get_string_value() },
                "size" => lambda {|n| @size = n.get_number_value() },
            })
        end
        ## 
        ## Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
        ## @return a boolean
        ## 
        def  is_inline
            return @is_inline
        end
        ## 
        ## Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
        ## @param value Value to set for the isInline property.
        ## @return a void
        ## 
        def  is_inline=(isInline)
            @is_inline = isInline
        end
        ## 
        ## Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
        ## @return a date_time
        ## 
        def  last_modified_date_time
            return @last_modified_date_time
        end
        ## 
        ## Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
        ## @param value Value to set for the lastModifiedDateTime property.
        ## @return a void
        ## 
        def  last_modified_date_time=(lastModifiedDateTime)
            @last_modified_date_time = lastModifiedDateTime
        end
        ## 
        ## Gets the name property value. The attachment's file name.
        ## @return a string
        ## 
        def  name
            return @name
        end
        ## 
        ## Sets the name property value. The attachment's file name.
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
            super
            writer.write_string_value("contentType", @content_type)
            writer.write_boolean_value("isInline", @is_inline)
            writer.write_date_time_value("lastModifiedDateTime", @last_modified_date_time)
            writer.write_string_value("name", @name)
            writer.write_number_value("size", @size)
        end
        ## 
        ## Gets the size property value. The length of the attachment in bytes.
        ## @return a integer
        ## 
        def  size
            return @size
        end
        ## 
        ## Sets the size property value. The length of the attachment in bytes.
        ## @param value Value to set for the size property.
        ## @return a void
        ## 
        def  size=(size)
            @size = size
        end
    end
end
