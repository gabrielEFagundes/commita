#include <stdio.h>
#include "../tools.h"

void exec_git_docs(int argc, char *argv[]){
    exec_command(":memo: doc: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_docs(argc, argv);
}