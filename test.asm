init:
li sp, 0x10010000
j main

main:
li t0, -0x5
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x4
sub sp, sp, t0
lw t2, 0(sp)
sub sp, sp, t0
lw t1, 0(sp)
add t0, t1, t2
sw t0, 0(sp)
addi sp, sp, 0x4