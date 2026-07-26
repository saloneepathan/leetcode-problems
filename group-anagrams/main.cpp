#include <iostream>
#include <vector>
#include <unordered_map>
#include <algorithm>
using namespace std;

vector<vector<string>> groupAnagrams(vector<string>& str){
    
    unordered_map<string, vector<string>> map;
    
    for (string s : str){
        string key = s;
        sort(key.begin(), key.end());
        map[key].push_back(s);
    }
    
    vector<vector<string>> result;
    
    for (auto element: map){
        result.push_back(element.second);
    }
    
    return result;
    
}

int main(){
    vector<string> str = {"hello", "llhoe", "meo", "ome"};
    
    vector<vector<string>> ans = groupAnagrams(str);
    
    for (auto group: ans){
        cout<<"[";
        for(string s: group){
            cout << s << " ";
        }
        cout <<"]\n";
    }
    
    return 0;
}