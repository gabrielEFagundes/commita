#include <stdio.h>
#include "../tools.h"

void exec_git_feat(int argc, char *argv[]){
    exec_command(":sparkles: feat: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_feat(argc, argv);
}
// rendundant, todo later