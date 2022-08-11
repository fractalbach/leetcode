// https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution {
public:	
	array<int, 256> lookup;

    int lengthOfLongestSubstring(string s) {

		lookup.fill(-1);
		
		char c;
		int cur = 0;
		int len = 0;
		int best_len = 0;			    
		const int n = s.size();
        
        for (int i = 0; i < n; ++i) {
        
        	c = s[i];
        
			if (lookup[c] != -1 && lookup[c] >= cur) {
        		cur = lookup[c] + 1;
        		len = i - cur;
        	}
        	
        	lookup[c] = i;
        	len = len + 1;
        	
        	best_len = max(len, best_len);
        }
        
        return best_len;
        
    }
};

