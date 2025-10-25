# asmr-compiler

#### COLLABORATORS
@Alfred-Jijo

[@GarethSimmans](https://github.com/GarethSimmans)

@AlxFG

@samtowler

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

### LIST

to declare a list 
```asmr
list type elem... tsil
```
`tsil` is used to end a list decl

### FUNCTION
to declare a function
```asmr
decl name arg ... return 
begin
    ...
end
```

### primitives
the only types supported in asmr lang is `byte` and `byted8` for a single bit

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
use the `gcm` keyword 
