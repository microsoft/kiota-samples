package graphjavav4.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.Consumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class FollowupFlag implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private Map<String, Object> _additionalData;
    /** The completedDateTime property */
    private DateTimeTimeZone _completedDateTime;
    /** The dueDateTime property */
    private DateTimeTimeZone _dueDateTime;
    /** The flagStatus property */
    private FollowupFlagStatus _flagStatus;
    /** The startDateTime property */
    private DateTimeTimeZone _startDateTime;
    /**
     * Instantiates a new followupFlag and sets the default values.
     * @return a void
     */
    public FollowupFlag() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a followupFlag
     */
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
     * Gets the completedDateTime property value. The completedDateTime property
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getCompletedDateTime() {
        return this._completedDateTime;
    }
    /**
     * Gets the dueDateTime property value. The dueDateTime property
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getDueDateTime() {
        return this._dueDateTime;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, Consumer<ParseNode>> getFieldDeserializers() {
        final FollowupFlag currentObject = this;
        return new HashMap<>(4) {{
            this.put("completedDateTime", (n) -> { currentObject.setCompletedDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
            this.put("dueDateTime", (n) -> { currentObject.setDueDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
            this.put("flagStatus", (n) -> { currentObject.setFlagStatus(n.getEnumValue(FollowupFlagStatus.class)); });
            this.put("startDateTime", (n) -> { currentObject.setStartDateTime(n.getObjectValue(DateTimeTimeZone::createFromDiscriminatorValue)); });
        }};
    }
    /**
     * Gets the flagStatus property value. The flagStatus property
     * @return a followupFlagStatus
     */
    @javax.annotation.Nullable
    public FollowupFlagStatus getFlagStatus() {
        return this._flagStatus;
    }
    /**
     * Gets the startDateTime property value. The startDateTime property
     * @return a dateTimeTimeZone
     */
    @javax.annotation.Nullable
    public DateTimeTimeZone getStartDateTime() {
        return this._startDateTime;
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
     * Sets the completedDateTime property value. The completedDateTime property
     * @param value Value to set for the completedDateTime property.
     * @return a void
     */
    public void setCompletedDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._completedDateTime = value;
    }
    /**
     * Sets the dueDateTime property value. The dueDateTime property
     * @param value Value to set for the dueDateTime property.
     * @return a void
     */
    public void setDueDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._dueDateTime = value;
    }
    /**
     * Sets the flagStatus property value. The flagStatus property
     * @param value Value to set for the flagStatus property.
     * @return a void
     */
    public void setFlagStatus(@javax.annotation.Nullable final FollowupFlagStatus value) {
        this._flagStatus = value;
    }
    /**
     * Sets the startDateTime property value. The startDateTime property
     * @param value Value to set for the startDateTime property.
     * @return a void
     */
    public void setStartDateTime(@javax.annotation.Nullable final DateTimeTimeZone value) {
        this._startDateTime = value;
    }
}
