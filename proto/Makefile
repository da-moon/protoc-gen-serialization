# OS specific part
# -----------------
ifeq ($(OS),Windows_NT)
    CLEAR = cls
    LS = dir
    TOUCH =>> 
    RM = del /F /Q
    CPF = copy /y
    RMDIR = -RMDIR /S /Q
    MKDIR = -mkdir
    ERRIGNORE = 2>NUL || (exit 0)
    GO_PATH = $(subst \,/,${GOPATH})
    SEP=\\
else
    CLEAR = clear
    GO_PATH = ${GOPATH}
    LS = ls
    TOUCH = touch
    CPF = cp -f
    RM = rm -rf 
    RMDIR = rm -rf 
    MKDIR = mkdir -p
    ERRIGNORE = 2>/dev/null
    SEP=/
endif

DIRS = $(notdir $(patsubst %/,%,$(dir $(wildcard */.))))
PSEP = $(strip $(SEP))
.PHONY: clean build

clean:	
ifneq ("$(wildcard ./*.go)","")
	- $(RM) .$(PSEP)*.go  
endif
# Definitions for nullstring and space
# -------------------------------------
nullstring :=
space := $(nullstring) #End

