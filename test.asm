j .init
.init:
li sp, 0x10010000
li s0, 0x10040000
j .main

.main:
li t0, 0x0
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
.186a3390_9dba_4b8e_9388_bef663dfce56:
addi t0, s0, 0x0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
lw t0, 0(t0)
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x1
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
add t0, t1, t2
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
addi t0, s0, 0x0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
lw t0, 0(t0)
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
li t0, 1
beq t1, t2, .a4ce7541_6d07_4b9e_b796_944576ec1450
li t0, 0
.a4ce7541_6d07_4b9e_b796_944576ec1450:
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
beqz t0, .186a3390_9dba_4b8e_9388_bef663dfce56