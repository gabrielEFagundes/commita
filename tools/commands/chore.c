#include <stdio.h>
#include "../tools.h"

void exec_git_chore(int argc, char *argv[]){
    exec_command(":art: chore: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_chore(argc, argv);
}