package petstore.utilities.models;

import com.microsoft.kiota.serialization.ValuedEnum;
import java.util.Objects;

/**
 * pet status in the store
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public enum PetStatus implements ValuedEnum {
    Available("available"),
    Pending("pending"),
    Sold("sold");
    public final String value;
    PetStatus(final String value) {
        this.value = value;
    }
    @jakarta.annotation.Nonnull
    public String getValue() { return this.value; }
    @jakarta.annotation.Nullable
    public static PetStatus forValue(@jakarta.annotation.Nonnull final String searchValue) {
        Objects.requireNonNull(searchValue);
        switch(searchValue) {
            case "available": return Available;
            case "pending": return Pending;
            case "sold": return Sold;
            default: return null;
        }
    }
}
