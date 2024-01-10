package access

// Casbin Usage Guide
//
// There are two files to create for the casbin configuration:
// * a configuration file
// * a policy file
//
// ## THe configuration file
//
// Uses the PERM metamodeal:
// Policy, Effect, Request, Matchers.
//
// The configuration file defines request and policy.
//
// In the HTTP case,
//
// • subject -> user role
// • object -> path the user wants to access
// • action -> request method
//
// The `matchers` define how the policy parts are matched,
// either directly like the subject or with a helper method
// like `keyMatch`, which can also match wildcards.
//
// ## The policy file
//
// A simple csv file describing which role can access which path.
// In the `policy.csv` file, we define that
// * the `admin` role can access everything
// * the `member` role can access everything after `/member` as well as `logout`
// * all unauthenticated users can use the `login`.
