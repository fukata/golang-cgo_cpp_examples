#include <sstream>
#include "hello.h"

char *hello() {
    std::ostringstream os;
    os << "Hello" << " CPP " << "World";
    return const_cast<char*>(os.str().c_str());
}
