package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class DateTimeTimeZone implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.  */
    private String _dateTime;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.  */
    private String _timeZone;
    /**
     * Instantiates a new dateTimeTimeZone and sets the default values.
     * @return a void
     */
    public DateTimeTimeZone() {
        this.setAdditionalData(new HashMap<>());
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
     * Gets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getDateTime() {
        return this._dateTime;
    }
    /**
     * Gets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getTimeZone() {
        return this._timeZone;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("dateTime", (o, n) -> { ((DateTimeTimeZone)o).setDateTime(n.getStringValue()); });
            this.put("timeZone", (o, n) -> { ((DateTimeTimeZone)o).setTimeZone(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("dateTime", this.getDateTime());
        writer.writeStringValue("timeZone", this.getTimeZone());
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
     * Sets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
     * @param value Value to set for the dateTime property.
     * @return a void
     */
    public void setDateTime(@javax.annotation.Nullable final String value) {
        this._dateTime = value;
    }
    /**
     * Sets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
     * @param value Value to set for the timeZone property.
     * @return a void
     */
    public void setTimeZone(@javax.annotation.Nullable final String value) {
        this._timeZone = value;
    }
}
