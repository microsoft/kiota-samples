
export interface MessageItemRequestBuilderGetQueryParameters {
    /**
     * Expand related entities
     */
    expand?: string[] | undefined;
    /**
     * Include Hidden Messages
     */
    includeHiddenMessages?: string | undefined;
    /**
     * Select properties to be returned
     */
    select?: string[] | undefined;
}
