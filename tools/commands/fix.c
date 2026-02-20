#include <stdio.h>
#include "../headers/tools.h"

void exec_git_fix(int argc, char *argv[]){
    exec_command(":bug: fix: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_fix(argc, argv);
}