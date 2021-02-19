# assertion
This is a library that is used for testing in golang. It can be used to create

All assertions are performed by an Asserter which wraps a *testing.T and calls
failure functions on it when assertions made by the Asserter fail.
