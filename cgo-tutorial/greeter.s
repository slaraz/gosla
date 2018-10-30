	.file	"greeter.c"
	.section .rdata,"dr"
	.align 8
.LC0:
	.ascii "Greetings, %s from %d! We come in peace :)\0"
	.text
	.globl	greet
	.def	greet;	.scl	2;	.type	32;	.endef
	.seh_proc	greet
greet:
	pushq	%rbp
	.seh_pushreg	%rbp
	movq	%rsp, %rbp
	.seh_setframe	%rbp, 0
	subq	$48, %rsp
	.seh_stackalloc	48
	.seh_endprologue
	movq	%rcx, 16(%rbp)
	movl	%edx, 24(%rbp)
	movq	%r8, 32(%rbp)
	movl	24(%rbp), %edx
	movq	32(%rbp), %rax
	movl	%edx, %r9d
	movq	16(%rbp), %r8
	leaq	.LC0(%rip), %rdx
	movq	%rax, %rcx
	call	sprintf
	movl	%eax, -4(%rbp)
	movl	-4(%rbp), %eax
	addq	$48, %rsp
	popq	%rbp
	ret
	.seh_endproc
	.ident	"GCC: (x86_64-posix-seh-rev1, Built by MinGW-W64 project) 7.2.0"
	.def	sprintf;	.scl	2;	.type	32;	.endef
