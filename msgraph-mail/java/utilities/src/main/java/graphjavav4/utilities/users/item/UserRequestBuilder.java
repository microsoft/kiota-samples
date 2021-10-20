package graphjavav4.utilities.users.item;

import com.microsoft.kiota.RequestAdapter;
import graphjavav4.utilities.users.item.inferenceClassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.item.MailFolderRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.item.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.item.messages.MessagesRequestBuilder;
import java.util.HashMap;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}  */
public class UserRequestBuilder {
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        return new InferenceClassificationRequestBuilder(urlTemplateParameters, requestAdapter);
    }
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        return new MailFoldersRequestBuilder(urlTemplateParameters, requestAdapter);
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(urlTemplateParameters, requestAdapter);
    }
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /** Url template parameters for the request  */
    private final HashMap<String, String> urlTemplateParameters;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @param urlTemplateParameters Url template parameters for the request
     * @return a void
     */
    public UserRequestBuilder(@javax.annotation.Nonnull final HashMap<String, String> urlTemplateParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(requestAdapter);
        Objects.requireNonNull(urlTemplateParameters);
        this.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}";
        var urlTplParams = new HashMap<String, String>(urlTemplateParameters);
        this.urlTemplateParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public UserRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}";
        var urlTplParams = new HashMap<String, String>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.urlTemplateParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @return a mailFolderRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, String>(this.urlTemplateParameters);
        urlTplParams.put("mailFolder_id", id);
        return new MailFolderRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @return a messageRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, String>(this.urlTemplateParameters);
        urlTplParams.put("message_id", id);
        return new MessageRequestBuilder(urlTplParams, requestAdapter);
    }
}
