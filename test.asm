j .init
.init:
li sp, 0x10010000
li s0, 0x10040000
j .main

.main:
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x14
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x1e
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x10010000
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x10010020
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x3
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
addi sp, sp, -0x4
lw t2, 0(sp)
.928bc1ce_4da2_492e_849f_72264f824a45:
lw t3, 0(t2)
sw t3, 0(t1)
addi t0, t0, -0x1
addi t1, t1, 0x4
addi t2, t2, 0x4
bnez t0, .928bc1ce_4da2_492e_849f_72264f824a45