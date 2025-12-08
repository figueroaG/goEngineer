package main

// STEP 1: Setup and Imports
// Import the following packages: fmt, context, sync, time.

// STEP 2: Define the Worker Function
// Define a function named processData.
// Arguments:
// 1. ctx of type context.Context
// 2. wg as a pointer to sync.WaitGroup
// 3. data as an empty interface (interface{})
func processData( /* arguments here */ ) {

	// STEP 3: Implement WaitGroup Signal
	// Strictly as the first line: defer the call to signal the wait group (wg.Done()).

	// STEP 4: Implement Type Assertion (The "Check")
	// Use the "comma-ok" idiom to check if 'data' is a string.
	// If it is a string, print: "Checking string length...".
	// Do not perform other logic here.

	// STEP 5: Implement the Context/Timeout Logic
	// Create a select statement.
	// Case 1: Check if ctx.Done().
	//    Inside this case, print "Context cancelled for data: " followed by the data value.
	//    Return immediately.
	// Case 2: Simulate work using time.After(500 * time.Millisecond).
	//    This case will contain the logic for Step 6.

		// STEP 6: Implement Type Switch (The "Processing")
		// Inside the time.After case:
		// Create a type switch on 'data'.
		// Case string: Print "Processed String: " followed by the string value.
		// Case int: Print "Processed Int: " followed by the integer value.
		// Default: Print "Unknown type encountered".

}

// STEP 7: The Main Routine
func main() {
	// Initialize a sync.WaitGroup variable.

	// Create a derived context using context.WithTimeout based on context.Background().
	// Set the timeout duration to 200 milliseconds.
	// Defer the cancellation function to avoid leaks.

	// STEP 8: Dispatch Goroutines
	// Increment the WaitGroup counter by 3.

	// Launch 3 separate goroutines calling processData.
	// 1. Pass the string "Alpha".
	// 2. Pass the integer 42.
	// 3. Pass a boolean true.

	// STEP 9: Wait for Completion
	// Block execution until all goroutines have finished using the WaitGroup.
	// Print "Program exit".
}