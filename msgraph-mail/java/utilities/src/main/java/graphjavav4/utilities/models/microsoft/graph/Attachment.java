package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Attachment extends Entity implements Parsable {
    /** The MIME type.  */
    private String _contentType;
    /** true if the attachment is an inline attachment; otherwise, false.  */
    private Boolean _isInline;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    private OffsetDateTime _lastModifiedDateTime;
    /** The display name of the attachment. This does not need to be the actual file name.  */
    private String _name;
    /** The length of the attachment in bytes.  */
    private Integer _size;
    /**
     * Instantiates a new attachment and sets the default values.
     * @return a void
     */
    public Attachment() {
        super();
    }
    /**
     * Gets the contentType property value. The MIME type.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getContentType() {
        return this._contentType;
    }
    /**
     * Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsInline() {
        return this._isInline;
    }
    /**
     * Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @return a OffsetDateTime
     */
    @javax.annotation.Nullable
    public OffsetDateTime getLastModifiedDateTime() {
        return this._lastModifiedDateTime;
    }
    /**
     * Gets the name property value. The display name of the attachment. This does not need to be the actual file name.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getName() {
        return this._name;
    }
    /**
     * Gets the size property value. The length of the attachment in bytes.
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getSize() {
        return this._size;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("contentType", (o, n) -> { ((Attachment)o).setContentType(n.getStringValue()); });
            this.put("isInline", (o, n) -> { ((Attachment)o).setIsInline(n.getBooleanValue()); });
            this.put("lastModifiedDateTime", (o, n) -> { ((Attachment)o).setLastModifiedDateTime(n.getOffsetDateTimeValue()); });
            this.put("name", (o, n) -> { ((Attachment)o).setName(n.getStringValue()); });
            this.put("size", (o, n) -> { ((Attachment)o).setSize(n.getIntegerValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeStringValue("contentType", this.getContentType());
        writer.writeBooleanValue("isInline", this.getIsInline());
        writer.writeOffsetDateTimeValue("lastModifiedDateTime", this.getLastModifiedDateTime());
        writer.writeStringValue("name", this.getName());
        writer.writeIntegerValue("size", this.getSize());
    }
    /**
     * Sets the contentType property value. The MIME type.
     * @param value Value to set for the contentType property.
     * @return a void
     */
    public void setContentType(@javax.annotation.Nullable final String value) {
        this._contentType = value;
    }
    /**
     * Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
     * @param value Value to set for the isInline property.
     * @return a void
     */
    public void setIsInline(@javax.annotation.Nullable final Boolean value) {
        this._isInline = value;
    }
    /**
     * Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @param value Value to set for the lastModifiedDateTime property.
     * @return a void
     */
    public void setLastModifiedDateTime(@javax.annotation.Nullable final OffsetDateTime value) {
        this._lastModifiedDateTime = value;
    }
    /**
     * Sets the name property value. The display name of the attachment. This does not need to be the actual file name.
     * @param value Value to set for the name property.
     * @return a void
     */
    public void setName(@javax.annotation.Nullable final String value) {
        this._name = value;
    }
    /**
     * Sets the size property value. The length of the attachment in bytes.
     * @param value Value to set for the size property.
     * @return a void
     */
    public void setSize(@javax.annotation.Nullable final Integer value) {
        this._size = value;
    }
}
