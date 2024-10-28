j .init
.init:
li sp, 0x10010000
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
.91c2ad3f_b367_4cb6_a130_74ee38e9c4b9:
li t0, 0xa
sw t0, 0(sp)
addi sp, sp, 0x4
sw t6, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t4, 0(sp)
addi sp, sp, -0x4
lw t3, 0(sp)
.91fc4db8_cb04_4e1b_bd59_084de77a2c39:
sw t6, 0(sp)
addi sp, sp, 0x4
sw t4, 0(sp)
addi sp, sp, 0x4
addi sp, sp, -0x4
lw t2, 0(sp)
addi sp, sp, -0x4
lw t1, 0(sp)
mul t0, t1, t2
sw t0, 0(sp)
addi sp, sp, 0x4
addi t4, t4, 0x1
bne t3, t4, .91fc4db8_cb04_4e1b_bd59_084de77a2c39
addi t6, t6, 0x1
bne t5, t6, .91c2ad3f_b367_4cb6_a130_74ee38e9c4b9