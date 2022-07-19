/**
 * @date: 2022/7/19
 * @desc: 前缀树的基本实现
 */

package tea

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配(狂野的粗狂的)，part 含有 : 或 * 时为true  , 为 false 时为精准匹配
}

// matchChild 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 匹配所有成功的节点, 用于查找
func (n *node) matchChildren(part string) (resNodes []*node) {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			resNodes = append(resNodes, child)
		}
	}
	return resNodes
}

// insert 注册插入节点
// 递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，
// 有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。
// p和:lang节点的pattern属性皆为空。因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
// 例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败。
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { // parts 长度与 height-1 相等为最后一层, 最后一层设置n.pattern
		n.pattern = pattern
		return
	}
	part := parts[height]       // 获取当前part
	child := n.matchChild(part) // 从前缀树中获取第一个成功的节点
	if child == nil {           // 如果没有则新建 part 节点,并加入children切片中
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	height += 1
	child.insert(pattern, parts, height)
}

// search 路由匹配节点
// 查询功能，同样也是递归查询每一层的节点，退出规则是: 匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		fmt.Println(result)
		if result != nil {
			return result
		}
	}
	return nil
}
