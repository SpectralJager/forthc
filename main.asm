j .init
.init:
li sp, 0x10010000
li s0, 0x10040000
j .main

.main:
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
li t0, 0x0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t6, 0(sp)
addi sp, sp, -0x4
lw t5, 0(sp)
.bb6ba40a_9e44_4f86_bc91_feed3605e551:
sw t6, 0(sp)
addi sp, sp, 0x4
addi t6, t6, 0x1
bne t5, t6, .bb6ba40a_9e44_4f86_bc91_feed3605e551
