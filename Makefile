
# Install present with this command:
# $ go get golang.org/x/tools/cmd/present

.PHONY: present
present: nma.slide *.jpg *.png *.c *.go
	present -play nma.slide
