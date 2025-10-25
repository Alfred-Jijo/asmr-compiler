const Op (
	ldr int = iota
	str
	add
	sub
	gcme
	gcml
	gcmg
	dmp
	udp
)
type struct Reg {
	r0 int
	r1 int
	r2 int
	r3 int
}

reg := Reg {
	.r0 = 0,
	.r1 = 0,
	.r2 = 0,
	.r3 = 0,
};


func ldr(destreg int, mem int) {
	switch destreg {
	case 0:
		reg.r0 = mem
	case 1:
		reg.r1 = mem
	case 2:
		reg.r2 = mem
	case 3:
		reg.r3 = mem
	}
}

func str(mem *int, srcreg int) {
	switch srcreg {
	case 0:
		mem = reg.r0
	case 1:
		mem = reg.r1
	case 2:
		mem = reg.r2
	case 3:
		mem = reg.r3
	}
}

func add(reg int) {
	switch reg {
	case 1:
		reg.r0 += reg.r1
	case 2:
		reg.r0 += reg.r2
	case 3:
		reg.r0 += reg.r3
	}
}

func sub(reg int) {
	switch reg {
	case 1:
		reg.r0 -= reg.r1
	case 2:
		reg.r0 -= reg.r2
	case 3:
		reg.r0 -= reg.r3
	}
}

func _cmp(bool res) int {
	if res {
		return 1;
	} else {
		return 0;
	}
}

func gcme(reg int) {
	switch reg {
	case 1:
		reg.r0 = _cmp(reg.r1 == rg.r0)
	case 2:
		reg.r0 = _cmp(reg.r2 == rg.r0)
	case 3:
		reg.r0 = _cmp(reg.r3 == rg.r0)
	}
}

func gcml(reg int) {
	switch reg {
	case 1:
		reg.r0 = _cmp(reg.r1 < rg.r0)
	case 2:
		reg.r0 = _cmp(reg.r2 < rg.r0)
	case 3:
		reg.r0 = _cmp(reg.r3 < rg.r0)
	}
}

func gcmg(reg int) {
	switch reg {
	case 1:
		reg.r0 = _cmp(reg.r1 > rg.r0)
	case 2:
		reg.r0 = _cmp(reg.r2 > rg.r0)
	case 3:
		reg.r0 = _cmp(reg.r3 > rg.r0)
	}
}

func dmp() {
	fmt.Printf("%d", reg.r0)
}

func udp() {
	fmt.Scanf("%u", reg.r0)
}
