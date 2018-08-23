bool hasCycle(struct ListNode *head) {
struct ListNode *begin;
struct ListNode *end;
if(head == NULL){
return false;
}
begin = head;
end = head->next;
while(end != NULL){
if(begin == end){
return true;
}
end = end->next;
if(end != NULL){
end = end->next;
}
begin = begin->next;
}
return false;
}