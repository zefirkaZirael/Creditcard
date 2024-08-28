

1. Validation Feature:
Luhn's Algorithm: Implemented to check if a credit card number is valid.
Minimum Length: The number must be at least 13 digits long.
Multiple Entries: Supports validation of multiple entries.
--stdin Flag: Allows passing numbers from stdin for validation.
Output: Prints "OK" to stdout if valid and exits with status 0. If invalid, prints "INCORRECT" to stderr and exits with status 1.

3. Generate Feature:
Asterisk Replacement: Replaces up to 4 asterisks (*) with digits. If more than 4 asterisks are present or they are not at the end of the card number, an error is returned.
Number Generation: Generates possible credit card numbers and prints them to stdout in ascending order.
Error Handling: Exits with status 1 if there is any error.
--pick Flag: Supports picking a single random entry from the generated numbers.
Example: ./creditcard generate --pick "440043018030****"

4. Information Feature:
Card Information: Outputs details about the card number, including validity, brand, and issuer, based on data from brands.txt and issuers.txt.
--stdin Flag: Allows passing numbers from stdin to retrieve card information.
Multiple Entries: Supports passing multiple entries to retrieve information.
Example: ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"

5. Issue Feature:
Random Number Generation: Generates a random valid credit card number for the specified brand and issuer.
Error Handling: Exits with status 1 if there is any error.
Example: ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"

Additional Requirements:
Code Formatting: The code adheres to gofumpt formatting standards.
Compilation: The project can be compiled successfully using the following command:
go build -o creditcard .
No Panics: The program handles errors gracefully and does not exit unexpectedly (e.g., no nil-pointer dereference, no index out of range errors).
Built-in Packages: Only built-in Go packages are used in the implementation.
