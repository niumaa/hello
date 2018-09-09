/*修改映射
1. 在映射m中插入或修改元素
 m[key] = elem
2. 获取元素：
 elem = m[key]
3.删除元素：
 delete(m, key)
4.通过双赋值检测某个键是否存在：
 elem, ok = m[key]
 若key在m中，ok为true
 若key不在映射中，那么elem是该映射元素类型的零值。
 同样，从映射读取某个不存在的键时， 结果是映射的元素类型的零值。
 注：若elem或ok还未声明，你可以使用短变量声明：
 elem, ok : =m[key]
 */
