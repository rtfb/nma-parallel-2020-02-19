// CPU1                             CPU2

r5 = flag
while (r5 == false);                x = 42
// memory barrier (won't help)      // memory barrier (won't help)
print x                             flag = true

