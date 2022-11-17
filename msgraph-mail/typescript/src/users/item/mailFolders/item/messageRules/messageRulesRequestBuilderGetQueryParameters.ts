
<<<<<<< HEAD
export interface MessageRulesRequestBuilderGetQueryParameters {
=======
/** Get all the messageRule objects defined for the user's inbox. */
export class MessageRulesRequestBuilderGetQueryParameters {
>>>>>>> main
    /** Include count of items */
    count?: boolean | undefined;
    /** Filter items by property values */
    filter?: string | undefined;
    /** Order items by property values */
    orderby?: string[] | undefined;
    /** Select properties to be returned */
    select?: string[] | undefined;
    /** Skip the first n items */
    skip?: number | undefined;
    /** Show only the first n items */
    top?: number | undefined;
}
