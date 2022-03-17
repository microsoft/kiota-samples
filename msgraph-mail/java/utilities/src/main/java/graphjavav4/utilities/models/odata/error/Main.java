package graphjavav4.utilities.models.odata.error;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Main implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private String _code;
    private java.util.List<Detail> _details;
    /** The structure of this object is service-specific  */
    private Innererror _innererror;
    private String _message;
    private String _target;
    /**
     * Instantiates a new main and sets the default values.
     * @return a void
     */
    public Main() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a main
     */
    @javax.annotation.Nonnull
    public static Main createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new Main();
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
     * Gets the code property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getCode() {
        return this._code;
    }
    /**
     * Gets the details property value. 
     * @return a detail
     */
    @javax.annotation.Nullable
    public java.util.List<Detail> getDetails() {
        return this._details;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(5) {{
            this.put("code", (o, n) -> { ((Main)o).setCode(n.getStringValue()); });
            this.put("details", (o, n) -> { ((Main)o).setDetails(n.getCollectionOfObjectValues(Detail::createFromDiscriminatorValue)); });
            this.put("innererror", (o, n) -> { ((Main)o).setInnererror(n.getObjectValue(Innererror::createFromDiscriminatorValue)); });
            this.put("message", (o, n) -> { ((Main)o).setMessage(n.getStringValue()); });
            this.put("target", (o, n) -> { ((Main)o).setTarget(n.getStringValue()); });
        }};
    }
    /**
     * Gets the innererror property value. The structure of this object is service-specific
     * @return a innererror
     */
    @javax.annotation.Nullable
    public Innererror getInnererror() {
        return this._innererror;
    }
    /**
     * Gets the message property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getMessage() {
        return this._message;
    }
    /**
     * Gets the target property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getTarget() {
        return this._target;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("code", this.getCode());
        writer.writeCollectionOfObjectValues("details", this.getDetails());
        writer.writeObjectValue("innererror", this.getInnererror());
        writer.writeStringValue("message", this.getMessage());
        writer.writeStringValue("target", this.getTarget());
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
     * Sets the code property value. 
     * @param value Value to set for the code property.
     * @return a void
     */
    public void setCode(@javax.annotation.Nullable final String value) {
        this._code = value;
    }
    /**
     * Sets the details property value. 
     * @param value Value to set for the details property.
     * @return a void
     */
    public void setDetails(@javax.annotation.Nullable final java.util.List<Detail> value) {
        this._details = value;
    }
    /**
     * Sets the innererror property value. The structure of this object is service-specific
     * @param value Value to set for the innererror property.
     * @return a void
     */
    public void setInnererror(@javax.annotation.Nullable final Innererror value) {
        this._innererror = value;
    }
    /**
     * Sets the message property value. 
     * @param value Value to set for the message property.
     * @return a void
     */
    public void setMessage(@javax.annotation.Nullable final String value) {
        this._message = value;
    }
    /**
     * Sets the target property value. 
     * @param value Value to set for the target property.
     * @return a void
     */
    public void setTarget(@javax.annotation.Nullable final String value) {
        this._target = value;
    }
}
