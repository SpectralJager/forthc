j .init
.init:
li sp, 0x10010000
j .main

.main:
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
addi sp, sp, 0x4
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x64
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
li t0, 1
beq t1, t2, .800abbdd_9391_4e06_a97d_b5f2f149a58f
li t0, 0
.800abbdd_9391_4e06_a97d_b5f2f149a58f:
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
beq t0, zero, .820285f7_36b1_4215_be95_6b77d987eb43_else
li t0, 0x64
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
add t0, t1, t2
sw t0, 0(sp)
addi sp, sp, 0x4
j .820285f7_36b1_4215_be95_6b77d987eb43
.820285f7_36b1_4215_be95_6b77d987eb43_else:
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
add t0, t1, t2
sw t0, 0(sp)
addi sp, sp, 0x4
.820285f7_36b1_4215_be95_6b77d987eb43: