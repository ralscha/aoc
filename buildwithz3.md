```
brew install z3
```

```
export CGO_CFLAGS=-I/opt/homebrew/include
export CGO_LDFLAGS=-L/opt/homebrew/lib
export LD_LIBRARY_PATH=/opt/homebrew/lib
```



Windows
Download llvm-mingw and z3
https://github.com/Z3Prover/z3
https://github.com/mstorsjo/llvm-mingw
```
set CGO_ENABLED=1
set CC="<path_to>\llvm-mingw-20231128-msvcrt-x86_64\bin\gcc.exe"
set LD_LIBRARY_PATH=<path_to>\z3-4.12.2-x64-win\bin\
set CGO_CFLAGS=-I<path_to>\z3-4.12.2-x64-win\include\
set CGO_LDFLAGS=-L<path_to>\z3-4.12.2-x64-win\bin\
```