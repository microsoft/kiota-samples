require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './users'

module Graphrubyv4::Users
    class Attachment < Graphrubyv4::Users::Entity
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
        # The display name of the attachment. This does not need to be the actual file name.
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
        ## Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
        ## @return a boolean
        ## 
        def  is_inline
            return @is_inline
        end
        ## 
        ## Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
        ## @return a date_time_offset
        ## 
        def  last_modified_date_time
            return @last_modified_date_time
        end
        ## 
        ## Gets the name property value. The display name of the attachment. This does not need to be the actual file name.
        ## @return a string
        ## 
        def  name
            return @name
        end
        ## 
        ## Gets the size property value. The length of the attachment in bytes.
        ## @return a integer
        ## 
        def  size
            return @size
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "contentType" => lambda {|o, n| o.content_type = n.get_string_value() },
                "isInline" => lambda {|o, n| o.is_inline = n.get_boolean_value() },
                "lastModifiedDateTime" => lambda {|o, n| o.last_modified_date_time = n.get_date_value() },
                "name" => lambda {|o, n| o.name = n.get_string_value() },
                "size" => lambda {|o, n| o.size = n.get_number_value() },
            })
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
            writer.write_date_value("lastModifiedDateTime", @last_modified_date_time)
            writer.write_string_value("name", @name)
            writer.write_number_value("size", @size)
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
        ## Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
        ## @param value Value to set for the isInline property.
        ## @return a void
        ## 
        def  is_inline=(isInline)
            @is_inline = isInline
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
        ## Sets the name property value. The display name of the attachment. This does not need to be the actual file name.
        ## @param value Value to set for the name property.
        ## @return a void
        ## 
        def  name=(name)
            @name = name
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
