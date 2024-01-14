package astgo

import (
	"errors"
	"os"
	"path/filepath"
)

func NewPackage(dir string) (*Package, error) {
	pkg := &Package{Dir: dir}
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

	}
	if len(pkg.Files) == 0 {
		return nil, errors.New("not found go ext file")
	}
	pkg.Name = pkg.Files[0].Package()

	return pkg, nil
}

type Package struct {
	Dir   string
	Name  string
	Dirs  []string
	Files []*File
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
	var comment = newComment(nil)
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
