// Two-Sums
// https://leetcode.com/problems/two-sum/

class Solution {
public:
	vector<int> twoSum(vector<int>& nums, int target);
	vector<int> twoSumNaive(vector<int>& nums, int target);
	vector<int> twoSumExperiment(vector<int>& nums, int target);
private:
    int find(vector<int>& nums, int lower, int goal);
};

// solution using a hashmap
vector<int> Solution::twoSum(vector<int>& nums, int target)
{
    unordered_map<int, int> m; // (value, index)
    int n = nums.size();
    for (int i = 0; i < n; i++) {
        int ideal = target - nums[i];
        auto res = m.find(ideal);
        if (res != m.end()) {
            return vector<int>{i, res->second};
        } else {
            m.insert({nums[i], i});
        }
    }
    return vector<int>{};
}

// Naive Solution with O(n^2) runtime complexity
// returns empty vector if there is an error.
vector<int> Solution::twoSumNaive(vector<int>& nums, int target)
{
    // asumming that length of array is greater than 2
    int n = nums.size();
    for (int i = 0; i < n-1; i++) {
        int a = nums[i];
        for (int j = i+1; j < n; j++) {
            int b = nums[j];
            if (a + b == target) {
                return vector<int>{i, j};
            }
        }
    }
    return vector<int>{};
}

// Slightly Improved by sorting first, allowing one to cull some of 
// the possibilities
vector<int> Solution::twoSumExperiment(vector<int>& nums, int target)
{
    int n = nums.size();
    sort(nums.begin(), nums.end()); // O(n * log(n))
    for (int i = 0; i < n-1; i++) {
        int a = nums[i];
        int ideal = target - a;
        int ans = find(nums, i+1, ideal); // O(log(n))
        if (ans != -1) {
            return vector<int>{i, ans};
        }
    }
    return vector<int>{};
}


// Does a binary search of a sorted array for a specific number,
// returns the index if successfull, otherwise returns -1.
// Will only search between indices 'lower' and nums.size()
int Solution::find(vector<int>& nums, int lower, int goal)
{
    int upper = nums.size();
    while (lower < upper) {
        int i = (upper - lower) / 2;
        if (goal == nums[i]) {
            return i;
        } else if (goal < nums[i]) {
            upper = i;
        } else {
            lower = i;
        }
    }
    return -1;
}


