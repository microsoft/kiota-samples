package graphjavav4.utilities.users.item;

import com.microsoft.kiota.RequestAdapter;
import graphjavav4.utilities.users.item.inferenceclassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.item.MailFolderItemRequestBuilder;
import graphjavav4.utilities.users.item.mailfolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.item.messages.item.MessageItemRequestBuilder;
import graphjavav4.utilities.users.item.messages.MessagesRequestBuilder;
import java.util.HashMap;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}  */
public class UserItemRequestBuilder {
    /** The inferenceClassification property  */
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        return new InferenceClassificationRequestBuilder(pathParameters, requestAdapter);
    }
    /** The mailFolders property  */
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        return new MailFoldersRequestBuilder(pathParameters, requestAdapter);
    }
    /** The messages property  */
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(pathParameters, requestAdapter);
    }
    /** Path parameters for the request  */
    private final HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /**
     * Instantiates a new UserItemRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public UserItemRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/users/{user%2Did}";
        var urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new UserItemRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public UserItemRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/users/{user%2Did}";
        var urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @return a mailFolderItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderItemRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("mailFolder%2Did", id);
        return new MailFolderItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @return a messageItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageItemRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        var urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("message%2Did", id);
        return new MessageItemRequestBuilder(urlTplParams, requestAdapter);
    }
}
