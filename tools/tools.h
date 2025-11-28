#include <stdio.h>
#include <windows.h>
#include <string.h>

#ifndef EXEC_COMMAND_UTILITARY
#define EXEC_COMMAND_UTILITARY

/// @brief Auxiliar command, executes the commands and pushes to your branch
/// 
/// Usually, it stages all your files, commits them with your message and then push it to your repo's branch
/// @param prefix The prefix (feat, fix, etc)
/// @param argcount Counts the amount of arguments, usually 3 (git <prefix> <message> <branch> <repo_url>)
/// @param argvec Vector with your command and message
int exec_command(char *prefix, int argcount, char *argvec[]);
#endif

/// @brief Executes a series of git commands, uses the user message to build the commit message, also uses the repo URL
/// @param argcount Counts the amount of arguments, usually 3 (git <prefix> <message> <repo_url>)
/// @param argvect Vector with your command and message
void exec_git_feat(int argcount, char *argvect[]);
void exec_git_fix(int argcount, char *argvect[]);
void exec_git_chore(int argcount, char *argvect[]);
void exec_git_docs(int argcount, char *argvect[]);
void exec_git_refactor(int argcount, char *argvect[]);
void exec_git_remove(int argcount, char *argvect[]);

/// @brief Shows a message with all commands, in case of help needed
void exec_help();