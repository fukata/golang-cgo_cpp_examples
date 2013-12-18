#include <sstream>
#include <string>
#include <cstring>
#include "say.h"

using namespace std;

char *say(char *name) {
    ostringstream os;
    os << "Hi! " << name;
    return const_cast<char*>(os.str().c_str());
}

char *say_cstring(char *name) {
    ostringstream os;
    os << "Hi! " << name;
    char *buf = new char[128];
    std::strcpy(buf, os.str().c_str());
    return buf;
}
