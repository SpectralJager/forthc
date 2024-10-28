j .init
.init:
li sp, 0x10010000
j .main

.main:
li t0, -0x1
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, -0x1
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
li t0, 0
beqz t1, .17f4e84b_67d7_4aaf_a402_524bc4a95281
beqz t2, .17f4e84b_67d7_4aaf_a402_524bc4a95281
li t0, 1
.17f4e84b_67d7_4aaf_a402_524bc4a95281:
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4