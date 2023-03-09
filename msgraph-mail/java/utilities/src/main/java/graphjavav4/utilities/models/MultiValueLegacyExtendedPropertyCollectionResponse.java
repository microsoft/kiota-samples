package graphjavav4.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class MultiValueLegacyExtendedPropertyCollectionResponse implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private Map<String, Object> additionalData;
    /** The OdataNextLink property */
    private String odataNextLink;
    /** The value property */
    private java.util.List<MultiValueLegacyExtendedProperty> value;
    /**
     * Instantiates a new MultiValueLegacyExtendedPropertyCollectionResponse and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public MultiValueLegacyExtendedPropertyCollectionResponse() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a MultiValueLegacyExtendedPropertyCollectionResponse
     */
    @javax.annotation.Nonnull
    public static MultiValueLegacyExtendedPropertyCollectionResponse createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new MultiValueLegacyExtendedPropertyCollectionResponse();
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(2);
        deserializerMap.put("@odata.nextLink", (n) -> { this.setOdataNextLink(n.getStringValue()); });
        deserializerMap.put("value", (n) -> { this.setValue(n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the @odata.nextLink property value. The OdataNextLink property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getOdataNextLink() {
        return this.odataNextLink;
    }
    /**
     * Gets the value property value. The value property
     * @return a multiValueLegacyExtendedProperty
     */
    @javax.annotation.Nullable
    public java.util.List<MultiValueLegacyExtendedProperty> getValue() {
        return this.value;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    @javax.annotation.Nonnull
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("@odata.nextLink", this.getOdataNextLink());
        writer.writeCollectionOfObjectValues("value", this.getValue());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the @odata.nextLink property value. The OdataNextLink property
     * @param value Value to set for the OdataNextLink property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setOdataNextLink(@javax.annotation.Nullable final String value) {
        this.odataNextLink = value;
    }
    /**
     * Sets the value property value. The value property
     * @param value Value to set for the value property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setValue(@javax.annotation.Nullable final java.util.List<MultiValueLegacyExtendedProperty> value) {
        this.value = value;
    }
}
