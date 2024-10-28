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
beq t1, t2, .e96cdd71_f7d4_4dee_bcd8_84ad06ce05b5
li t0, 0
.e96cdd71_f7d4_4dee_bcd8_84ad06ce05b5:
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t0, 0(sp)
beq t0, zero, .de1e0ff9_db59_4a08_8ed4_4a983122d776_else
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
j .de1e0ff9_db59_4a08_8ed4_4a983122d776
.de1e0ff9_db59_4a08_8ed4_4a983122d776_else:
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
.de1e0ff9_db59_4a08_8ed4_4a983122d776:
