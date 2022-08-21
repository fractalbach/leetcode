
def generate(n):
	# output array, answers are stored here
	out = []

	# contains index for each of the left-parentheses
	loc = [i for i in range(n)]
	
	# current paren we are working on (starts at the right-most)
	cur = n-1
	
	backtrack(cur, loc, out) # begin backtrack algorithm
	
	return out
	
	
# backtrack
# cur = index of the character we currently are moving
# loc = indices tracking the position of parentheses
# n   = length of each string
# out = array to append solutions to
def backtrack(cur, loc, out):

	output(loc, out)
	
	while loc[cur] < 2*cur\
		  and cur != 0\
		  and (cur >= len(loc)-1 or loc[cur]+1 < loc[cur+1]):
		loc[cur] += 1
		backtrack(cur-1, list.copy(loc), out)



def output(loc, out):
	s = [')'] * (2*len(loc))
	for v in loc:
		s[v] = '('
	out += [''.join(s)]

	
if __name__ == '__main__':
	arr = generate(4)
	for i, x in enumerate(arr):
		print(f'{i:3}: {x}',)
		


