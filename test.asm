j .init
.init:
li sp, 0x10010000
j .main

.main:
li t0, 0x64
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
slt t0, t1, t2
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
beq t0, zero, .1d072fa9_e8dc_4afc_b546_3b8c7aa9b460
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
.1d072fa9_e8dc_4afc_b546_3b8c7aa9b460: