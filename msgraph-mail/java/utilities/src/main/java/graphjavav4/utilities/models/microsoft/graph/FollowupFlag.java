package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class FollowupFlag implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private DateTimeTimeZone _completedDateTime;
    private DateTimeTimeZone _dueDateTime;
    private FollowupFlagStatus _flagStatus;
    private DateTimeTimeZone _startDateTime;
    /**
     * Instantiates a new followupFlag and sets the default values.
     * @return a void
     */
    public FollowupFlag() {
        this.setAdditionalData(new HashMap<>());
    }
    @javax.annotation.Nonnull
    public static FollowupFlag createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new FollowupFlag();
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the completedDateTime property value. 
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getCompletedDateTime() {
        return this._completedDateTime;
    }
    /**
     * Gets the dueDateTime property value. 
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getDueDateTime() {
        return this._dueDateTime;
    }
    /**
     * Gets the flagStatus property value. 
     * @return a followupFlagStatus
     */
    @javax.annotation.Nullable
    public FollowupFlagStatus getFlagStatus() {
        return this._flagStatus;
    }
    /**
     * Gets the startDateTime property value. 
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getStartDateTime() {
        return this._startDateTime;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(4) {{
            this.put("completedDateTime", (o, n) -> { ((FollowupFlag)o).setCompletedDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
            this.put("dueDateTime", (o, n) -> { ((FollowupFlag)o).setDueDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
            this.put("flagStatus", (o, n) -> { ((FollowupFlag)o).setFlagStatus(n.getEnumValue(FollowupFlagStatus.class)); });
            this.put("startDateTime", (o, n) -> { ((FollowupFlag)o).setStartDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeObjectValue("completedDateTime", this.getCompletedDateTime());
        writer.writeObjectValue("dueDateTime", this.getDueDateTime());
        writer.writeEnumValue("flagStatus", this.getFlagStatus());
        writer.writeObjectValue("startDateTime", this.getStartDateTime());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the completedDateTime property value. 
     * @param value Value to set for the completedDateTime property.
     * @return a void
     */
    public void setCompletedDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._completedDateTime = value;
    }
    /**
     * Sets the dueDateTime property value. 
     * @param value Value to set for the dueDateTime property.
     * @return a void
     */
    public void setDueDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._dueDateTime = value;
    }
    /**
     * Sets the flagStatus property value. 
     * @param value Value to set for the flagStatus property.
     * @return a void
     */
    public void setFlagStatus(@javax.annotation.Nullable final FollowupFlagStatus value) {
        this._flagStatus = value;
    }
    /**
     * Sets the startDateTime property value. 
     * @param value Value to set for the startDateTime property.
     * @return a void
     */
    public void setStartDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._startDateTime = value;
    }
}
