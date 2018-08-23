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