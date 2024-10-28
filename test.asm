j .init
.init:
li sp, 0x10010000
li s0, 0x10040000
j .main

.main:
li t0, 0x7b
sw t0, 0(sp)
addi sp, sp, 0x4
addi t0, s0, 0x0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
sw t1, 0(t0)
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
addi t0, s0, 0x4
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
sw t1, 0(t0)
addi t0, s0, 0x0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
lw t0, 0(t0)
sw t0, 0(sp)
addi sp, sp, 0x4
addi t0, s0, 0x4
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
lw t0, 0(t0)
sw t0, 0(sp)
addi sp, sp, 0x4