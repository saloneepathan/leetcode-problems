class Solution {
    const int LL = 0, LR = 1, RL = 2, RR = 3; 
public:
    int maxWalls(vector<int>& robots, vector<int>& distance, vector<int>& walls) {
        robots.push_back(-1), robots.push_back(2e9);
        distance.push_back(0), distance.push_back(0);
        vector<int> rIndices(robots.size());
        iota(rIndices.begin(), rIndices.end(), 0);
        sort(rIndices.begin(), rIndices.end(), [&](int idx1, int idx2) {
            return robots[idx1] < robots[idx2];
        });
        int wallSize = walls.size();
        sort(walls.begin(), walls.end());
        int curr = 0;
        int lbot = rIndices[curr], rbot = rIndices[curr + 1];
        // dp[LL/LR/RL/RR]: the directions of the robots around the wall fire to
        int dp[4] = {0};
        for (int i = 0; i < wallSize; i++) {
            while (walls[i] > robots[rbot]) {
                ++curr;
                lbot = rIndices[curr];
                rbot = rIndices[curr + 1];
                int maxL = max(dp[LL], dp[RL]);
                int maxR = max(dp[LR], dp[RR]);
                dp[LL] = dp[LR] = maxL;
                dp[RL] = dp[RR] = maxR;
            }
            dp[LL] += walls[i] >= robots[rbot] - distance[rbot];
            dp[LR] += walls[i] == robots[rbot];
            dp[RL] += walls[i] <= robots[lbot] + distance[lbot] 
                        || walls[i] >= robots[rbot] - distance[rbot];
            dp[RR] += walls[i] <= robots[lbot] + distance[lbot]
                        || walls[i] == robots[rbot];
        }
        return *max_element(dp, dp + 4);
    }
};