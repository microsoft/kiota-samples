// ignore_for_file: type=lint
import 'package:microsoft_kiota_abstractions/microsoft_kiota_abstractions.dart';

/// auto generated
/// Get posts
class PostsRequestBuilderGetQueryParameters implements AbstractQueryParameters {
    ///  Filter results by title
    /// @QueryParameter('title')
    String? title;
    ///  Filter results by user ID
    /// @QueryParameter('userId')
    int? userId;
    /// Extracts the query parameters into a map for the URI template parsing.
    @override
    Map<String, dynamic> toMap() {
        return {
            'title' : title,
            'userId' : userId,
        };
    }
}
