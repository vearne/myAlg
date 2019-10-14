#include <iostream>
#include <cstdio>
#include <cstring>
#include <cstdlib>
#include <vector>
#include <climits>
#include <algorithm>
#include <cmath>
#define LL long long
#define INF 0x3f3f3f
using namespace std;
const int maxn = 1010;
vector<int>g[maxn];
int llink[maxn],n,m;
bool used[maxn];
bool dfs(int u) {
    for(int i = 0; i < g[u].size(); i++) {
        if(!used[g[u][i]]) {
            used[g[u][i]] = true;
            if(llink[g[u][i]] == -1 || dfs(llink[g[u][i]])) {
                llink[g[u][i]] = u;
                return true;
            }
        }
    }
    return false;
}
int main() {
    int t,i,j,u,v;
    scanf("%d",&t);
    while(t--) {
        scanf("%d%d",&n,&m);
        v = n<<1;
        for(i = 0; i <= v; i++) {
            llink[i] = -1;
            g[i].clear();
        }
        for(i = 0; i < m; i++) {
            scanf("%d%d",&u,&v);
            g[u].push_back(n+v);
        }
        n <<= 1;
        int ans = 0;
        for(i = 1; i <= n; i++) {
            memset(used,false,sizeof(used));
            if(dfs(i)) {
                cout << llink << endl;
                ans++;
            }
        }
        printf("%d\n",ans);
    }
    return 0;
}