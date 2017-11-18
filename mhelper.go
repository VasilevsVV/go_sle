package main

// SwpLines swaps lines in matrix
func swpLines(m *[][]float64, l1, l2 int) {
	buf := (*m)[l1]
	(*m)[l1] = (*m)[l2]
	(*m)[l2] = buf
}
