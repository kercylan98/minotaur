package astgo

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/utils/str"
	"os"
	"path/filepath"
)

func NewPackage(dir string) (*Package, error) {
	pkg := &Package{Dir: dir, Functions: map[string]*Function{}}
	fs, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		path := filepath.Join(pkg.Dir, f.Name())
		if f.IsDir() {
			pkg.Dirs = append(pkg.Dirs, path)
			continue
		}
		if filepath.Ext(path) != ".go" {
			continue
		}
		f, err := newFile(pkg, path)
		if err != nil {
			continue
		}
		pkg.Files = append(pkg.Files, f)
		for _, function := range f.Functions {
			function := function
			key := function.Name
			if function.Struct != nil {
				key = fmt.Sprintf("%s.%s", function.Struct.Name, key)
			}
			pkg.Functions[key] = function
		}
	}
	if len(pkg.Files) == 0 {
		return nil, errors.New("not found go ext file")
	}
	pkg.Name = pkg.Files[0].Package()

	return pkg, nil
}

type Package struct {
	Dir       string
	Name      string
	Dirs      []string
	Files     []*File
	Functions map[string]*Function
}

func (p *Package) StructFunc(name string) []*Function {
	var fs []*Function
	for _, file := range p.Files {
		for _, function := range file.Functions {
			if function.Struct == nil {
				continue
			}
			if function.Struct.Type.Name == name {
				fs = append(fs, function)
			}
		}
	}
	return fs
}

func (p *Package) PackageFunc() []*Function {
	var fs []*Function
	for _, file := range p.Files {
		for _, function := range file.Functions {
			if function.Struct == nil {
				fs = append(fs, function)
			}
		}
	}
	return fs
}

func (p *Package) Structs() []*Struct {
	var structs []*Struct
	for _, file := range p.Files {
		for _, s := range file.Structs {
			structs = append(structs, s)
		}
	}
	return structs
}

func (p *Package) FileComments() *Comment {
	var comment = newComment("", nil)
	for _, file := range p.Files {
		for _, c := range file.Comment.Comments {
			comment.Comments = append(comment.Comments, c)
		}
		for _, c := range file.Comment.Clear {
			comment.Clear = append(comment.Clear, c)
		}
	}
	return comment
}

func (p *Package) GetUnitTest(f *Function) *Function {
	if f.Struct == nil {
		return p.Functions[fmt.Sprintf("Test%s", str.FirstUpper(f.Name))]
	}
	return p.Functions[fmt.Sprintf("Test%s_%s", f.Struct.Type.Name, f.Name)]
}

func (p *Package) GetExampleTest(f *Function) *Function {
	if f.Struct == nil {
		return p.Functions[fmt.Sprintf("Example%s", str.FirstUpper(f.Name))]
	}
	return p.Functions[fmt.Sprintf("Example%s_%s", f.Struct.Type.Name, f.Name)]
}

func (p *Package) GetBenchmarkTest(f *Function) *Function {
	if f.Struct == nil {
		return p.Functions[fmt.Sprintf("Benchmark%s", str.FirstUpper(f.Name))]
	}
	return p.Functions[fmt.Sprintf("Benchmark%s_%s", f.Struct.Type.Name, f.Name)]
}
