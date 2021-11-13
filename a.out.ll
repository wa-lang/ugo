define i32 @main() {
	%t1 = add i32 0, 1; %t1 = 1
	%t3 = add i32 0, 2; %t3 = 2
	%t5 = add i32 0, 3; %t5 = 3
	%t6 = add i32 0, 4; %t6 = 4
	%t4 = add i32 %t5, %t6
	%t2 = mul i32 %t3, %t4
	%t0 = add i32 %t1, %t2
	ret i32 %t0
}
