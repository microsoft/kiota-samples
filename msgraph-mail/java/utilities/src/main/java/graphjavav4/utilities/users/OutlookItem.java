package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class OutlookItem extends Entity implements Parsable {
    /** The categories associated with the item  */
    private List<String> _categories;
    /** Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.  */
    private String _changeKey;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    private OffsetDateTime _createdDateTime;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    private OffsetDateTime _lastModifiedDateTime;
    /**
     * Instantiates a new outlookItem and sets the default values.
     * @return a void
     */
    public OutlookItem() {
        super();
    }
    /**
     * Gets the categories property value. The categories associated with the item
     * @return a string
     */
    @javax.annotation.Nullable
    public List<String> getCategories() {
        return this._categories;
    }
    /**
     * Gets the changeKey property value. Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getChangeKey() {
        return this._changeKey;
    }
    /**
     * Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @return a OffsetDateTime
     */
    @javax.annotation.Nullable
    public OffsetDateTime getCreatedDateTime() {
        return this._createdDateTime;
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
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("categories", (o, n) -> { ((OutlookItem)o).setCategories(n.getCollectionOfPrimitiveValues(String.class)); });
            this.put("changeKey", (o, n) -> { ((OutlookItem)o).setChangeKey(n.getStringValue()); });
            this.put("createdDateTime", (o, n) -> { ((OutlookItem)o).setCreatedDateTime(n.getOffsetDateTimeValue()); });
            this.put("lastModifiedDateTime", (o, n) -> { ((OutlookItem)o).setLastModifiedDateTime(n.getOffsetDateTimeValue()); });
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
        writer.writeCollectionOfPrimitiveValues("categories", this.getCategories());
        writer.writeStringValue("changeKey", this.getChangeKey());
        writer.writeOffsetDateTimeValue("createdDateTime", this.getCreatedDateTime());
        writer.writeOffsetDateTimeValue("lastModifiedDateTime", this.getLastModifiedDateTime());
    }
    /**
     * Sets the categories property value. The categories associated with the item
     * @param value Value to set for the categories property.
     * @return a void
     */
    public void setCategories(@javax.annotation.Nullable final List<String> value) {
        this._categories = value;
    }
    /**
     * Sets the changeKey property value. Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     * @param value Value to set for the changeKey property.
     * @return a void
     */
    public void setChangeKey(@javax.annotation.Nullable final String value) {
        this._changeKey = value;
    }
    /**
     * Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @param value Value to set for the createdDateTime property.
     * @return a void
     */
    public void setCreatedDateTime(@javax.annotation.Nullable final OffsetDateTime value) {
        this._createdDateTime = value;
    }
    /**
     * Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @param value Value to set for the lastModifiedDateTime property.
     * @return a void
     */
    public void setLastModifiedDateTime(@javax.annotation.Nullable final OffsetDateTime value) {
        this._lastModifiedDateTime = value;
    }
}
