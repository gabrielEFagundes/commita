#include <stdio.h>
#include <windows.h>
#include <string.h>

char buffer[1000];

/// @brief Auxiliar command, executes the message and pushes to your branch
/// @param branch Your branch
/// @param prefix The prefix (feat, fix, etc)
/// @param argcount Counts the amount of arguments, usually 2 (git <prefix> <message>)
/// @param argvec Vector with your command and message
int exec_command(const char *branch, char *prefix, int argcount, char *argvec[]);

/// @brief Execute the command that the user inputs, this also executes "git -A"
/// @param argcount Counts the amount of arguments, usually 2 (the git command and message)
/// @param argvect Vector with your command and message
void exec_git_feat(int argcount, char *argvect[]);
void exec_git_fix(int argcount, char *argvect[]);
void exec_git_chore(int argcount, char *argvect[]);
void exec_git_docs(int argcount, char *argvect[]);
void exec_git_refactor(int argcount, char *argvect[]);
void exec_git_remove(int argcount, char *argvect[]);

/// @brief Shows a message with all commands, in case of help needed
void exec_help();