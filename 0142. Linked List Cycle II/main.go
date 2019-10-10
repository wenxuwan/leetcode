/*
快慢指针的问题，这个问题的关键在于，如果存在交点

那么快慢指针在遇到的那一点开始，到重合点的距离和从头到重合点的距离
相同
 */
 */
 */
*/


struct ListNode *detectCycle(struct ListNode *head) {
	struct ListNode *begin;
	struct ListNode *end;
	if(head == NULL){
		return head;
	}
	begin = head;
	end = head;
	while(end != NULL){
		end = end->next;
		if(end != NULL){
		end = end->next;
	}
	begin = begin->next;
	if(begin == end){
		break;
	}
	}
	if(end == NULL){
		return NULL;
	}
	begin = head;
	while(begin != end){
		begin = begin->next;
		end = end->next;
	}
	return begin;
}