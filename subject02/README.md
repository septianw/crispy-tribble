# Subject #2

Goal: Dynamic symbol

In previous subject i've learn how to implement basic go plugin, there is a case when we don't know any symbol inside those library. This subject will cover those case.

In development environment of C/C++ there is a header file that need to load to make benefit of dynamic library. Go plugin is equal implementation of that case, we need to prepare our own header file to make .so file usefull for other application. In some case of environment, loading *.so file and recompile main app is not an option, and i curious about those case.