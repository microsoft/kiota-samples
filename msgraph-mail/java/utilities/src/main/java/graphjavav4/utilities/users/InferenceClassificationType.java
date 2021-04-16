package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.ValuedEnum;

public enum InferenceClassificationType implements ValuedEnum {
    Focused("focused"),
    Other("other");
    public final String value;
    InferenceClassificationType(final String value) {
        this.value = value;
    }
    @javax.annotation.Nonnull
    public String getValue() { return this.value; }
    @javax.annotation.Nullable
    public static InferenceClassificationType forValue(@javax.annotation.Nonnull final String searchValue) {
        switch(searchValue) {
            case "focused": return Focused;
            case "other": return Other;
            default: return null;
        }
    }
}
