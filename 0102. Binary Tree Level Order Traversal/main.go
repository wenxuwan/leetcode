class Solution {
	public:
	vector<vector<int>> levelOrder(TreeNode* root) {
	vector<vector<int>> ret;
	if (root == NULL) { return ret; }
	vector<int> level;
	deque<TreeNode*> d;
	d.push_back(root);
	d.push_back(NULL);
	while (!d.empty()) {
	TreeNode* n = d.front();
	d.pop_front();
	if (n == NULL) {
	if (level.size() != 0) {
	ret.push_back(level);
	level.clear();
	}
	if (d.empty()) { return ret; }
	else { d.push_back(NULL); }
	} else {
	if (n->left != NULL) {
	d.push_back(n->left);
	}
	if (n->right != NULL) {
	d.push_back(n->right);
	}
	level.push_back(n->val);
	}
	}
	return ret;
	}
};