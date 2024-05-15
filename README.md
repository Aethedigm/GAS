# GAS
 Go Autocomplete Service

### Purpose
GAS provides a resource for compounding data (typically URLs) for your application behind keywords. This works well for 
things such as autocomplete for a search feature, where you might allow a user to begin typing in search, and retrieve 
their results with followable links to their destination.

### Usage
The Go Autocomplete System allows multiple "GAS", think of each as a category of results you might want to separate from
each other. You can easily maintain separate information for users, documents, posts, etc. and only query for the types
that you might want.

#### GAS exposes 4 routes
- Add Result
- - [POST] ( "/" ) `{ "gas": "XXX", "key": "YYY", "value": "ZZZ" }` Adds a value to a given key inside of the provided
    GAS, assuming the GAS exists. Creates the key if it didn't exist. Will not create a GAS.
- Get Result
- - [GET] ( "/?gas=XXX&key=YYY" ) Retrieves all matching words to the prefix provided, along with their values.
- Add Gas
- - [POST] ( "/gas" ) `{ "gas": "XXX" }` Allows the user to create a new GAS category for use
- Get Gas 
- - [GET] ( "/gas" ) Retrieves an array of all GAS categories

#### TODO
1. Replace all maps with rune arrays
2. Configuration Settings
   1. Allow HTTPS
   2. Allow restricting characters
   3. Allow setting result limits
3. Save state
4. Load save state