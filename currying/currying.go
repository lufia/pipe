// Package currying provides utilities for currying.
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

// First2 transforms f into callable as f(v1)(v2).
func First2[T1, T2, R1 any](f func(T1, T2) R1) func(T2) func(T1) {
	return func(v2 T2) func(T1) {
		return func(v1 T1) {
			f(v1, v2)
		}
	}
}

// First2R1 is like [First2] but it returning R1.
func First2R1[T1, T2, R1 any](f func(T1, T2) R1) func(T2) func(T1) R1 {
	return func(v2 T2) func(T1) R1 {
		return func(v1 T1) R1 {
			return f(v1, v2)
		}
	}
}

// First2R2 is like [First2] but it returning R1 and R2.
func First2R2[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2)) func(T2) func(T1) (R1, R2) {
	return func(v2 T2) func(T1) (R1, R2) {
		return func(v1 T1) (R1, R2) {
			return f(v1, v2)
		}
	}
}

// First2R3 is like [First2] but it returning R1, R2 and R3.
func First2R3[T1, T2, R1, R2, R3 any](f func(T1, T2) (R1, R2, R3)) func(T2) func(T1) (R1, R2, R3) {
	return func(v2 T2) func(T1) (R1, R2, R3) {
		return func(v1 T1) (R1, R2, R3) {
			return f(v1, v2)
		}
	}
}

// First3 transforms f into callable as f(v2, v3)(v1).
func First3[T1, T2, T3 any](f func(T1, T2, T3)) func(T2, T3) func(T1) {
	return func(v2 T2, v3 T3) func(T1) {
		return func(v1 T1) {
			f(v1, v2, v3)
		}
	}
}

// First3R1 is like [First3] but it returning R1.
func First3R1[T1, T2, T3, R1 any](f func(T1, T2, T3) R1) func(T2, T3) func(T1) R1 {
	return func(v2 T2, v3 T3) func(T1) R1 {
		return func(v1 T1) R1 {
			return f(v1, v2, v3)
		}
	}
}

// First3R2 is like [First3] but it returning R1 and R2.
func First3R2[T1, T2, T3, R1, R2 any](f func(T1, T2, T3) (R1, R2)) func(T2, T3) func(T1) (R1, R2) {
	return func(v2 T2, v3 T3) func(T1) (R1, R2) {
		return func(v1 T1) (R1, R2) {
			return f(v1, v2, v3)
		}
	}
}

// First3R3 is like [First3] but it returning R1, R2 and R3.
func First3R3[T1, T2, T3, R1, R2, R3 any](f func(T1, T2, T3) (R1, R2, R3)) func(T2, T3) func(T1) (R1, R2, R3) {
	return func(v2 T2, v3 T3) func(T1) (R1, R2, R3) {
		return func(v1 T1) (R1, R2, R3) {
			return f(v1, v2, v3)
		}
	}
}
