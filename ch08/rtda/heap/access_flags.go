package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method 0000 0000 0000 0001
	ACC_PRIVATE      = 0x0002 //       field method 0000 0000 0000 0010
	ACC_PROTECTED    = 0x0004 //       field method 0000 0000 0000 0100
	ACC_STATIC       = 0x0008 //       field method 0000 0000 0000 1000
	ACC_FINAL        = 0x0010 // class field method 0000 0000 0001 0000
	ACC_SUPER        = 0x0020 // class              0000 0000 0010 0000
	ACC_SYNCHRONIZED = 0x0020 //             method 0000 0000 0010 0000
	ACC_VOLATILE     = 0x0040 //       field        0000 0000 0100 0000
	ACC_BRIDGE       = 0x0040 //             method 0000 0000 0100 0000
	ACC_TRANSIENT    = 0x0080 //       field        0000 0000 1000 0000
	ACC_VARARGS      = 0x0080 //             method 0000 0000 1000 0000
	ACC_NATIVE       = 0x0100 //             method 0000 0001 0000 0000
	ACC_INTERFACE    = 0x0200 // class              0000 0010 0000 0000
	ACC_ABSTRACT     = 0x0400 // class       method 0000 0100 0000 0000
	ACC_STRICT       = 0x0800 //             method 0000 1000 0000 0000
	ACC_SYNTHETIC    = 0x1000 // class field method 0001 0000 0000 0000
	ACC_ANNOTATION   = 0x2000 // class              0010 0000 0000 0000
	ACC_ENUM         = 0x4000 // class field        0100 0000 0000 0000
)
