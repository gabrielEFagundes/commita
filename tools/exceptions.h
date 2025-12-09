#include <stdio.h>

/// @brief Handles exceptions and errors that Git might return
///
/// Prints a user-friendly error message
///
/// @param err Parsed error message, used by the algorithm to find the exact error
void exception_handler(char *err);

/*
    I'm not exactly sure if I need to have methods just to print an error message,
    so I won't define them here just yet.
*/