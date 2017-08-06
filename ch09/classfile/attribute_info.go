package classfile
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/*将所有的属性出来*/
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo{
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}
/*将一个属性读出来*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo{
	//读取属性名索引
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	//读取属性长度
	attrLen := reader.readUint32()
	//实例化属性
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
func newAttributeInfo(attrName string,attrLen uint32,cp ConstantPool) AttributeInfo {
	//23种预定义属性分为3组，第一组是实现虚拟机所必须得，共5种，第二组是java类库所必须得，共有12种；第三组属性提供给工具使用共6种，是可选的

	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}//1组
	case "ConstantValue": return &ConstantValueAttribute{}//1组
	case "Deprecated": return &DeprecatedAttribute{}//3组
	case "Exceptions": return &ExceptionsAttribute{}//1组
	case "LineNumberTable": return &LineNumberTableAttribute{}//3组 在异常堆栈中显示行号
	case "LocalVariableTable": return &LocalVariableTableAttribute{}//3组
	case "SourceFile": return &SourceFileAttribute{}//3组
	case "Synthetic": return &SyntheticAttribute{}//2组
	default: return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
