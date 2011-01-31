include $(GOROOT)/src/Make.inc

TARG=com.abneptis.oss/cryptools
GOFILES=\
	common_interfaces.go\
	signable_interfaces.go\
	signer_interfaces.go\
	verifier_interfaces.go\

include $(GOROOT)/src/Make.pkg

