j .init
.init:
li sp, 0x10010000
j .main

.main:
li t0, 0xa
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
li t0, 1
bne t1, t2, .ac592080_5df6_4607_84e3_82c44dc892c8
li t0, 0
.ac592080_5df6_4607_84e3_82c44dc892c8:
neg t0, t0
sw t0, 0(sp)
addi sp, sp, 0x4
