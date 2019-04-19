/**
刚看到题目的时候第一反应是遍历长度，然后长的先往后移动，移动到和短链表
长度一样的地方，然后继续一起往后面走，直到NULL或者两个链表的元素相同。

但这里还有个更好点的做法就是两个指针同时遍历两个链表，然后先遍历完的指针
指向另外一条链表的头，继续遍历。

这个原理就是两个指针都遍历一下a,b。因为不管是指针A还是指针B遍历过的总的
长度都是链表A和链表B的长度，所以总是会碰到。和第一种做法差不多个人感觉，只是
感觉更优雅一点。

**/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {

        ListNode *FlagA = headA;
        ListNode *FlagB = headB;
        if (FlagA == NULL || FlagB == NULL){
            return NULL;
        }
        while(FlagA != FlagB){
            if(FlagA == NULL){
                FlagA = headB;
            }else{
                FlagA = FlagA->next;
            }
            if(FlagB == NULL){
                FlagB = headA;
            }else{
                FlagB = FlagB->next;
            }
        }
        return FlagA;
    }
};