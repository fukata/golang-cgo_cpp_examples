CXX= g++
CXXFLAGS= -Wall -O2 -fPIC -g
WARN= -pedantic -Wall
INCS= 
PREFIX=/usr

MYNAME= hello
MYLIB= $(MYNAME)
T= lib$(MYLIB).so
S= lib$(MYLIB).a
H= $(MYLIB).h
OBJS= $(MYLIB).o
LINKS=-L.

all:	$T $S 

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $<

so:	$T

a:	$S

$T:	$(OBJS)
	$(CXX) -o $@ -shared $(OBJS)

$S:	$(OBJS)
	ar rv $@ $(OBJS)

clean:
	rm -f $(OBJS) $T $S

install:
	cp -p $T $S $(PREFIX)/lib64/
	cp -p $H    $(PREFIX)/include/
