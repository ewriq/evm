.section .text.startup
.global _start
.type _start, @function

_start:

    call0 app_main

.Lhang:
    j .Lhang

.size _start, .-_start
