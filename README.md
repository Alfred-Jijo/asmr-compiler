# asmr-compiler

#### COLLABORATORS
[@Alfred-Jijo](https://github.com/Alfred-Jijo)

[@GarethSimmans](https://github.com/GarethSimmans)

[@AlxFG](https://github.com/AlxFG)

[@samtowler](https://github.com/samtowler)

---
`NOTE: CURRENT IMPLEMTATION IS JUST AN INTERPRETER`
## Language

### IO
for output to the console
```asmr
call dmp literal
```
for reading input from the console
```asmr
call udp var
```

### primitives
the only types supported in asmr lang is `byte` 

### ASSIGNMENT
to assign variables
```asmr
ldv type name literal
```

### OPERATORS
```asmr
... plus ...
... sulp ...
```
asmr lang only supports addition with `plus`
and subtraction with `sulp`.

### COMPARISONS
use the `gcm` (gocompare) keyword with relevant letters
```asmr
    gmce ... ...
    gmcg ... ... 
    gmcl ... ...
    gmcn ... ...
```
Letters are self-explanatory like the rest of this language.
