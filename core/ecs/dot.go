package ecs

//
//func (w *world) GenerateDotFile(path string) {
//	var buffer bytes.Buffer
//	buffer.WriteString("digraph G {\n")
//
//	for _, art := range w.archetypes {
//		currKey := fmt.Sprint(parseDynamicBitSetKey(art.mask.Key()).Bits())
//		for key := range art.delEdges {
//			buffer.WriteString(fmt.Sprintf("  \"%v\" -> \"%v\" [label=\"del\"];\n", currKey, parseDynamicBitSetKey(key).Bits()))
//		}
//		for key := range art.addEdges {
//			buffer.WriteString(fmt.Sprintf("  \"%v\" -> \"%v\" [label=\"add\"];\n", currKey, parseDynamicBitSetKey(key).Bits()))
//		}
//	}
//
//	buffer.WriteString("}\n")
//
//	dotFileName := path
//	err := os.WriteFile(dotFileName, buffer.Bytes(), 0644)
//	if err != nil {
//		fmt.Println("Error writing dot file:", err)
//		return
//	}
//
//	// 生成图形文件
//	outFileName := dotFileName + ".png"
//	cmd := exec.Command("dot", "-Tpng", dotFileName, "-o", outFileName)
//	err = cmd.Run()
//	if err != nil {
//		fmt.Println("Error generating graph:", err)
//		return
//	}
//
//	fmt.Println("Graph generated and saved as", outFileName)
//}
//
