// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 40: // ['(','(']
			return 3
		case r == 41: // [')',')']
			return 4
		case r == 42: // ['*','*']
			return 5
		case r == 43: // ['+','+']
			return 6
		case r == 45: // ['-','-']
			return 7
		case r == 46: // ['.','.']
			return 8
		case r == 47: // ['/','/']
			return 9
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 60: // ['<','<']
			return 11
		case r == 63: // ['?','?']
			return 12
		case r == 65: // ['A','A']
			return 13
		case r == 66: // ['B','B']
			return 14
		case r == 67: // ['C','C']
			return 15
		case 68 <= r && r <= 69: // ['D','E']
			return 16
		case r == 70: // ['F','F']
			return 17
		case 71 <= r && r <= 72: // ['G','H']
			return 16
		case r == 73: // ['I','I']
			return 18
		case 74 <= r && r <= 75: // ['J','K']
			return 16
		case r == 76: // ['L','L']
			return 19
		case r == 77: // ['M','M']
			return 16
		case r == 78: // ['N','N']
			return 20
		case 79 <= r && r <= 82: // ['O','R']
			return 16
		case r == 83: // ['S','S']
			return 21
		case r == 84: // ['T','T']
			return 22
		case r == 85: // ['U','U']
			return 23
		case r == 86: // ['V','V']
			return 24
		case r == 87: // ['W','W']
			return 25
		case 88 <= r && r <= 90: // ['X','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case r == 97: // ['a','a']
			return 26
		case 98 <= r && r <= 122: // ['b','z']
			return 27
		case r == 123: // ['{','{']
			return 28
		case r == 124: // ['|','|']
			return 29
		case r == 125: // ['}','}']
			return 30
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		default:
			return 2
		}
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 62: // ['>','>']
			return 33
		default:
			return 11
		}
	},
	// S12
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 37
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 69: // ['A','E']
			return 16
		case r == 70: // ['F','F']
			return 38
		case 71 <= r && r <= 83: // ['G','S']
			return 16
		case r == 84: // ['T','T']
			return 39
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 40
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 41
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 42
		case 80 <= r && r <= 81: // ['P','Q']
			return 16
		case r == 82: // ['R','R']
			return 43
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 77: // ['A','M']
			return 16
		case r == 78: // ['N','N']
			return 44
		case 79 <= r && r <= 90: // ['O','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 72: // ['A','H']
			return 16
		case r == 73: // ['I','I']
			return 45
		case 74 <= r && r <= 90: // ['J','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case r == 65: // ['A','A']
			return 46
		case 66 <= r && r <= 90: // ['B','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 47
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 48
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 77: // ['A','M']
			return 16
		case r == 78: // ['N','N']
			return 49
		case 79 <= r && r <= 90: // ['O','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 50
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 71: // ['A','G']
			return 16
		case r == 72: // ['H','H']
			return 51
		case 73 <= r && r <= 90: // ['I','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 52
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 54
		case r == 95: // ['_','_']
			return 52
		case 97 <= r && r <= 122: // ['a','z']
			return 55
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 37
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 37
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 37
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 37
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 56
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 69: // ['A','E']
			return 16
		case r == 70: // ['F','F']
			return 57
		case 71 <= r && r <= 90: // ['G','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 84: // ['A','T']
			return 16
		case r == 85: // ['U','U']
			return 58
		case 86 <= r && r <= 90: // ['V','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 59
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 60
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 82: // ['A','R']
			return 16
		case r == 83: // ['S','S']
			return 61
		case 84 <= r && r <= 90: // ['T','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 76: // ['A','L']
			return 16
		case r == 77: // ['M','M']
			return 62
		case 78 <= r && r <= 82: // ['N','R']
			return 16
		case r == 83: // ['S','S']
			return 63
		case 84 <= r && r <= 90: // ['T','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 76: // ['A','L']
			return 16
		case r == 77: // ['M','M']
			return 64
		case 78 <= r && r <= 90: // ['N','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 75: // ['A','K']
			return 16
		case r == 76: // ['L','L']
			return 65
		case 77 <= r && r <= 90: // ['M','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 72: // ['A','H']
			return 16
		case r == 73: // ['I','I']
			return 66
		case 74 <= r && r <= 90: // ['J','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 67
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 68
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 52
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 54
		case r == 95: // ['_','_']
			return 52
		case 97 <= r && r <= 122: // ['a','z']
			return 55
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 52
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 54
		case r == 95: // ['_','_']
			return 52
		case 97 <= r && r <= 122: // ['a','z']
			return 55
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 52
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 54
		case r == 95: // ['_','_']
			return 52
		case 97 <= r && r <= 122: // ['a','z']
			return 55
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 52
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 54
		case r == 95: // ['_','_']
			return 52
		case 97 <= r && r <= 122: // ['a','z']
			return 55
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 69
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 70
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 77: // ['A','M']
			return 16
		case r == 78: // ['N','N']
			return 71
		case 79 <= r && r <= 90: // ['O','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 76: // ['A','L']
			return 16
		case r == 77: // ['M','M']
			return 72
		case 78 <= r && r <= 90: // ['N','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 73
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 72: // ['A','H']
			return 16
		case r == 73: // ['I','I']
			return 74
		case 74 <= r && r <= 90: // ['J','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 75
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 76
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 77
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 78
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 82: // ['A','R']
			return 16
		case r == 83: // ['S','S']
			return 79
		case 84 <= r && r <= 90: // ['T','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 80
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 81
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 82
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 83
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 81: // ['A','Q']
			return 16
		case r == 82: // ['R','R']
			return 84
		case 83 <= r && r <= 90: // ['S','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 85
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 82: // ['A','R']
			return 16
		case r == 83: // ['S','S']
			return 86
		case 84 <= r && r <= 90: // ['T','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 66: // ['A','B']
			return 16
		case r == 67: // ['C','C']
			return 87
		case 68 <= r && r <= 90: // ['D','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 77: // ['A','M']
			return 16
		case r == 78: // ['N','N']
			return 88
		case 79 <= r && r <= 90: // ['O','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 72: // ['A','H']
			return 16
		case r == 73: // ['I','I']
			return 89
		case 74 <= r && r <= 90: // ['J','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 90
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 68: // ['A','D']
			return 16
		case r == 69: // ['E','E']
			return 91
		case 70 <= r && r <= 90: // ['F','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 92
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 83: // ['A','S']
			return 16
		case r == 84: // ['T','T']
			return 93
		case 85 <= r && r <= 90: // ['U','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S89
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 78: // ['A','N']
			return 16
		case r == 79: // ['O','O']
			return 94
		case 80 <= r && r <= 90: // ['P','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S90
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S91
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S92
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S93
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S94
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 77: // ['A','M']
			return 16
		case r == 78: // ['N','N']
			return 95
		case 79 <= r && r <= 90: // ['O','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S95
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 82: // ['A','R']
			return 16
		case r == 83: // ['S','S']
			return 96
		case 84 <= r && r <= 90: // ['T','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S96
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 10
		case r == 58: // [':',':']
			return 32
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
}
