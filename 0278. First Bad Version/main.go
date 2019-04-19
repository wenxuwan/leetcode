package main


// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
	int firstBadVersion(int n) {
		if(isBadVersion(1)) return 1;
		int begin = 1;
		int end = n;
		int middle = 0;
		while(begin < end){
			middle = begin / 2 + end / 2;
			if(isBadVersion(middle)){
			end  = middle;
			}else{
			begin = middle + 1;
			}
		}
		return begin;
		}
};