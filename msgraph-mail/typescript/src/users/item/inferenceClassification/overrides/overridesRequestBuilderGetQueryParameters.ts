
<<<<<<< HEAD
export interface OverridesRequestBuilderGetQueryParameters {
=======
/** Get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides. */
export class OverridesRequestBuilderGetQueryParameters {
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
