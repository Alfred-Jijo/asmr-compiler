# asmr-compiler

#### COLLABORATORS
[@Alfred-Jijo](https://github.com/Alfred-Jijo)

[@GarethSimmans](https://github.com/GarethSimmans)

[@AlxFG](https://github.com/AlxFG)

[@samtowler](https://github.com/samtowler)

---
`NOTE: CURRENT IMPLEMTATION IS JUST AN INTERPRETER`
## Language

### Types
```asmr
byte
```
#### Assignment
```asmr
ldv byte [variable] [var|number]
```
### Control Flow
```asmr
gcm[e|l|g]
    [operation]
alt
    [operation]
end
```
### Operators
```asmr
plus dest [var|num] [var|num]
sulp dest [var|num] [var|num]
dmp stream [var|num]
udp stream [var|num]
```
> [!NOTE]  
> `stream` for the `dump` and `undump` operations are like `C`:
>     - `stdin` and `stdout` or `file` 
