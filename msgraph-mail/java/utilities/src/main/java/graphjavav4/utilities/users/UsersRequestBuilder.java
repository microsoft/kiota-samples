package graphjavav4.utilities.users;

import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
/** Builds and executes requests for operations under /users  */
public class UsersRequestBuilder {
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /** Url template parameters for the request  */
    private final HashMap<String, String> urlTemplateParameters;
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @param urlTemplateParameters Url template parameters for the request
     * @return a void
     */
    public UsersRequestBuilder(@javax.annotation.Nonnull final HashMap<String, String> urlTemplateParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(requestAdapter);
        Objects.requireNonNull(urlTemplateParameters);
        this.urlTemplate = "https://graph.microsoft.com/v1.0/users";
        var urlTplParams = new HashMap<String, String>(urlTemplateParameters);
        this.urlTemplateParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public UsersRequestBuilder(@javax.annotation.Nonnull final String rawUrl, final RequestAdapter requestAdapter) {
        this.urlTemplate = "https://graph.microsoft.com/v1.0/users";
        var urlTplParams = new HashMap<String, String>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.urlTemplateParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
}
