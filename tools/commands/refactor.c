#include <stdio.h>
#include "../tools.h"

void exec_git_refactor(int argc, char *argv[]){
    exec_command(":recycle: refactor: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_refactor(argc, argv);
}