package currying

// Last2 transforms f into callable as f(v1)(v2).
func Last2[T1, T2, R1 any](f func(T1, T2) R1) func(T1) func(T2) {
	return func(v1 T1) func(T2) {
		return func(v2 T2) {
			f(v1, v2)
		}
	}
}

// Last2R1 is like [Last2] but it returning R1.
func Last2R1[T1, T2, R1 any](f func(T1, T2) R1) func(T1) func(T2) R1 {
	return func(v1 T1) func(T2) R1 {
		return func(v2 T2) R1 {
			return f(v1, v2)
		}
	}
}

// Last2R2 is like [Last2] but it returning R1 and R2.
func Last2R2[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2)) func(T1) func(T2) (R1, R2) {
	return func(v1 T1) func(T2) (R1, R2) {
		return func(v2 T2) (R1, R2) {
			return f(v1, v2)
		}
	}
}

// Last2R3 is like [Last2] but it returning R1, R2 and R3.
func Last2R3[T1, T2, R1, R2, R3 any](f func(T1, T2) (R1, R2, R3)) func(T1) func(T2) (R1, R2, R3) {
	return func(v1 T1) func(T2) (R1, R2, R3) {
		return func(v2 T2) (R1, R2, R3) {
			return f(v1, v2)
		}
	}
}

// Last3 transforms f into callable as f(v1, v2)(v3).
func Last3[T1, T2, T3 any](f func(T1, T2, T3)) func(T1, T2) func(T3) {
	return func(v1 T1, v2 T2) func(T3) {
		return func(v3 T3) {
			f(v1, v2, v3)
		}
	}
}

// Last3R1 is like [Last3] but it returning R1.
func Last3R1[T1, T2, T3, R1 any](f func(T1, T2, T3) R1) func(T1, T2) func(T3) R1 {
	return func(v1 T1, v2 T2) func(T3) R1 {
		return func(v3 T3) R1 {
			return f(v1, v2, v3)
		}
	}
}

// Last3R2 is like [Last3] but it returning R1 and R2.
func Last3R2[T1, T2, T3, R1, R2 any](f func(T1, T2, T3) (R1, R2)) func(T1, T2) func(T3) (R1, R2) {
	return func(v1 T1, v2 T2) func(T3) (R1, R2) {
		return func(v3 T3) (R1, R2) {
			return f(v1, v2, v3)
		}
	}
}

// Last3R3 is like [Last3] but it returning R1, R2 and R3.
func Last3R3[T1, T2, T3, R1, R2, R3 any](f func(T1, T2, T3) (R1, R2, R3)) func(T1, T2) func(T3) (R1, R2, R3) {
	return func(v1 T1, v2 T2) func(T3) (R1, R2, R3) {
		return func(v3 T3) (R1, R2, R3) {
			return f(v1, v2, v3)
		}
	}
}

func All3[T1, T2, T3 any](f func(T1, T2, T3)) func(T1) func(T2) func(T3) {
	return Last2R1(Last3(f))
}

func All3R1[T1, T2, T3, Out any](f func(T1, T2, T3) Out) func(T1) func(T2) func(T3) Out {
	return Last2R1(Last3R1(f))
}

func All3R2[T1, T2, T3, Out1, Out2 any](f func(T1, T2, T3) (Out1, Out2)) func(T1) func(T2) func(T3) (Out1, Out2) {
	return Last2R1(Last3R2(f))
}

func All3R3[T1, T2, T3, Out1, Out2, Out3 any](f func(T1, T2, T3) (Out1, Out2, Out3)) func(T1) func(T2) func(T3) (Out1, Out2, Out3) {
	return Last2R1(Last3R3(f))
}
