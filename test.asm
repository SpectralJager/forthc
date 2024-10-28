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
.2f82db7f_c3fc_4a77_9dad_dd6bf32dc44b:
sw t6, 0(sp)
addi sp, sp, 0x4
addi t6, t6, 0x1
bne t5, t6, .2f82db7f_c3fc_4a77_9dad_dd6bf32dc44b
